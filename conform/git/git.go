package git

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
)

// Info contains the status of the working tree.
type Info struct {
	Branch     string
	SHA        string
	Tag        string
	Prerelease string
	IsTag      bool
	IsDirty    bool
}

// NewInfo instantiates and returns info.
func NewInfo() (info *Info, err error) {
	branch, err := Branch()
	if err != nil {
		return
	}

	sha, err := SHA()
	if err != nil {
		return
	}

	tag, isTag, err := Tag()
	if err != nil {
		return
	}

	_, isDirty, err := Status()
	if err != nil {
		return
	}

	info = &Info{
		Branch:  branch,
		SHA:     sha,
		Tag:     strings.TrimSuffix(tag, "\n"),
		IsTag:   isTag,
		IsDirty: isDirty,
	}

	return
}

// Branch returns the current git branch name.
func Branch() (branch string, err error) {
	branchBytes, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return
	}
	branch = strings.TrimSuffix(string(branchBytes), "\n")
	err = ExportConformVar("branch", branch)
	fmt.Printf("Branch: %s\n", branch)

	return
}

// SHA returns the sha of the current commit.
func SHA() (sha string, err error) {
	shaBytes, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return
	}
	sha = strings.TrimSuffix(string(shaBytes), "\n")
	err = ExportConformVar("sha", sha)
	fmt.Printf("SHA: %s\n", sha)

	return
}

// Tag returns the tag name if HEAD is a tag.
func Tag() (tag string, isTag bool, err error) {
	tagBytes, isTagErr := exec.Command("git", "describe", "--exact-match", "--tags", "HEAD").Output()
	if isTagErr == nil {
		isTag = true
	}
	tag = strings.TrimSuffix(string(tagBytes), "\n")
	if isTag {
		_, err = semver.NewVersion(tag[1:])
		if err != nil {
			return
		}
	}
	err = ExportConformVar("tag", tag)
	if err != nil {
		return
	}
	err = ExportConformVar("is_tag", strconv.FormatBool(isTag))
	if err != nil {
		return
	}
	fmt.Printf("IsTag: %v\n", isTag)
	fmt.Printf("Tag: %s\n", tag)

	return
}

// Status returns the status of the working tree.
func Status() (status string, isDirty bool, err error) {
	statusBytes, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return
	}
	status = strings.TrimSuffix(string(statusBytes), "\n")
	if status != "" {
		isDirty = true
	}
	err = ExportConformVar("is_dirty", strconv.FormatBool(isDirty))
	if err != nil {
		return
	}
	fmt.Printf("Status: %s\n", status)
	fmt.Printf("IsDirty: %v\n", isDirty)

	return
}

// ExportConformVar exports variable prefixed with CONFORM_
func ExportConformVar(name, value string) (err error) {
	variable := fmt.Sprintf("CONFORM_%s", strings.ToUpper(name))
	err = os.Setenv(variable, value)

	return
}
