/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"unicode"
	"unicode/utf8"

	"github.com/pkg/errors"
	"github.com/talos-systems/conform/internal/policy"
)

// HeaderCaseCheck enforces the case of the first word in the header.
type HeaderCaseCheck struct {
	headerCase string
	errors     []error
}

// Name returns the name of the check.
func (h HeaderCaseCheck) Name() string {
	return "Header Case"
}

// Message returns to check message.
func (h HeaderCaseCheck) Message() string {
	if len(h.errors) != 0 {
		return h.errors[0].Error()
	}
	return "Header case is valid"
}

// Errors returns any violations of the check.
func (h HeaderCaseCheck) Errors() []error {
	return h.errors
}

// ValidateHeaderCase checks the header length.
func (c Commit) ValidateHeaderCase() policy.Check {
	check := &HeaderCaseCheck{headerCase: c.Header.Case}

	firstWord, err := c.firstWord()
	if err != nil {
		check.errors = append(check.errors, err)
		return check
	}

	first, _ := utf8.DecodeRuneInString(firstWord)
	if first == utf8.RuneError {
		check.errors = append(check.errors, errors.New("Header does not start with valid UTF-8 text"))
		return check
	}

	var valid bool
	switch c.Header.Case {
	case "upper":
		valid = unicode.IsUpper(first)
	case "lower":
		valid = unicode.IsLower(first)
	default:
		check.errors = append(check.errors, errors.Errorf("Invalid configured case %s", c.Header.Case))
		return check
	}
	if !valid {
		check.errors = append(check.errors, errors.Errorf("Commit header case is not %s", c.Header.Case))
	}
	return check
}
