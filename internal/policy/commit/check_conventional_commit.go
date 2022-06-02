// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package commit

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/siderolabs/conform/internal/policy"
)

// Conventional implements the policy.Policy interface and enforces commit
// messages to conform the Conventional Commit standard.
type Conventional struct {
	Types             []string `mapstructure:"types"`
	Scopes            []string `mapstructure:"scopes"`
	DescriptionLength int      `mapstructure:"descriptionLength"`
}

// HeaderRegex is the regular expression used for Conventional Commits 1.0.0.
var HeaderRegex = regexp.MustCompile(`^(\w*)(\(([^)]+)\))?(!)?:\s{1}(.*)($|\n{2})`)

const (
	// TypeFeat is a commit of the type fix patches a bug in your codebase
	// (this correlates with PATCH in semantic versioning).
	TypeFeat = "feat"

	// TypeFix is a commit of the type feat introduces a new feature to the
	// codebase (this correlates with MINOR in semantic versioning).
	TypeFix = "fix"
)

// ConventionalCommitCheck ensures that the commit message is a valid
// conventional commit.
type ConventionalCommitCheck struct {
	errors []error
}

// Name returns the name of the check.
func (c ConventionalCommitCheck) Name() string {
	return "Conventional Commit"
}

// Message returns to check message.
func (c ConventionalCommitCheck) Message() string {
	if len(c.errors) != 0 {
		return c.errors[0].Error()
	}

	return "Commit message is a valid conventional commit"
}

// Errors returns any violations of the check.
func (c ConventionalCommitCheck) Errors() []error {
	return c.errors
}

// ValidateConventionalCommit returns the commit type.
func (c Commit) ValidateConventionalCommit() policy.Check { //nolint:ireturn
	check := &ConventionalCommitCheck{}
	groups := parseHeader(c.msg)

	if len(groups) != 7 {
		check.errors = append(check.errors, errors.Errorf("Invalid conventional commits format: %q", c.msg))

		return check
	}

	// conventional commit sections
	ccType := groups[1]
	ccScope := groups[3]
	ccDesc := groups[5]

	c.Conventional.Types = append(c.Conventional.Types, TypeFeat, TypeFix)
	typeIsValid := false

	for _, t := range c.Conventional.Types {
		if t == ccType {
			typeIsValid = true
		}
	}

	if !typeIsValid {
		check.errors = append(check.errors, errors.Errorf("Invalid type %q: allowed types are %v", groups[1], c.Conventional.Types))

		return check
	}

	// Scope is optional.
	if ccScope != "" {
		scopeIsValid := false

		for _, scope := range c.Conventional.Scopes {
			re := regexp.MustCompile(scope)
			if re.Match([]byte(ccScope)) {
				scopeIsValid = true

				break
			}
		}

		if !scopeIsValid {
			check.errors = append(check.errors, errors.Errorf("Invalid scope %q: allowed scopes are %v", groups[3], c.Conventional.Scopes))

			return check
		}
	}

	// Provide a good default value for DescriptionLength
	if c.Conventional.DescriptionLength == 0 {
		c.Conventional.DescriptionLength = 72
	}

	if len(ccDesc) <= c.Conventional.DescriptionLength && len(ccDesc) != 0 {
		return check
	}

	check.errors = append(check.errors, errors.Errorf("Invalid description: %s", ccDesc))

	return check
}

func parseHeader(msg string) []string {
	// To circumvent any policy violation due to the leading \n that GitHub
	// prefixes to the commit message on a squash merge, we remove it from the
	// message.
	header := strings.Split(strings.TrimPrefix(msg, "\n"), "\n")[0]
	groups := HeaderRegex.FindStringSubmatch(header)

	return groups
}
