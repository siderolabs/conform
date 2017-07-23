package conventionalcommit

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/autonomy/conform/pkg/metadata"
	"github.com/autonomy/conform/pkg/pipeline"
	"github.com/autonomy/conform/pkg/policy"
	"github.com/autonomy/conform/pkg/task"
)

// Conventional implements the policy.Policy interface and enforces commit
// messages to conform the Conventional Commit standard.
type Conventional struct {
	Types  []string `mapstructure:"types"`
	Scopes []string `mapstructure:"scopes"`
}

// MaxNumberOfCommitCharacters is the maximium number of characters allowed in
// a commit header.
const MaxNumberOfCommitCharacters = 72

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
	report = policy.Report{}
	if !metadata.Git.IsClean {
		return
	}
	groups := parseHeader(metadata.Git.Message)
	if len(groups) != 5 {
		report.Errors = append(report.Errors, fmt.Errorf("Invalid commit format"))
		return
	}
	ValidateHeaderLength(&report, groups)
	ValidateType(&report, groups, c.Types)
	ValidateScope(&report, groups, c.Scopes)
	ValidateDescription(&report, groups)

	return
}

// Pipeline implements the policy.Policy.Pipeline function.
func (c *Conventional) Pipeline(*pipeline.Pipeline) policy.Option {
	return func(args *policy.Options) {}
}

// Tasks implements the policy.Policy.Tasks function.
func (c *Conventional) Tasks(map[string]*task.Task) policy.Option {
	return func(args *policy.Options) {}
}

// ValidateHeaderLength checks the header length.
func ValidateHeaderLength(report *policy.Report, groups []string) {
	if len(groups[1]) > MaxNumberOfCommitCharacters {
		report.Errors = append(report.Errors, fmt.Errorf("Commit header is %d characters", len(groups[1])))
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
	report.Errors = append(report.Errors, fmt.Errorf("Invalid type: %s", groups[1]))
}

// ValidateScope returns the commit scope.
func ValidateScope(report *policy.Report, groups []string, scopes []string) {
	for _, scope := range scopes {
		if scope == groups[2] {
			return
		}
	}
	report.Errors = append(report.Errors, fmt.Errorf("Invalid scope: %s", groups[2]))
}

// ValidateDescription returns the commit description.
func ValidateDescription(report *policy.Report, groups []string) {
	if len(groups[3]) <= 72 && len(groups[3]) != 0 {
		return
	}
	report.Errors = append(report.Errors, fmt.Errorf("Invalid description: %s", groups[3]))
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
