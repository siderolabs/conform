/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"strings"

	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
)

// RequiredBodyThreshold is the default minimum number of line changes required
// to trigger the body check.
var RequiredBodyThreshold = 10

// Body enforces a maximum number of charcters on the commit
// header.
type Body struct {
	errors []error
}

// Name returns the name of the check.
func (h Body) Name() string {
	return "Commit Body"
}

// Message returns to check message.
func (h Body) Message() string {
	if len(h.errors) != 0 {
		return h.errors[0].Error()
	}
	return "Commit body is valid"
}

// Errors returns any violations of the check.
func (h Body) Errors() []error {
	return h.errors
}

// ValidateBody checks the header length.
func (c Commit) ValidateBody() policy.Check {
	check := &Body{}

	if c.HeaderLength != 0 {
		MaxNumberOfCommitCharacters = c.HeaderLength
	}

	lines := strings.Split(strings.TrimPrefix(c.msg, "\n"), "\n")
	valid := false
	for _, line := range lines[1:] {
		if DCORegex.MatchString(strings.TrimSpace(line)) {
			continue
		}
		if line != "" {
			valid = true
			break
		}
	}

	if !valid {
		check.errors = append(check.errors, errors.New("Commit body is empty"))
	}

	return check
}
