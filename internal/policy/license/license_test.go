// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//go:build !some_test_tag
// +build !some_test_tag

package license_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/siderolabs/conform/internal/policy/license"
)

func TestLicense(t *testing.T) {
	const header = `
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.`

	const otherHeader = "// some-other-header"

	t.Run("Default", func(t *testing.T) {
		l := license.Licenses{
			{
				SkipPaths:              []string{"subdir1/"},
				IncludeSuffixes:        []string{".txt"},
				AllowPrecedingComments: false,
				Header:                 header,
			},
		}
		check := l.ValidateLicenseHeaders()
		assert.Equal(t, "Found 1 files without license header", check.Message())
	})

	t.Run("AllowPrecedingComments", func(t *testing.T) {
		l := license.Licenses{
			{
				SkipPaths:              []string{"subdir1/"},
				IncludeSuffixes:        []string{".txt"},
				AllowPrecedingComments: true,
				Header:                 header,
			},
		}
		check := l.ValidateLicenseHeaders()
		assert.Equal(t, "All files have a valid license header", check.Message())
	})

	// File "testdata/subdir1/subdir2/data.txt" is valid for the root license, but "testdata/subdir1/" is skipped.
	// It is invalid for the additional license, but that license skips "subdir2/" relative to itself.
	// The check should pass.
	t.Run("AdditionalValid", func(t *testing.T) {
		l := license.Licenses{
			{
				IncludeSuffixes:        []string{".txt"},
				SkipPaths:              []string{"testdata/subdir1/"},
				AllowPrecedingComments: true,
				Header:                 header,
			},
			{
				Root:            "testdata/subdir1/",
				SkipPaths:       []string{"subdir2/"},
				IncludeSuffixes: []string{".txt"},
				Header:          otherHeader,
			},
		}
		check := l.ValidateLicenseHeaders()
		assert.Equal(t, "All files have a valid license header", check.Message())
	})

	// File "testdata/subdir1/subdir2/data.txt" is valid for the root license, but "testdata/subdir1/" is skipped.
	// However, it is invalid for the additional license.
	// The check should fail.
	t.Run("AdditionalInvalid", func(t *testing.T) {
		l := license.Licenses{
			{
				IncludeSuffixes:        []string{".txt"},
				SkipPaths:              []string{"testdata/subdir1/"},
				AllowPrecedingComments: true,
				Header:                 header,
			},

			{
				Root:            "testdata/subdir1/",
				IncludeSuffixes: []string{".txt"},
				Header:          otherHeader,
			},
		}
		check := l.ValidateLicenseHeaders()
		assert.Equal(t, "Found 1 files without license header", check.Message())
	})
}
