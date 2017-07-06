package metadata

import (
	"fmt"
	"time"

	"github.com/autonomy/conform/conform/git"
	"github.com/autonomy/conform/conform/utilities"
)

// Metadata contains metadata.
type Metadata struct {
	Repository string `yaml:"repository"`
	Docker     *Docker
	Git        *Git
	Version    *Version
	Built      string
}

type Docker struct {
	Image string
}

type Git struct {
	Branch   string
	Message  string
	SHA      string
	Tag      string
	IsBranch bool
	IsClean  bool
	IsTag    bool
}

type Version struct {
	Major        string
	Minor        string
	Prerelease   string
	IsPrerelease bool
}

// UnmarshalYAML implements the yaml.UnmarshalYAML interface.
func (m *Metadata) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var aux struct {
		Repository string `yaml:"repository"`
	}
	if err := unmarshal(&aux); err != nil {
		return err
	}
	m.Repository = aux.Repository
	m.Built = time.Now().Format(time.RFC3339)
	if err := addMetadataForGit(m); err != nil {
		return err
	}
	if err := addMetadataForDocker(m); err != nil {
		return err
	}

	fmt.Printf("%#v", m.Git)
	return nil
}

func addMetadataForDocker(m *Metadata) error {
	image, err := utilities.ImageName(m.Repository, m.Git.SHA, m.Git.IsClean)
	if err != nil {
		return err
	}
	dockerMetadata := &Docker{
		Image: image,
	}
	m.Docker = dockerMetadata

	return nil
}

func addMetadataForGit(m *Metadata) error {
	g, err := git.NewGit()
	if err != nil {
		return err
	}
	branch, isBranch, err := g.Branch()
	if err != nil {
		return err
	}
	message, err := g.Message()
	if err != nil {
		return err
	}
	sha, err := g.SHA()
	if err != nil {
		return err
	}
	_, isClean, err := g.Status()
	if err != nil {
		return err
	}
	tag, isTag, err := g.Tag()
	if err != nil {
		return err
	}
	gitMetadata := &Git{
		Branch:   branch,
		Message:  message,
		SHA:      sha,
		Tag:      tag,
		IsBranch: isBranch,
		IsClean:  isClean,
		IsTag:    isTag,
	}
	m.Git = gitMetadata

	return nil
}
