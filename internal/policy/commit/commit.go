/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/autonomy/conform/internal/git"
	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
)

// Commit implements the policy.Policy interface and enforces commit
// messages to conform the Conventional Commit standard.
type Commit struct {
	// HeaderLength is the maximum length of the commit subject.
	HeaderLength int `mapstructure:"headerLength"`
	// DCO enables the Developer Certificate of Origin check.
	DCO bool `mapstructure:"dco"`
	// GPG enables the GPG signature check.
	GPG bool `mapstructure:"gpg"`
	// Imperative enforces the use of imperative verbs as the first word of a
	// commit message.
	Imperative bool `mapstructure:"imperative"`
	// MaximumOfOneCommit enforces that the current commit is only one commit
	// ahead of a specified ref.
	MaximumOfOneCommit bool `mapstructure:"maximumOfOneCommit"`
	// RequireCommitBody enforces that the current commit has a body.
	RequireCommitBody bool `mapstructure:"requireCommitBody"`
	// Conventional is the user specified settings for conventional commits.
	Conventional *Conventional `mapstructure:"conventional"`

	msg string
}

// FirstWordRegex is theregular expression used to find the first word in a
// commit.
var FirstWordRegex = regexp.MustCompile(`^\s*([a-zA-Z0-9]+)`)

// Compliance implements the policy.Policy.Compliance function.
// nolint: gocyclo
func (c *Commit) Compliance(options *policy.Options) (*policy.Report, error) {
	var err error

	report := &policy.Report{}

	// Setup the policy for all checks.

	var g *git.Git
	if g, err = git.NewGit(); err != nil {
		return report, errors.Errorf("failed to open git repo: %v", err)
	}

	var msg string
	if options.CommitMsgFile != nil {
		var contents []byte
		if contents, err = ioutil.ReadFile(*options.CommitMsgFile); err != nil {
			return report, errors.Errorf("failed to read commit message file: %v", err)
		}
		msg = string(contents)
	} else if msg, err = g.Message(); err != nil {
		return report, errors.Errorf("failed to get commit message: %v", err)
	}
	c.msg = msg

	if c.HeaderLength != 0 {
		report.AddCheck(c.ValidateHeaderLength())
	}

	if c.DCO {
		report.AddCheck(c.ValidateDCO())
	}

	if c.GPG {
		report.AddCheck(c.ValidateGPGSign(g))
	}

	if c.Imperative {
		report.AddCheck(c.ValidateImperative())
	}

	if c.Conventional != nil {
		report.AddCheck(c.ValidateConventionalCommit())
	}

	if c.MaximumOfOneCommit {
		report.AddCheck(c.ValidateNumberOfCommits(g, "refs/heads/master"))
	}

	if c.RequireCommitBody {
		report.AddCheck(c.ValidateBody())
	}

	return report, nil
}

func (c Commit) firstWord() (string, error) {
	var header string
	var groups []string
	var msg string
	if c.Conventional != nil {
		groups = parseHeader(c.msg)
		msg = groups[4]
	} else {
		msg = c.msg
	}
	if header = strings.Split(strings.TrimPrefix(msg, "\n"), "\n")[0]; header == "" {
		return "", errors.Errorf("Invalid msg: %s", msg)
	}
	if groups = FirstWordRegex.FindStringSubmatch(header); groups == nil {
		return "", errors.Errorf("Invalid msg: %s", msg)
	}
	return groups[0], nil
}
