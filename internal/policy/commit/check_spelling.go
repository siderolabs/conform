/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"fmt"
	"strings"

	"github.com/talos-systems/conform/internal/policy"

	"github.com/golangci/misspell"
)

// SpellCheck represents to spell check policy.
type SpellCheck struct {
	Locale string `mapstructure:"locale"`
}

// SpellingCheck enforces correct spelling.
type SpellingCheck struct {
	errors []error
}

// Name returns the name of the check.
func (h SpellingCheck) Name() string {
	return "Spellcheck"
}

// Message returns to check message.
func (h SpellingCheck) Message() string {
	return fmt.Sprintf("Commit contains %d misspellings", len(h.errors))
}

// Errors returns any violations of the check.
func (h SpellingCheck) Errors() []error {
	return h.errors
}

// ValidateSpelling checks the spelling.
func (c Commit) ValidateSpelling() policy.Check {
	check := &SpellingCheck{}

	r := misspell.Replacer{
		Replacements: misspell.DictMain,
	}

	switch strings.ToUpper(c.SpellCheck.Locale) {
	case "":
	case "US":
		r.AddRuleList(misspell.DictAmerican)
	case "UK", "GB":
		r.AddRuleList(misspell.DictBritish)
	case "NZ", "AU", "CA":
		check.errors = append(check.errors, fmt.Errorf("unknown locale: %q", c.SpellCheck.Locale))
	}

	r.Compile()

	_, diffs := r.Replace(c.msg)

	for _, diff := range diffs {
		check.errors = append(check.errors, fmt.Errorf("`%s` is a misspelling of `%s`", diff.Original, diff.Corrected))
	}

	return check
}
