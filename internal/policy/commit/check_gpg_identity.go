// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package commit

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/google/go-github/github"
	"golang.org/x/sync/errgroup"

	"github.com/talos-systems/conform/internal/git"
	"github.com/talos-systems/conform/internal/policy"
)

// GPGIdentityCheck ensures that the commit is cryptographically signed using known identity.
//
//nolint:govet
type GPGIdentityCheck struct {
	errors   []error
	identity string
}

// Name returns the name of the check.
func (g GPGIdentityCheck) Name() string {
	return "GPG Identity"
}

// Message returns to check message.
func (g GPGIdentityCheck) Message() string {
	if len(g.errors) != 0 {
		return g.errors[0].Error()
	}

	return fmt.Sprintf("Signed by %q", g.identity)
}

// Errors returns any violations of the check.
func (g GPGIdentityCheck) Errors() []error {
	return g.errors
}

// ValidateGPGIdentity checks the commit GPG signature for a known identity.
func (c Commit) ValidateGPGIdentity(g *git.Git) policy.Check {
	check := &GPGIdentityCheck{}

	switch {
	case c.GPG.Identity.GitHubOrganization != "":
		githubClient := github.NewClient(nil)

		list, _, err := githubClient.Organizations.ListMembers(context.Background(), c.GPG.Identity.GitHubOrganization, &github.ListMembersOptions{})
		if err != nil {
			check.errors = append(check.errors, err)

			return check
		}

		members := make([]string, len(list))

		for i := range list {
			members[i] = list[i].GetLogin()
		}

		keyrings, err := getKeyring(context.Background(), members)
		if err != nil {
			check.errors = append(check.errors, err)

			return check
		}

		entity, err := g.VerifyPGPSignature(keyrings)
		if err != nil {
			check.errors = append(check.errors, err)

			return check
		}

		for identity := range entity.Identities {
			check.identity = identity

			break
		}
	default:
		check.errors = append(check.errors, fmt.Errorf("no signature identity configuration found"))
	}

	return check
}

func getKeyring(ctx context.Context, members []string) ([]string, error) {
	var (
		result []string
		mu     sync.Mutex
	)

	eg, ctx := errgroup.WithContext(ctx)

	for _, member := range members {
		member := member

		eg.Go(func() error {
			key, err := getKey(ctx, member)

			mu.Lock()
			result = append(result, key)
			mu.Unlock()

			return err
		})
	}

	err := eg.Wait()

	return result, err
}

func getKey(ctx context.Context, login string) (string, error) {
	// GitHub client doesn't have a method to fetch a key unauthenticated
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://github.com/%s.gpg", login), nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() //nolint:errcheck

	buf, err := ioutil.ReadAll(resp.Body)

	return string(buf), err
}
