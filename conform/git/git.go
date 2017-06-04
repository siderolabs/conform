package git

import (
	"fmt"
	"log"
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
func NewInfo() *Info {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatalf("Failed to get branch [%v]", err)
	}
	sha, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		log.Fatalf("Failed to get sha [%v]", err)
	}
	isTag := false
	_, err = exec.Command("git", "symbolic-ref", "HEAD").Output()
	if err != nil {
		isTag = true
	}
	tag := "undefined"
	if isTag {
		_tag, err2 := exec.Command("git", "describe", "--exact-match", "HEAD").Output()
		if err2 == nil {
			tag = string(_tag)
		}
	}
	status, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		log.Fatal(err)
	}
	isDirty := false
	if strings.TrimSuffix(string(status), "\n") != "" {
		isDirty = true
	}

	prerelease := ""
	if isTag {
		sv, err := semver.NewVersion(strings.TrimSuffix(string(tag[1:]), "\n"))
		if err != nil {
			log.Fatal(err)
		}

		prerelease = sv.Prerelease()
	}

	fmt.Printf("Branch: %s\n", strings.TrimSuffix(string(branch), "\n"))
	os.Setenv("CONFORM_BRANCH", strings.TrimSuffix(string(branch), "\n"))
	fmt.Printf("SHA: %s\n", strings.TrimSuffix(string(sha), "\n"))
	os.Setenv("CONFORM_SHA", strings.TrimSuffix(string(sha), "\n"))
	fmt.Printf("Tag: %s\n", strings.TrimSuffix(string(tag), "\n"))
	os.Setenv("CONFORM_TAG", strings.TrimSuffix(string(tag), "\n"))
	fmt.Printf("Status: %s\n", strings.TrimSuffix(string(status), "\n"))
	fmt.Printf("IsTag: %v\n", isTag)
	os.Setenv("CONFORM_IS_TAG", strconv.FormatBool(isTag))
	fmt.Printf("Prerelease: %s\n", prerelease)
	os.Setenv("CONFORM_PRERELEASE", prerelease)
	fmt.Printf("IsDirty: %v\n", isDirty)
	os.Setenv("CONFORM_IS_DIRTY", strconv.FormatBool(isDirty))
	// os.Setenv("CONFORM_IMAGE", strconv.FormatBool(isTag))

	return &Info{
		Branch:     strings.TrimSuffix(string(branch), "\n"),
		SHA:        strings.TrimSuffix(string(sha), "\n"),
		Tag:        strings.TrimSuffix(string(tag), "\n"),
		Prerelease: prerelease,
		IsTag:      isTag,
		IsDirty:    isDirty,
	}
}
