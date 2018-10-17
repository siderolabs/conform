package git

import (
	"os"
	"path"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
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

// Branch returns the current git branch name.
func (g *Git) Branch() (branch string, isBranch bool, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}
	if ref.Name().IsBranch() {
		isBranch = true
		branch = ref.Name().Short()
	}

	return branch, isBranch, err
}

// Ref returns the current git ref name.
func (g *Git) Ref() (ref string, err error) {
	r, err := g.repo.Head()
	if err != nil {
		return
	}

	ref = r.Name().String()

	return ref, err
}

// SHA returns the sha of the current commit.
func (g *Git) SHA() (sha string, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}
	sha = ref.Hash().String()[0:7]

	return sha, err
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

	return tag, isTag, err
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

	return status, isClean, err
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
