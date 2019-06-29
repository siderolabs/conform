/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"regexp"
	"strings"

	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
)

// DCORegex is the regular expression used for Developer Certificate of Origin.
var DCORegex = regexp.MustCompile(`^Signed-off-by: ([^<]+) <([^<>@]+@[^<>]+)>$`)

// DCOCheck ensures that the commit message contains a
// Developer Certificate of Origin.
type DCOCheck struct {
	errors []error
}

// Name returns the name of the check.
func (d DCOCheck) Name() string {
	return "DCO"
}

// Message returns to check message.
func (d DCOCheck) Message() string {
	if len(d.errors) != 0 {
		return d.errors[0].Error()
	}
	return "Developer Certificate of Origin was found"
}

// Errors returns any violations of the check.
func (d DCOCheck) Errors() []error {
	return d.errors
}

// ValidateDCO checks the commit message for a Developer Certificate of Origin.
func (c Commit) ValidateDCO() policy.Check {
	check := &DCOCheck{}
	for _, line := range strings.Split(c.msg, "\n") {
		if DCORegex.MatchString(strings.TrimSpace(line)) {
			return check
		}
	}

	check.errors = append(check.errors, errors.Errorf("Commit does not have a DCO"))

	return check
}
