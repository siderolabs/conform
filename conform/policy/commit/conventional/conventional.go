package conventional

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/policy"
)

// Conventional implements the policy.Policy interface and enforces commit
// messages to conform the Conventional Commit standard.
type Conventional struct {
	Types  []string
	Scopes []string
}

// HeaderRegex is the regular expression used for Conventional Commits
// 1.0.0-beta.1.
const HeaderRegex = `^(\w*)\(([^)]+)\):\s{1}(.*)($|\n{2})`

// TypeFeat is a commit of the type fix patches a bug in your codebase
// (this correlates with PATCH in semantic versioning).
const TypeFeat = "feat"

// TypeFix is a commit of the type feat introduces a new feature to the
// codebase (this correlates with MINOR in semantic versioning).
const TypeFix = "fix"

// Compliance implements the policy.Policy.Compliance function.
func (c *Conventional) Compliance(metadata *metadata.Metadata, options ...policy.Option) (report policy.Report) {
	report = policy.Report{Valid: false}
	if !metadata.Git.IsClean {
		report.Valid = true

		return
	}
	groups := parseHeader(metadata.Git.Message)
	if len(groups) != 5 {
		report.Errors = append(report.Errors, fmt.Errorf("Invalid commit format"))
		return
	}
	if err := ValidateType(groups, c.Types); err != nil {
		report.Errors = append(report.Errors, err)
	}

	if err := ValidateScope(groups, c.Scopes); err != nil {
		report.Errors = append(report.Errors, err)
	}

	if err := ValidateDescription(groups); err != nil {
		report.Errors = append(report.Errors, err)
	}
	if len(report.Errors) == 0 {
		report.Valid = true
	}

	return
}

// ValidateType returns the commit type.
func ValidateType(groups []string, types []string) error {
	types = append(types, TypeFeat, TypeFix)
	for _, t := range types {
		if t == groups[1] {
			return nil
		}
	}

	return fmt.Errorf("Invalid type: %s", groups[1])
}

// ValidateScope returns the commit scope.
func ValidateScope(groups []string, scopes []string) error {
	for _, scope := range scopes {
		if scope == groups[2] {
			return nil
		}
	}

	return fmt.Errorf("Invalid scope: %s", groups[2])
}

// ValidateDescription returns the commit description.
func ValidateDescription(groups []string) error {
	if len(groups[3]) <= 72 && len(groups[3]) != 0 {
		return nil
	}

	return fmt.Errorf("Invalid description: %s", groups[3])
}

func parseHeader(message string) []string {
	re, err := regexp.Compile(HeaderRegex)
	if err != nil {
		return nil
	}
	header := strings.Split(message, "\n")[0]
	groups := re.FindStringSubmatch(header)

	return groups
}
