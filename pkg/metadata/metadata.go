package metadata

import (
	"time"

	"github.com/Masterminds/semver"
	"github.com/autonomy/conform/pkg/git"
)

// Metadata contains metadata.
type Metadata struct {
	Repository string `yaml:"repository"`
	Docker     *Docker
	Git        *Git
	Version    *Version
	Variables  VariablesMap `yaml:"variables"`
	Built      string
}

// Docker contains docker specific metadata.
type Docker struct {
	Image         *Image
	PreviousStage string
	CurrentStage  string
	NextStage     string
}

// Image contains information used to identity an image.
type Image struct {
	Name string
	Tag  string
}

// Git contains git specific metadata.
type Git struct {
	Branch   string
	Message  string
	SHA      string
	Tag      string
	IsBranch bool
	IsClean  bool
	IsTag    bool
}

// Version contains version specific metadata.
type Version struct {
	Major        int64
	Minor        int64
	Patch        int64
	Prerelease   string
	IsPrerelease bool
}

// VariablesMap is a map for user defined metadata.
type VariablesMap = map[string]interface{}

// UnmarshalYAML implements the yaml.UnmarshalYAML interface.
func (m *Metadata) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var aux struct {
		Repository string       `yaml:"repository"`
		Variables  VariablesMap `yaml:"variables"`
	}
	if err := unmarshal(&aux); err != nil {
		return err
	}

	m.Repository = aux.Repository
	m.Variables = aux.Variables
	m.Built = time.Now().UTC().Format(time.RFC1123)

	if err := addMetadataForGit(m); err != nil {
		return err
	}
	if err := addMetadataForDocker(m); err != nil {
		return err
	}
	if err := addMetadataForVersion(m); err != nil {
		return err
	}

	return nil
}

func addMetadataForVersion(m *Metadata) error {
	m.Version = &Version{}
	if m.Git.IsTag {
		var ver *semver.Version
		ver, err := semver.NewVersion(m.Git.Tag[1:])
		if err != nil {
			return err
		}
		m.Version.Major = ver.Major()
		m.Version.Minor = ver.Minor()
		m.Version.Patch = ver.Patch()
		m.Version.Prerelease = ver.Prerelease()
		if ver.Prerelease() != "" {
			m.Version.IsPrerelease = true
		}
	}

	return nil
}

func addMetadataForDocker(m *Metadata) error {
	tag := m.Git.SHA
	name := m.Repository

	dockerMetadata := &Docker{
		Image: &Image{
			Name: name,
			Tag:  tag,
		},
	}
	m.Docker = dockerMetadata

	return nil
}

func addMetadataForGit(m *Metadata) error {
	g, err := git.NewGit()
	if err != nil {
		return err
	}
	m.Git = &Git{}
	if err = addBranchMetadataForGit(g, m); err != nil {
		return err
	}
	if err = addMessageMetadataForGit(g, m); err != nil {
		return err
	}
	if err = addStatusMetadataForGit(g, m); err != nil {
		return err
	}
	if err = addSHAMetadataForGit(g, m); err != nil {
		return err
	}
	if err = addTagMetadataForGit(g, m); err != nil {
		return err
	}

	return nil
}

func addBranchMetadataForGit(g *git.Git, m *Metadata) error {
	branch, isBranch, err := g.Branch()
	if err != nil {
		return err
	}
	m.Git.Branch = branch
	m.Git.IsBranch = isBranch

	return nil
}

func addMessageMetadataForGit(g *git.Git, m *Metadata) error {
	message, err := g.Message()
	if err != nil {
		return err
	}
	m.Git.Message = message

	return nil
}

func addSHAMetadataForGit(g *git.Git, m *Metadata) error {
	sha, err := g.SHA()
	if err != nil {
		return err
	}

	if !m.Git.IsClean {
		sha = sha + "-dirty"
	}

	m.Git.SHA = sha

	return nil
}

func addStatusMetadataForGit(g *git.Git, m *Metadata) error {
	_, isClean, err := g.Status()
	if err != nil {
		return err
	}
	m.Git.IsClean = isClean

	return nil
}

func addTagMetadataForGit(g *git.Git, m *Metadata) error {
	tag, isTag, err := g.Tag()
	if err != nil {
		return err
	}
	m.Git.Tag = tag
	m.Git.IsTag = isTag

	return nil
}
