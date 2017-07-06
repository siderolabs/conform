package git

import (
	"fmt"

	"github.com/Masterminds/semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Git is a helper for git.
type Git struct {
	repo *git.Repository
}

// NewGit instantiates and returns a Git struct.
func NewGit() (g *Git, err error) {
	repo, err := git.PlainOpen("./")
	if err != nil {
		return
	}
	g = &Git{repo: repo}

	return
}

// Branch returns the current git branch name.
func (g *Git) Branch() (branch string, isBranch bool, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}
	if ref.IsBranch() {
		isBranch = true
		branch = ref.Name().Short()
	}

	return
}

// SHA returns the sha of the current commit.
func (g *Git) SHA() (sha string, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}
	sha = ref.Hash().String()[0:7]

	return
}

// Tag returns the tag name if HEAD is a tag.
func (g *Git) Tag() (tag string, isTag bool, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}
	tags, err := g.repo.Tags()
	if err != nil {
		return
	}
	err = tags.ForEach(func(_ref *plumbing.Reference) error {
		if _ref.Hash().String() == ref.Hash().String() {
			isTag = true
			tag = _ref.Name().Short()
			return nil
		}
		return nil
	})
	if err != nil {
		return
	}

	return
}

// Prerelease returns the prerelease name if the tag is a prerelease.
func (g *Git) Prerelease(tag string, isTag bool) (prerelease string, isPrerelease bool, err error) {
	if isTag {
		var ver *semver.Version
		ver, err = semver.NewVersion(tag[1:])
		if err != nil {
			return
		}
		if ver.Prerelease() != "" {
			prerelease = ver.Prerelease()
			isPrerelease = true
		}
	}

	return
}

// Status returns the status of the working tree.
func (g *Git) Status() (status string, isClean bool, err error) {
	worktree, err := g.repo.Worktree()
	if err != nil {
		return
	}
	worktreeStatus, err := worktree.Status()
	if err != nil {
		return
	}
	if worktreeStatus.IsClean() {
		isClean = true
		status = " nothing to commit, working tree clean"
	} else {
		status = worktreeStatus.String()
	}

	return
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
	fmt.Println(commit.NumParents())
	if commit.NumParents() > 1 {
		parents := commit.Parents()
		for i := 1; i <= commit.NumParents(); i++ {
			next, err := parents.Next()
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

	return
}
