package git

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/autonomy/conform/conform/utilities"
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
	Message      string
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

	message, err := Message(repo, isDirty)
	if err != nil {
		return
	}

	info = &Info{
		Branch:       branch,
		SHA:          sha,
		Tag:          tag,
		Prerelease:   prerelease,
		Status:       status,
		Message:      message,
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

	fmt.Printf("Branch: %s\n", branch)
	err = utilities.ExportConformVar("branch", branch)
	if err != nil {
		return
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

	fmt.Printf("SHA: %s\n", sha)
	err = utilities.ExportConformVar("sha", sha)
	if err != nil {
		return
	}

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

	fmt.Printf("Tag: %s\n", tag)
	err = utilities.ExportConformVar("tag", tag)
	if err != nil {
		return
	}
	fmt.Printf("IsTag: %s\n", strconv.FormatBool(isTag))
	err = utilities.ExportConformVar("is_tag", strconv.FormatBool(isTag))
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

	fmt.Printf("Prerelease: %s\n", prerelease)
	err = utilities.ExportConformVar("prerelease", prerelease)
	if err != nil {
		return
	}
	fmt.Printf("IsPrerelease: %s\n", strconv.FormatBool(isPrerelease))
	err = utilities.ExportConformVar("is_prerelease", strconv.FormatBool(isPrerelease))
	if err != nil {
		return
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
	if err != nil {
		return
	}
	if worktreeStatus.IsClean() {
		status = " nothing to commit, working tree clean"
	} else {
		isDirty = true
		status = worktreeStatus.String()
	}

	fmt.Printf("Status: \n%s\n", strings.TrimRight(status, "\n"))
	err = utilities.ExportConformVar("status", strconv.FormatBool(isDirty))
	if err != nil {
		return
	}
	fmt.Printf("IsDirty: %s\n", strconv.FormatBool(isDirty))
	err = utilities.ExportConformVar("is_dirty", strconv.FormatBool(isDirty))
	if err != nil {
		return
	}

	return
}

// Message returns the commit message. In the case that a commit has multiple
// parents, the message of the last parent is returned.
func Message(repo *git.Repository, isDirty bool) (message string, err error) {
	ref, err := repo.Head()
	if err != nil {
		return
	}
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return
	}
	if commit.NumParents() != 1 {
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

	if !isDirty {
		fmt.Printf("Message: %s\n", strings.TrimRight(message, "\n"))
		err = utilities.ExportConformVar("message", message)
		if err != nil {
			return
		}
	}

	return
}
