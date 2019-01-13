/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package conventionalcommit

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/autonomy/conform/internal/git"
	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
)

// Conventional implements the policy.Policy interface and enforces commit
// messages to conform the Conventional Commit standard.
type Conventional struct {
	Types  []string `mapstructure:"types"`
	Scopes []string `mapstructure:"scopes"`
}

// HeaderRegex is the regular expression used for Conventional Commits
// 1.0.0-beta.1.
var HeaderRegex = regexp.MustCompile(`^(\w*)(\(([^)]+)\))?:\s{1}(.*)($|\n{2})`)

// TypeFeat is a commit of the type fix patches a bug in your codebase
// (this correlates with PATCH in semantic versioning).
const TypeFeat = "feat"

// TypeFix is a commit of the type feat introduces a new feature to the
// codebase (this correlates with MINOR in semantic versioning).
const TypeFix = "fix"

// Compliance implements the policy.Policy.Compliance function.
func (c *Conventional) Compliance(options *policy.Options) (report policy.Report) {
	report = policy.Report{}

	var msg string
	if options.CommitMsgFile != nil {
		contents, err := ioutil.ReadFile(*options.CommitMsgFile)
		if err != nil {
			report.Errors = append(report.Errors, errors.Errorf("failed to read commit message file: %v", err))
			return
		}
		msg = string(contents)
	} else {
		g, err := git.NewGit()
		if err != nil {
			report.Errors = append(report.Errors, errors.Errorf("failed to open git repo: %v", err))
			return
		}
		if msg, err = g.Message(); err != nil {
			report.Errors = append(report.Errors, errors.Errorf("failed to get commit message: %v", err))
			return
		}
	}
	groups := parseHeader(msg)
	if len(groups) != 6 {
		report.Errors = append(report.Errors, errors.Errorf("Invalid commit format: %s", msg))
		return
	}

	ValidateType(&report, groups, c.Types)
	ValidateScope(&report, groups, c.Scopes)
	ValidateDescription(&report, groups)

	return report
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

func parseHeader(msg string) []string {
	// To circumvent any policy violation due to the leading \n that GitHub
	// prefixes to the commit message on a squash merge, we remove it from the
	// message.
	header := strings.Split(strings.TrimPrefix(msg, "\n"), "\n")[0]
	groups := HeaderRegex.FindStringSubmatch(header)

	return groups
}
