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
	prose "gopkg.in/jdkato/prose.v2"
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
	// Conventional is the user specified settings for conventional commits.
	Conventional *Conventional `mapstructure:"conventional"`
}

// Conventional implements the policy.Policy interface and enforces commit
// messages to conform the Conventional Commit standard.
type Conventional struct {
	Types  []string `mapstructure:"types"`
	Scopes []string `mapstructure:"scopes"`
}

// MaxNumberOfCommitCharacters is the default maximium number of characters
// allowed in a commit header.
var MaxNumberOfCommitCharacters = 89

// DCORegex is the regular expression used for Developer Certificate of Origin.
var DCORegex = regexp.MustCompile(`^Signed-off-by: ([^<]+) <([^<>@]+@[^<>]+)>$`)

// FirstWordRegex is theregular expression used to find the first word in a
// commit.
var FirstWordRegex = regexp.MustCompile(`^\s*([a-zA-Z0-9]+)`)

// HeaderRegex is the regular expression used for Conventional Commits
// 1.0.0-beta.1.
var HeaderRegex = regexp.MustCompile(`^(\w*)(\(([^)]+)\))?:\s{1}(.*)($|\n{2})`)

const (
	// TypeFeat is a commit of the type fix patches a bug in your codebase
	// (this correlates with PATCH in semantic versioning).
	TypeFeat = "feat"

	// TypeFix is a commit of the type feat introduces a new feature to the
	// codebase (this correlates with MINOR in semantic versioning).
	TypeFix = "fix"
)

// Compliance implements the policy.Policy.Compliance function.
// nolint: gocyclo
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

	var word string
	if word, err = firstWord(msg, &report); err != nil {
		return
	}

	if c.Conventional != nil {
		groups := parseHeader(msg)
		if len(groups) != 6 {
			report.Errors = append(report.Errors, errors.Errorf("Invalid conventional commits format: %s", msg))
			return
		}
		if word, err = firstWord(groups[4], &report); err != nil {
			return
		}

		ValidateType(&report, groups, c.Conventional.Types)
		ValidateScope(&report, groups, c.Conventional.Scopes)
		ValidateDescription(&report, groups)
	}

	if c.Imperative {
		ValidateImperative(&report, word)
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

// ValidateImperative checks the commit message for a GPG signature.
func ValidateImperative(report *policy.Report, word string) {
	doc, err := prose.NewDocument("I " + strings.ToLower(word))
	if err != nil {
		report.Errors = append(report.Errors, errors.Errorf("Failed to create document: %v", err))
	}
	if len(doc.Tokens()) != 2 {
		report.Errors = append(report.Errors, errors.Errorf("Expected 2 tokens, got %d", len(doc.Tokens())))
		return
	}
	tokens := doc.Tokens()
	tok := tokens[1]
	for _, tag := range []string{"VBD", "VBG", "VBZ"} {
		if tok.Tag == tag {
			report.Errors = append(report.Errors, errors.Errorf("First word of commit must be an imperative verb: %q", word))
		}
	}
}

// ValidateType returns the commit type.
func ValidateType(report *policy.Report, groups []string, types []string) {
	types = append(types, TypeFeat, TypeFix)
	for _, t := range types {
		if t == groups[1] {
			return
		}
	}
	report.Errors = append(report.Errors, errors.Errorf("Invalid type: %s, allowed types are: %v", groups[1], types))
}

// ValidateScope returns the commit scope.
func ValidateScope(report *policy.Report, groups []string, scopes []string) {
	// Scope is optional.
	if groups[3] == "" {
		return
	}
	for _, scope := range scopes {
		if scope == groups[3] {
			return
		}
	}
	report.Errors = append(report.Errors, errors.Errorf("Invalid scope: %s, allowed scopes are: %v", groups[3], scopes))
}

// ValidateDescription returns the commit description.
func ValidateDescription(report *policy.Report, groups []string) {
	if len(groups[4]) <= 72 && len(groups[4]) != 0 {
		return
	}
	report.Errors = append(report.Errors, errors.Errorf("Invalid description: %s", groups[4]))
}

func firstWord(msg string, report *policy.Report) (string, error) {
	var header string
	var groups []string
	if header = strings.Split(strings.TrimPrefix(msg, "\n"), "\n")[0]; header == "" {
		report.Errors = append(report.Errors, errors.Errorf("Invalid conventional commits (empty)"))
		return "", errors.Errorf("Invalid msg: %s", msg)
	}
	if groups = FirstWordRegex.FindStringSubmatch(header); groups == nil {
		return "", errors.Errorf("Invalid msg: %s", msg)
	}
	return groups[0], nil
}

func parseHeader(msg string) []string {
	// To circumvent any policy violation due to the leading \n that GitHub
	// prefixes to the commit message on a squash merge, we remove it from the
	// message.
	header := strings.Split(strings.TrimPrefix(msg, "\n"), "\n")[0]
	groups := HeaderRegex.FindStringSubmatch(header)

	return groups
}
