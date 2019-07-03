/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package license

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
)

// License implements the policy.Policy interface and enforces source code
// license headers.
type License struct {
	// SkipPaths applies fnmatch-style patterns to file paths to skip completely
	// parts of the tree which shouldn't be scanned (e.g. .git/)
	SkipPaths []string `mapstructure:"skipPaths"`
	// IncludeSuffixes is the regex used to find files that the license policy
	// should be applied to.
	IncludeSuffixes []string `mapstructure:"includeSuffixes"`
	// ExcludeSuffixes is the Suffixes used to find files that the license policy
	// should not be applied to.
	ExcludeSuffixes []string `mapstructure:"excludeSuffixes"`
	// Header is the contents of the license header.
	Header string `mapstructure:"header"`
}

// Compliance implements the policy.Policy.Compliance function.
func (l *License) Compliance(options *policy.Options) (*policy.Report, error) {
	report := &policy.Report{}

	report.AddCheck(l.ValidateLicenseHeader())

	return report, nil
}

// HeaderCheck enforces a license header on source code files.
type HeaderCheck struct {
	errors []error
}

// Name returns the name of the check.
func (l HeaderCheck) Name() string {
	return "File Header"
}

// Message returns to check message.
func (l HeaderCheck) Message() string {
	if len(l.errors) != 0 {
		return fmt.Sprintf("Found %d files without license header", len(l.errors))
	}
	return "All files have a valid license header"
}

// Errors returns any violations of the check.
func (l HeaderCheck) Errors() []error {
	return l.errors
}

// ValidateLicenseHeader checks the header of a file and ensures it contains the
// provided value.
// nolint: gocyclo
func (l License) ValidateLicenseHeader() policy.Check {
	check := HeaderCheck{}
	if l.Header == "" {
		check.errors = append(check.errors, errors.New("Header is not defined"))
		return check
	}
	value := []byte(l.Header)
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		matchPath := path
		if info.IsDir() {
			// for directories, match against "dir/
			matchPath += "/"
		}
		for _, pattern := range l.SkipPaths {
			var matches bool
			matches, err = filepath.Match(pattern, matchPath)
			if err != nil {
				return err
			}
			if matches {
				if info.IsDir() {
					// skip whole directory tree
					return filepath.SkipDir
				}
				// skip single file
				return nil
			}
		}

		if info.Mode().IsRegular() {
			// Skip excluded suffixes.
			for _, suffix := range l.ExcludeSuffixes {
				if strings.HasSuffix(info.Name(), suffix) {
					return nil
				}
			}
			// Check files matching the included suffixes.
			for _, suffix := range l.IncludeSuffixes {
				if strings.HasSuffix(info.Name(), suffix) {
					var contents []byte
					if contents, err = ioutil.ReadFile(path); err != nil {
						check.errors = append(check.errors, errors.Errorf("Failed to open %s", path))
						return nil
					}
					// ValidateLicenseHeader(&report, info.Name(), contents, []byte(l.Header))
					if bytes.HasPrefix(contents, value) {
						continue
					}
					check.errors = append(check.errors, errors.Errorf("File %s does not contain a license header", info.Name()))
				}
			}
		}
		return nil
	})
	if err != nil {
		check.errors = append(check.errors, errors.Errorf("Failed to walk directory: %v", err))
	}

	return check
}
