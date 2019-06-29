/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"fmt"
	"strings"

	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
)

// MaxNumberOfCommitCharacters is the default maximium number of characters
// allowed in a commit header.
var MaxNumberOfCommitCharacters = 89

// HeaderLengthCheck enforces a maximum number of charcters on the commit
// header.
type HeaderLengthCheck struct {
	headerLength int
	errors       []error
}

// Name returns the name of the check.
func (h HeaderLengthCheck) Name() string {
	return "Header Length"
}

// Message returns to check message.
func (h HeaderLengthCheck) Message() string {
	return fmt.Sprintf("Header is %d characters", h.headerLength)
}

// Errors returns any violations of the check.
func (h HeaderLengthCheck) Errors() []error {
	return h.errors
}

// ValidateHeaderLength checks the header length.
func (c Commit) ValidateHeaderLength() policy.Check {
	check := &HeaderLengthCheck{}

	if c.HeaderLength != 0 {
		MaxNumberOfCommitCharacters = c.HeaderLength
	}

	header := strings.Split(strings.TrimPrefix(c.msg, "\n"), "\n")[0]
	check.headerLength = len(header)
	if check.headerLength > MaxNumberOfCommitCharacters {
		check.errors = append(check.errors, errors.Errorf("Commit header is %d characters", len(header)))
	}

	return check
}
