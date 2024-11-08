// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package commit provides commit-related policies.
package commit

import (
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/siderolabs/conform/internal/git"
	"github.com/siderolabs/conform/internal/policy"
)

// HeaderChecks is the configuration for checks on the header of a commit.
//
//nolint:govet
type HeaderChecks struct {
	// Length is the maximum length of the commit subject.
	Length int `mapstructure:"length"`
	// Imperative enforces the use of imperative verbs as the first word of a
	// commit message.
	Imperative bool `mapstructure:"imperative"`
	// HeaderCase is the case that the first word of the header must have ("upper" or "lower").
	Case string `mapstructure:"case"`
	// HeaderInvalidLastCharacters is a string containing all invalid last characters for the header.
	InvalidLastCharacters string `mapstructure:"invalidLastCharacters"`
	// Jira checks if the header containers a Jira project key.
	Jira *JiraChecks `mapstructure:"jira"`
}

// JiraChecks is the configuration for checks for Jira issues.
type JiraChecks struct {
	Keys []string `mapstructure:"keys"`
}

// BodyChecks is the configuration for checks on the body of a commit.
type BodyChecks struct {
	// Required enforces that the current commit has a body.
	Required bool `mapstructure:"required"`
}

// GPG is the configuration for checks GPG signature on the commit.
//
//nolint:govet
type GPG struct {
	// Required enforces that the current commit has a signature.
	Required bool `mapstructure:"required"`
	// Identity configures identity of the signature.
	Identity *struct {
		// GitHubOrganization enforces that commit should be signed with the key
		// of one of the organization public members.
		GitHubOrganization string `mapstructure:"gitHubOrganization"`
	} `mapstructure:"identity"`
}

// Commit implements the policy.Policy interface and enforces commit
// messages to conform the Conventional Commit standard.
//
//nolint:maligned,govet
type Commit struct {
	// SpellCheck enforces correct spelling.
	SpellCheck *SpellCheck `mapstructure:"spellcheck"`
	// Conventional is the user specified settings for conventional commits.
	Conventional *Conventional `mapstructure:"conventional"`
	// Header is the user specified settings for the header of each commit.
	Header *HeaderChecks `mapstructure:"header"`
	// Header is the user specified settings for the body of each commit.
	Body *BodyChecks `mapstructure:"body"`
	// DCO enables the Developer Certificate of Origin check.
	DCO bool `mapstructure:"dco"`
	// GPG is the user specified settings for the GPG signature check.
	GPG *GPG `mapstructure:"gpg"`
	// GPGSignatureGitHubOrganization enforces that GPG signature should come from
	// one of the members of the GitHub org.
	GPGSignatureGitHubOrganization string `mapstructure:"gpgSignatureGitHubOrg"`
	// MaximumOfOneCommit enforces that the current commit is only one commit
	// ahead of a specified ref.
	MaximumOfOneCommit bool `mapstructure:"maximumOfOneCommit"`

	msg string
	sha string
}

// FirstWordRegex is theregular expression used to find the first word in a
// commit.
var FirstWordRegex = regexp.MustCompile(`^\s*([a-zA-Z0-9]+)`)

// Compliance implements the policy.Policy.Compliance function.
func (c *Commit) Compliance(options *policy.Options) (*policy.Report, error) {
	var err error

	report := &policy.Report{}

	// Setup the policy for all checks.
	var g *git.Git

	if g, err = git.NewGit(); err != nil {
		return report, errors.Errorf("failed to open git repo: %v", err)
	}

	var commits []git.CommitData

	switch o := options; {
	case o.CommitMsgFile != nil:
		var contents []byte

		if contents, err = os.ReadFile(*options.CommitMsgFile); err != nil {
			return report, errors.Errorf("failed to read commit message file: %v", err)
		}

		commits = append(commits, git.CommitData{Message: string(contents)})
	case o.RevisionRange != "":
		revs, err := extractRevisionRange(options)
		if err != nil {
			return report, errors.Errorf("failed to get commit message: %v", err)
		}

		commits, err = g.Commits(revs[0], revs[1])
		if err != nil {
			return report, errors.Errorf("failed to get commit message: %v", err)
		}
	default:
		commit, err := g.Commit()
		if err != nil {
			return report, errors.Errorf("failed to get commit message: %v", err)
		}

		commits = append(commits, commit)
	}

	for i := range commits {
		c.msg = commits[i].Message
		c.sha = commits[i].SHA

		c.compliance(report, g, options)
	}

	return report, nil
}

// compliance checks the compliance with the policies of the given commit.
func (c *Commit) compliance(report *policy.Report, g *git.Git, options *policy.Options) {
	if c.Header != nil {
		if c.Header.Length != 0 {
			report.AddCheck(c.ValidateHeaderLength())
		}

		if c.Header.Imperative {
			report.AddCheck(c.ValidateImperative())
		}

		if c.Header.Case != "" {
			report.AddCheck(c.ValidateHeaderCase())
		}

		if c.Header.InvalidLastCharacters != "" {
			report.AddCheck(c.ValidateHeaderLastCharacter())
		}

		if c.Header.Jira != nil {
			report.AddCheck(c.ValidateJiraCheck())
		}
	}

	if c.DCO {
		report.AddCheck(c.ValidateDCO())
	}

	if c.GPG != nil {
		if c.GPG.Required {
			report.AddCheck(c.ValidateGPGSign(g))

			if c.GPG.Identity != nil {
				report.AddCheck(c.ValidateGPGIdentity(g))
			}
		}
	}

	if c.Conventional != nil {
		report.AddCheck(c.ValidateConventionalCommit())
	}

	if c.SpellCheck != nil {
		report.AddCheck(c.ValidateSpelling())
	}

	if c.MaximumOfOneCommit {
		report.AddCheck(c.ValidateNumberOfCommits(g, options.CommitRef))
	}

	if c.Body != nil {
		if c.Body.Required {
			report.AddCheck(c.ValidateBody())
		}
	}
}

func (c Commit) firstWord() (string, error) {
	var (
		groups []string
		msg    string
	)

	if c.Conventional != nil {
		groups = parseHeader(c.msg)
		if len(groups) != 7 {
			return "", errors.Errorf("Invalid conventional commit format")
		}

		msg = groups[5]
	} else {
		msg = c.msg
	}

	if msg == "" {
		return "", errors.Errorf("Invalid msg: %s", msg)
	}

	if groups = FirstWordRegex.FindStringSubmatch(msg); groups == nil {
		return "", errors.Errorf("Invalid msg: %s", msg)
	}

	return groups[0], nil
}

func (c Commit) header() string {
	return strings.Split(strings.TrimPrefix(c.msg, "\n"), "\n")[0]
}

func extractRevisionRange(options *policy.Options) ([]string, error) {
	revs := strings.Split(options.RevisionRange, "..")
	if len(revs) > 2 || len(revs) == 0 || revs[0] == "" || revs[1] == "" {
		return nil, errors.New("invalid revision range")
	} else if len(revs) == 1 {
		// if no final rev is given, use HEAD as default
		revs = append(revs, "HEAD")
	}

	return revs, nil
}
