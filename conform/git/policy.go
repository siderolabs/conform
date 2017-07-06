package git

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/autonomy/conform/conform/policy"
)

// ConventionalCommitsOneDotZeroDotZeroBeta1 is the regular expression used for Conventional Commits 1.0.0-beta.1.
const ConventionalCommitsOneDotZeroDotZeroBeta1 = `^(\w*)\(([^)]+)\):\s{1}(.*)($|\n{2})`

// TypeFeat is a commit of the type fix patches a bug in your codebase (this correlates with PATCH in semantic versioning).
const TypeFeat = "feat"

// TypeFix is a commit of the type feat introduces a new feature to the codebase (this correlates with MINOR in semantic versioning).
const TypeFix = "fix"

// ConventionalCommitsOptions are the configurable options used to check the copliance of conventional commits policy.
type ConventionalCommitsOptions struct {
	Types  []string
	Scopes []string
}

// Compliance implements the policy.Policy interface.
func (g *Git) Compliance(obj interface{}) (report *policy.Report, err error) {
	opts := obj.(*ConventionalCommitsOptions)
	report = &policy.Report{Valid: true}
	re, err := regexp.Compile(ConventionalCommitsOneDotZeroDotZeroBeta1)
	if err != nil {
		return
	}
	message, err := g.Message()
	if err != nil {
		return
	}
	lines := strings.Split(message, "\n")
	groups := re.FindStringSubmatch(lines[0])
	if len(groups) != 5 {
		err = fmt.Errorf("Invalid commit format")
		return
	}
	validType, err := ValidateType(groups, opts.Types)
	if !validType {
		report.Valid = false
		report.Errors = append(report.Errors, err)
	}
	validScope, err := ValidateScope(groups, opts.Scopes)
	if !validScope {
		report.Valid = false
		report.Errors = append(report.Errors, err)
	}
	validDescription, err := ValidateDescription(groups)
	if !validDescription {
		report.Valid = false
		report.Errors = append(report.Errors, err)
	}

	return
}

// ValidateType returns the commit type.
func ValidateType(groups []string, types []string) (valid bool, err error) {
	types = append(types, TypeFeat, TypeFix)
	for _, t := range types {
		if t == groups[1] {
			valid = true
		}
	}
	if !valid {
		err = fmt.Errorf("Invalid type: %s", groups[1])
	}

	return
}

// ValidateScope returns the commit scope.
func ValidateScope(groups []string, scopes []string) (valid bool, err error) {
	for _, scope := range scopes {
		if scope == groups[2] {
			valid = true
		}
	}
	if !valid {
		err = fmt.Errorf("Invalid scope: %s", groups[2])
	}

	return
}

// ValidateDescription returns the commit description.
func ValidateDescription(groups []string) (valid bool, err error) {
	if groups[3] != "" {
		valid = true
	}

	return
}
