// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package commit

import (
	"regexp"

	"github.com/pkg/errors"

	"github.com/talos-systems/conform/internal/policy"
)

// JiraCheck enforces that a Jira issue is mentioned in the header.
type JiraCheck struct {
	errors []error
}

// Name returns the name of the check.
func (j *JiraCheck) Name() string {
	return "Jira issues"
}

// Message returns to check message.
func (j *JiraCheck) Message() string {
	if len(j.errors) != 0 {
		return j.errors[0].Error()
	}

	return "Jira issues are valid"
}

// Errors returns any violations of the check.
func (j *JiraCheck) Errors() []error {
	return j.errors
}

// ValidateJiraCheck validates if a Jira issue is mentioned in the header.
func (c Commit) ValidateJiraCheck() policy.Check {
	check := &JiraCheck{}

	reg := regexp.MustCompile(`.* \[?([A-Z]*)-[1-9]{1}\d*\]?.*`)

	if reg.MatchString(c.msg) {
		submatch := reg.FindStringSubmatch(c.msg)
		jiraProject := submatch[1]

		if !find(c.Header.Jira.Keys, jiraProject) {
			check.errors = append(check.errors, errors.Errorf("Jira project %s is not a valid jira project", jiraProject))
		}
	} else {
		check.errors = append(check.errors, errors.Errorf("No Jira issue tag found in %q", c.msg))
	}

	return check
}

func find(slice []string, value string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}

	return false
}
