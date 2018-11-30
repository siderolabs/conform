package git

import (
	"os"
	"path"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
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
