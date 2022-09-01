// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package git provides helpers for SCM.
package git

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/keybase/go-crypto/openpgp"
	"github.com/pkg/errors"
)

// Git is a helper for git.
type Git struct {
	repo *git.Repository
}

func findDotGit(name string) (string, error) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return findDotGit(path.Join("..", name))
	}

	return filepath.Abs(name)
}

// NewGit instantiates and returns a Git struct.
func NewGit() (*Git, error) {
	p, err := findDotGit(".git")
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(path.Dir(p))
	if err != nil {
		return nil, err
	}

	return &Git{repo: repo}, nil
}

// Message returns the commit message. In the case that a commit has multiple
// parents, the message of the last parent is returned.
//
//nolint:nonamedreturns
func (g *Git) Message() (message string, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}

	commit, err := g.repo.CommitObject(ref.Hash())
	if err != nil {
		return
	}

	if commit.NumParents() > 1 {
		parents := commit.Parents()

		for i := 1; i <= commit.NumParents(); i++ {
			var next *object.Commit

			next, err = parents.Next()
			if err != nil {
				return
			}

			if i == commit.NumParents() {
				message = next.Message
			}
		}
	} else {
		message = commit.Message
	}

	return message, err
}

// Messages returns the list of commit messages in the range commit1..commit2.
func (g *Git) Messages(commit1, commit2 string) ([]string, error) {
	hash1, err := g.repo.ResolveRevision(plumbing.Revision(commit1))
	if err != nil {
		return nil, err
	}

	hash2, err := g.repo.ResolveRevision(plumbing.Revision(commit2))
	if err != nil {
		return nil, err
	}

	c2, err := g.repo.CommitObject(*hash2)
	if err != nil {
		return nil, err
	}

	c1, err := g.repo.CommitObject(*hash1)
	if err != nil {
		return nil, err
	}

	if ok, ancestorErr := c1.IsAncestor(c2); ancestorErr != nil || !ok {
		c, mergeBaseErr := c1.MergeBase(c2)
		if mergeBaseErr != nil {
			return nil, errors.Errorf("invalid ancestor %s", c1)
		}

		c1 = c[0]
	}

	msgs := make([]string, 0)

	for {
		msgs = append(msgs, c2.Message)

		c2, err = c2.Parents().Next()
		if err != nil {
			return nil, err
		}

		if c2.ID() == c1.ID() {
			break
		}
	}

	return msgs, nil
}

// HasGPGSignature returns the commit message. In the case that a commit has multiple
// parents, the message of the last parent is returned.
//
//nolint:nonamedreturns
func (g *Git) HasGPGSignature() (ok bool, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return false, err
	}

	commit, err := g.repo.CommitObject(ref.Hash())
	if err != nil {
		return false, err
	}

	ok = commit.PGPSignature != ""

	return ok, err
}

// VerifyPGPSignature validates PGP signature against a keyring.
func (g *Git) VerifyPGPSignature(armoredKeyrings []string) (*openpgp.Entity, error) {
	ref, err := g.repo.Head()
	if err != nil {
		return nil, err
	}

	commit, err := g.repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	var keyring openpgp.EntityList

	for _, armoredKeyring := range armoredKeyrings {
		var el openpgp.EntityList

		el, err = openpgp.ReadArmoredKeyRing(strings.NewReader(armoredKeyring))
		if err != nil {
			return nil, err
		}

		keyring = append(keyring, el...)
	}

	// Extract signature.
	signature := strings.NewReader(commit.PGPSignature)

	encoded := &plumbing.MemoryObject{}

	// Encode commit components, excluding signature and get a reader object.
	if err = commit.EncodeWithoutSignature(encoded); err != nil {
		return nil, err
	}

	er, err := encoded.Reader()
	if err != nil {
		return nil, err
	}

	return openpgp.CheckArmoredDetachedSignature(keyring, er, signature)
}

// FetchPullRequest fetches a remote PR.
//
//nolint:nonamedreturns
func (g *Git) FetchPullRequest(remote string, number int) (err error) {
	opts := &git.FetchOptions{
		RemoteName: remote,
		RefSpecs: []config.RefSpec{
			config.RefSpec(fmt.Sprintf("refs/pull/%d/head:pr/%d", number, number)),
		},
	}

	return g.repo.Fetch(opts)
}

// CheckoutPullRequest checks out pull request.
//
//nolint:nonamedreturns
func (g *Git) CheckoutPullRequest(number int) (err error) {
	w, err := g.repo.Worktree()
	if err != nil {
		return err
	}

	opts := &git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("pr/%d", number)),
	}

	return w.Checkout(opts)
}

// SHA returns the sha of the current commit.
//
//nolint:nonamedreturns
func (g *Git) SHA() (sha string, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return sha, err
	}

	sha = ref.Hash().String()

	return sha, nil
}

// AheadBehind returns the number of commits that HEAD is ahead and behind
// relative to the specified ref.
//
//nolint:nonamedreturns
func (g *Git) AheadBehind(ref string) (ahead, behind int, err error) {
	ref1, err := g.repo.Reference(plumbing.ReferenceName(ref), false)
	if err != nil {
		return 0, 0, err
	}

	ref2, err := g.repo.Head()
	if err != nil {
		return 0, 0, err
	}

	commit2, err := object.GetCommit(g.repo.Storer, ref2.Hash())
	if err != nil {
		return 0, 0, nil //nolint:nilerr
	}

	var count int

	iter := object.NewCommitPreorderIter(commit2, nil, nil)

	err = iter.ForEach(func(comm *object.Commit) error {
		if comm.Hash != ref1.Hash() {
			count++

			return nil
		}

		return storer.ErrStop
	})

	if err != nil {
		return 0, 0, nil //nolint:nilerr
	}

	return count, 0, nil
}
