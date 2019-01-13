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
}

// MaxNumberOfCommitCharacters is the default maximium number of characters
// allowed in a commit header.
var MaxNumberOfCommitCharacters = 89

// DCORegex is the regular expression used for Developer Certificate of Origin.
var DCORegex = regexp.MustCompile(`^Signed-off-by: ([^<]+) <([^<>@]+@[^<>]+)>$`)

// Compliance implements the policy.Policy.Compliance function.
func (c *Commit) Compliance(options *policy.Options) (report policy.Report) {
	var err error

	report = policy.Report{}

	var g *git.Git
	if g, err = git.NewGit(); err != nil {
		report.Errors = append(report.Errors, errors.Errorf("failed to open git repo: %v", err))
		return
	}

	var msg string
	if options.CommitMsgFile != nil {
		var contents []byte
		if contents, err = ioutil.ReadFile(*options.CommitMsgFile); err != nil {
			report.Errors = append(report.Errors, errors.Errorf("failed to read commit message file: %v", err))
			return
		}
		msg = string(contents)
	} else {
		if msg, err = g.Message(); err != nil {
			report.Errors = append(report.Errors, errors.Errorf("failed to get commit message: %v", err))
			return
		}
	}

	if c.HeaderLength != 0 {
		MaxNumberOfCommitCharacters = c.HeaderLength
	}
	ValidateHeaderLength(&report, msg)

	if c.DCO {
		ValidateDCO(&report, msg)
	}

	if c.GPG {
		ValidateGPGSign(&report, g)
	}

	return report
}

// ValidateHeaderLength checks the header length.
func ValidateHeaderLength(report *policy.Report, msg string) {
	header := strings.Split(strings.TrimPrefix(msg, "\n"), "\n")[0]
	if len(header) > MaxNumberOfCommitCharacters {
		report.Errors = append(report.Errors, errors.Errorf("Commit header is %d characters", len(header)))
	}
}

// ValidateDCO checks the commit message for a Developer Certificate of Origin.
func ValidateDCO(report *policy.Report, msg string) {
	for _, line := range strings.Split(msg, "\n") {
		if DCORegex.MatchString(line) {
			return
		}
	}

	report.Errors = append(report.Errors, errors.Errorf("Commit does not have a DCO"))
}

// ValidateGPGSign checks the commit message for a GPG signature.
func ValidateGPGSign(report *policy.Report, g *git.Git) {
	var err error
	var ok bool
	if ok, err = g.HasGPGSignature(); !ok {
		if err != nil {
			report.Errors = append(report.Errors, errors.Errorf("Commit does not have a GPG signature: %v", err))
		}
		report.Errors = append(report.Errors, errors.Errorf("Commit does not have a GPG signature"))
	}
}
