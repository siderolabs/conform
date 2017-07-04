package git

import (
	"github.com/Masterminds/semver"
	// git "github.com/libgit2/git2go"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Info contains the status of the working tree.
type Info struct {
	Branch       string
	SHA          string
	Tag          string
	Prerelease   string
	Status       string
	IsTag        bool
	IsPrerelease bool
	IsDirty      bool
}

// NewInfo instantiates and returns info.
func NewInfo() (info *Info, err error) {
	repo, err := git.PlainOpen("./")
	if err != nil {
		return
	}

	branch, err := Branch(repo)
	if err != nil {
		return
	}

	sha, err := SHA(repo)
	if err != nil {
		return
	}

	tag, isTag, err := Tag(repo)
	if err != nil {
		return
	}

	prerelease, isPrerelease, err := Prerelease(tag, isTag)
	if err != nil {
		return
	}

	status, isDirty, err := Status(repo)
	if err != nil {
		return
	}

	info = &Info{
		Branch:       branch,
		SHA:          sha,
		Tag:          tag,
		Prerelease:   prerelease,
		Status:       status,
		IsTag:        isTag,
		IsPrerelease: isPrerelease,
		IsDirty:      isDirty,
	}

	return
}

// Branch returns the current git branch name.
func Branch(repo *git.Repository) (branch string, err error) {
	ref, err := repo.Head()
	if err != nil {
		return
	}
	if ref.IsBranch() {
		branch = ref.Name().Short()
	}

	return
}

// SHA returns the sha of the current commit.
func SHA(repo *git.Repository) (sha string, err error) {
	ref, err := repo.Head()
	if err != nil {
		return
	}
	sha = ref.Hash().String()[0:7]

	return
}

// Tag returns the tag name if HEAD is a tag.
func Tag(repo *git.Repository) (tag string, isTag bool, err error) {
	ref, err := repo.Head()
	if err != nil {
		return
	}
	tags, err := repo.Tags()
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
func Prerelease(tag string, isTag bool) (prerelease string, isPrerelease bool, err error) {
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
func Status(repo *git.Repository) (status string, isDirty bool, err error) {
	worktree, err := repo.Worktree()
	if err != nil {
		return
	}
	worktreeStatus, err := worktree.Status()
	if worktreeStatus.IsClean() {
		status = " nothing to commit, working tree clean"
	} else {
		isDirty = true
		status = worktreeStatus.String()
	}

	return
}
