// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package commit

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/google/go-github/v60/github"
	"github.com/keybase/go-crypto/openpgp"
	"golang.org/x/sync/errgroup"

	"github.com/siderolabs/conform/internal/git"
	"github.com/siderolabs/conform/internal/policy"
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
func (c Commit) ValidateGPGIdentity() policy.Check { //nolint:ireturn
	check := &GPGIdentityCheck{}

	if c.GPG.Identity.GitHubOrganization == "" {
		check.errors = append(check.errors, errors.New("no signature identity configuration found"))

		return check
	}

	if !git.HasGPGSignature(c.commit) {
		check.errors = append(check.errors, fmt.Errorf("commit %s: no GPG signature", c.commit.Hash.String()[:shortHashLength]))

		return check
	}

	keyring, err := c.keyring.get(context.Background(), c.GPG.Identity.GitHubOrganization)
	if err != nil {
		check.errors = append(check.errors, err)

		return check
	}

	entity, err := git.VerifyPGPSignature(c.commit, keyring)
	if err != nil {
		check.errors = append(check.errors, fmt.Errorf("commit %s: %w", c.commit.Hash.String()[:shortHashLength], err))

		return check
	}

	for identity := range entity.Identities {
		check.identity = identity

		break
	}

	return check
}

// gpgKeyring fetches the keyring of the GitHub organization members once.
type gpgKeyring struct {
	err     error
	keyring openpgp.EntityList
	once    sync.Once
}

func (k *gpgKeyring) get(ctx context.Context, org string) (openpgp.EntityList, error) {
	k.once.Do(func() {
		k.keyring, k.err = fetchOrgKeyring(ctx, org)
	})

	return k.keyring, k.err
}

// fetchOrgKeyring is a variable to allow overriding it in tests.
var fetchOrgKeyring = githubOrgKeyring

// githubOrgKeyring builds the keyring from the GPG keys of the public members of the GitHub organization.
//
// If GITHUB_TOKEN (or INPUT_TOKEN) is set, the GitHub API is accessed with it to get higher rate limits.
func githubOrgKeyring(ctx context.Context, org string) (openpgp.EntityList, error) {
	githubClient := github.NewClient(nil)

	if token := githubToken(); token != "" {
		githubClient = githubClient.WithAuthToken(token)
	}

	opts := &github.ListMembersOptions{
		PublicOnly: true,
		PerPage:    100,
	}

	var members []string

	for {
		list, resp, err := githubClient.Organizations.ListMembers(ctx, org, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list members of GitHub organization %q: %w", org, err)
		}

		for _, member := range list {
			members = append(members, member.GetLogin())
		}

		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	armoredKeyrings, err := getKeyring(ctx, members)
	if err != nil {
		return nil, err
	}

	return git.ReadArmoredKeyrings(armoredKeyrings)
}

func githubToken() string {
	for _, env := range []string{"GITHUB_TOKEN", "INPUT_TOKEN"} {
		if token := os.Getenv(env); token != "" {
			return token
		}
	}

	return ""
}

func getKeyring(ctx context.Context, members []string) ([]string, error) {
	const concurrency = 8

	result := make([]string, len(members))

	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(concurrency)

	for i, member := range members {
		eg.Go(func() error {
			key, err := getKey(ctx, member)
			if err != nil {
				return err
			}

			result[i] = key

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return result, nil
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

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch GPG keys of %q: unexpected status %s", login, resp.Status)
	}

	buf, err := io.ReadAll(resp.Body)

	return string(buf), err
}
