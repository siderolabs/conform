// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package git provides helpers for SCM.
package git

import (
	"fmt"
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

// OpenRepository opens the repository containing path, including linked
// worktrees whose .git entry is a gitdir file.
func OpenRepository(path string) (*git.Repository, error) {
	return git.PlainOpenWithOptions(path, &git.PlainOpenOptions{
		DetectDotGit:          true,
		EnableDotGitCommonDir: true,
	})
}

// NewGit instantiates and returns a Git struct.
func NewGit() (*Git, error) {
	repo, err := OpenRepository(".")
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
		return "", err
	}

	commit, err := g.repo.CommitObject(ref.Hash())
	if err != nil {
		return "", err
	}

	if commit.NumParents() > 1 {
		parents := commit.Parents()

		for i := 1; i <= commit.NumParents(); i++ {
			var next *object.Commit

			next, err = parents.Next()
			if err != nil {
				return "", err
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

// HeadCommit returns the commit HEAD points to.
func (g *Git) HeadCommit() (*object.Commit, error) {
	ref, err := g.repo.Head()
	if err != nil {
		return nil, err
	}

	return g.repo.CommitObject(ref.Hash())
}

// Commits returns the list of commits in the range commit1..commit2: the commits
// reachable from commit2 which are not reachable from commit1 (as `git log commit1..commit2`).
//
// All parents of merge commits are followed, so commits merged into the range are included.
// The commits are returned newest first (breadth-first order starting from commit2).
func (g *Git) Commits(commit1, commit2 string) ([]*object.Commit, error) {
	c1, err := g.resolveCommit(commit1)
	if err != nil {
		return nil, err
	}

	c2, err := g.resolveCommit(commit2)
	if err != nil {
		return nil, err
	}

	excluded, err := g.ancestors(c1)
	if err != nil {
		return nil, err
	}

	if _, ok := excluded[c2.Hash]; ok {
		if c2.Hash == c1.Hash {
			return nil, nil
		}

		return nil, errors.Errorf("invalid revision range: %s is not an ancestor of %s", commit1, commit2)
	}

	var (
		commits     []*object.Commit
		reachedBase bool
	)

	seen := map[plumbing.Hash]struct{}{c2.Hash: {}}
	queue := []*object.Commit{c2}

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		commits = append(commits, c)

		for _, parentHash := range c.ParentHashes {
			if _, ok := excluded[parentHash]; ok {
				reachedBase = true

				continue
			}

			if _, ok := seen[parentHash]; ok {
				continue
			}

			seen[parentHash] = struct{}{}

			parent, err := g.repo.CommitObject(parentHash)
			if err != nil {
				return nil, errors.Errorf("failed to get parent %s of commit %s: %v", parentHash, c.Hash, err)
			}

			queue = append(queue, parent)
		}
	}

	if !reachedBase {
		return nil, errors.Errorf("invalid revision range: %s and %s have no common history", commit1, commit2)
	}

	return commits, nil
}

func (g *Git) resolveCommit(rev string) (*object.Commit, error) {
	hash, err := g.repo.ResolveRevision(plumbing.Revision(rev))
	if err != nil {
		return nil, err
	}

	return g.repo.CommitObject(*hash)
}

// ancestors returns the set of commits reachable from the commit, including the commit itself.
//
// Missing commits (e.g. beyond the boundary of a shallow clone) are skipped.
func (g *Git) ancestors(c *object.Commit) (map[plumbing.Hash]struct{}, error) {
	result := map[plumbing.Hash]struct{}{c.Hash: {}}
	queue := []*object.Commit{c}

	for len(queue) > 0 {
		c = queue[0]
		queue = queue[1:]

		for _, parentHash := range c.ParentHashes {
			if _, ok := result[parentHash]; ok {
				continue
			}

			parent, err := g.repo.CommitObject(parentHash)
			if err != nil {
				if errors.Is(err, plumbing.ErrObjectNotFound) {
					continue
				}

				return nil, err
			}

			result[parentHash] = struct{}{}

			queue = append(queue, parent)
		}
	}

	return result, nil
}

// HasGPGSignature returns true if the commit is signed.
func HasGPGSignature(commit *object.Commit) bool {
	return commit.PGPSignature != ""
}

// ReadArmoredKeyrings parses a list of armored keyrings into a single keyring.
func ReadArmoredKeyrings(armoredKeyrings []string) (openpgp.EntityList, error) {
	var keyring openpgp.EntityList

	for _, armoredKeyring := range armoredKeyrings {
		el, err := openpgp.ReadArmoredKeyRing(strings.NewReader(armoredKeyring))
		if err != nil {
			return nil, err
		}

		keyring = append(keyring, el...)
	}

	return keyring, nil
}

// VerifyPGPSignature validates the commit PGP signature against a keyring.
func VerifyPGPSignature(commit *object.Commit, keyring openpgp.EntityList) (*openpgp.Entity, error) {
	// Extract signature.
	signature := strings.NewReader(commit.PGPSignature)

	encoded := &plumbing.MemoryObject{}

	// Encode commit components, excluding signature and get a reader object.
	if err := commit.EncodeWithoutSignature(encoded); err != nil {
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
