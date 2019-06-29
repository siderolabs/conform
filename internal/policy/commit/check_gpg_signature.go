/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"github.com/autonomy/conform/internal/git"
	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
)

// GPGCheck ensures that the commit is cryptographically signed using GPG.
type GPGCheck struct {
	errors []error
}

// Name returns the name of the check.
func (g GPGCheck) Name() string {
	return "GPG"
}

// Message returns to check message.
func (g GPGCheck) Message() string {
	if len(g.errors) != 0 {
		return g.errors[0].Error()
	}
	return "GPG signature found"
}

// Errors returns any violations of the check.
func (g GPGCheck) Errors() []error {
	return g.errors
}

// ValidateGPGSign checks the commit message for a GPG signature.
func (c Commit) ValidateGPGSign(g *git.Git) policy.Check {
	check := &GPGCheck{}

	ok, err := g.HasGPGSignature()
	if err != nil {
		check.errors = append(check.errors, err)
		return check
	}

	if ok {
		return check
	}

	check.errors = append(check.errors, errors.Errorf("Commit does not have a GPG signature"))

	return check
}
