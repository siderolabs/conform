//go:build !some_test_tag
// +build !some_test_tag

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package license_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/talos-systems/conform/internal/policy/license"
)

func TestLicense(t *testing.T) {
	const header = `
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.`

	t.Run("Default", func(t *testing.T) {
		l := license.License{
			IncludeSuffixes:        []string{".go"},
			AllowPrecedingComments: false,
			Header:                 header,
		}
		check := l.ValidateLicenseHeader()
		assert.Equal(t, "Found 1 files without license header", check.Message())
	})

	t.Run("AllowPrecedingComments", func(t *testing.T) {
		l := license.License{
			IncludeSuffixes:        []string{".go"},
			AllowPrecedingComments: true,
			Header:                 header,
		}
		check := l.ValidateLicenseHeader()
		assert.Equal(t, "All files have a valid license header", check.Message())
	})
}
