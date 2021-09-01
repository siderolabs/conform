// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package commit

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/talos-systems/conform/internal/policy"
)

// MaxNumberOfCommitCharacters is the default maximium number of characters
// allowed in a commit header.
var MaxNumberOfCommitCharacters = 89

// HeaderLengthCheck enforces a maximum number of charcters on the commit
// header.
//
//nolint:govet
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

	if c.Header.Length != 0 {
		MaxNumberOfCommitCharacters = c.Header.Length
	}

	check.headerLength = len(c.header())
	if check.headerLength > MaxNumberOfCommitCharacters {
		check.errors = append(check.errors, errors.Errorf("Commit header is %d characters", check.headerLength))
	}

	return check
}
