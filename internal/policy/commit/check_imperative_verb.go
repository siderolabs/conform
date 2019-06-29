/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"strings"

	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
	"gopkg.in/jdkato/prose.v2"
)

// ImperativeCheck enforces that the first word of a commit message header is
// and imperative verb.
type ImperativeCheck struct {
	errors []error
}

// Name returns the name of the check.
func (i ImperativeCheck) Name() string {
	return "Imperative Mood"
}

// Message returns to check message.
func (i ImperativeCheck) Message() string {
	if len(i.errors) != 0 {
		return i.errors[0].Error()
	}
	return "Commit begins with imperative verb"
}

// Errors returns any violations of the check.
func (i ImperativeCheck) Errors() []error {
	return i.errors
}

// ValidateImperative checks the commit message for a GPG signature.
func (c Commit) ValidateImperative() policy.Check {
	check := &ImperativeCheck{}
	var (
		word string
		err  error
	)
	if word, err = c.firstWord(); err != nil {
		check.errors = append(check.errors, err)
		return check
	}
	doc, err := prose.NewDocument("I " + strings.ToLower(word))
	if err != nil {
		check.errors = append(check.errors, errors.Errorf("Failed to create document: %v", err))
		return check
	}
	if len(doc.Tokens()) != 2 {
		check.errors = append(check.errors, errors.Errorf("Expected 2 tokens, got %d", len(doc.Tokens())))
		return check
	}
	tokens := doc.Tokens()
	tok := tokens[1]
	for _, tag := range []string{"VBD", "VBG", "VBZ"} {
		if tok.Tag == tag {
			check.errors = append(check.errors, errors.Errorf("First word of commit must be an imperative verb: %q is invalid", word))
		}
	}

	return check
}
