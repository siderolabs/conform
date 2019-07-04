/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package git

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
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
func NewGit() (g *Git, err error) {
	p, err := findDotGit(".git")
	if err != nil {
		return
	}
	repo, err := git.PlainOpen(path.Dir(p))
	if err != nil {
		return
	}
	g = &Git{repo: repo}

	return g, err
}

// Message returns the commit message. In the case that a commit has multiple
// parents, the message of the last parent is returned.
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

// HasGPGSignature returns the commit message. In the case that a commit has multiple
// parents, the message of the last parent is returned.
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

// FetchPullRequest fetches a remote PR.
func (g *Git) FetchPullRequest(remote string, number int) (err error) {
	opts := &git.FetchOptions{
		RemoteName: remote,
		RefSpecs: []config.RefSpec{
			config.RefSpec(fmt.Sprintf("refs/pull/%d/head:pr/%d", number, number)),
		},
	}
	if err = g.repo.Fetch(opts); err != nil {
		return err
	}

	return nil
}

// CheckoutPullRequest checks out pull request.
func (g *Git) CheckoutPullRequest(number int) (err error) {
	w, err := g.repo.Worktree()
	if err != nil {
		return err
	}

	opts := &git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("pr/%d", number)),
	}

	if err := w.Checkout(opts); err != nil {
		return err
	}

	return nil
}

// SHA returns the sha of the current commit.
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
func (g *Git) AheadBehind(ref string) (ahead int, behind int, err error) {
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
		return 0, 0, nil
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
		return 0, 0, nil
	}

	return count, 0, nil
}
