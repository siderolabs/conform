// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package license provides license policy.
package license

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/denormal/go-gitignore"
	"github.com/pkg/errors"

	"github.com/siderolabs/conform/internal/policy"
)

// Licenses implement the policy.Policy interface and enforces source code license headers.
type Licenses []License

// License represents a single license policy.
//
//nolint:govet
type License struct {
	Root string `mapstructure:"root"`
	// SkipPaths applies gitignore-style patterns to file paths to skip completely
	// parts of the tree which shouldn't be scanned (e.g. .git/)
	SkipPaths []string `mapstructure:"skipPaths"`
	// IncludeSuffixes is the regex used to find files that the license policy
	// should be applied to.
	IncludeSuffixes []string `mapstructure:"includeSuffixes"`
	// ExcludeSuffixes is the Suffixes used to find files that the license policy
	// should not be applied to.
	ExcludeSuffixes []string `mapstructure:"excludeSuffixes"`
	// AllowPrecedingComments, when enabled, allows blank lines and `//` and `#` line comments
	// before the license header. Useful for code generators that put build constraints or
	// "DO NOT EDIT" lines before the license.
	AllowPrecedingComments bool `mapstructure:"allowPrecedingComments"`
	// Header is the contents of the license header.
	Header string `mapstructure:"header"`
}

// Compliance implements the policy.Policy.Compliance function.
func (l *Licenses) Compliance(_ *policy.Options) (*policy.Report, error) {
	report := &policy.Report{}

	report.AddCheck(l.ValidateLicenseHeaders())

	return report, nil
}

// HeaderCheck enforces a license header on source code files.
type HeaderCheck struct {
	licenseErrors []error
}

// Name returns the name of the check.
func (l HeaderCheck) Name() string {
	return "File Header"
}

// Message returns to check message.
func (l HeaderCheck) Message() string {
	if len(l.licenseErrors) != 0 {
		return fmt.Sprintf("Found %d files without license header", len(l.licenseErrors))
	}

	return "All files have a valid license header"
}

// Errors returns any violations of the check.
func (l HeaderCheck) Errors() []error {
	return l.licenseErrors
}

// ValidateLicenseHeaders checks the header of a file and ensures it contains the provided value.
func (l Licenses) ValidateLicenseHeaders() policy.Check { //nolint:ireturn
	check := HeaderCheck{}

	for _, license := range l {
		if license.Root == "" {
			license.Root = "."
		}

		check.licenseErrors = append(check.licenseErrors, validateLicenseHeader(license)...)
	}

	return check
}

//nolint:gocognit
func validateLicenseHeader(license License) []error {
	var errs []error

	var buf bytes.Buffer

	for _, pattern := range license.SkipPaths {
		fmt.Fprintf(&buf, "%s\n", pattern)
	}

	patternmatcher := gitignore.New(&buf, license.Root, func(e gitignore.Error) bool {
		errs = append(errs, e.Underlying())

		return true
	})

	if license.Header == "" {
		errs = append(errs, errors.New("Header is not defined"))

		return errs
	}

	value := []byte(strings.TrimSpace(license.Header))

	err := filepath.Walk(license.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if patternmatcher.Relative(path, info.IsDir()) != nil {
			if info.IsDir() {
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
			for _, suffix := range license.ExcludeSuffixes {
				if strings.HasSuffix(info.Name(), suffix) {
					return nil
				}
			}

			// Check files matching the included suffixes.
			for _, suffix := range license.IncludeSuffixes {
				if strings.HasSuffix(info.Name(), suffix) {
					if license.AllowPrecedingComments {
						err = validateFileWithPrecedingComments(path, value)
					} else {
						err = validateFile(path, value)
					}

					if err != nil {
						errs = append(errs, err)
					}
				}
			}
		}

		return nil
	})
	if err != nil {
		errs = append(errs, errors.Errorf("Failed to walk directory: %v", err))
	}

	return errs
}

func validateFile(path string, value []byte) error {
	contents, err := os.ReadFile(path)
	if err != nil {
		return errors.Errorf("Failed to read %s: %s", path, err)
	}

	if bytes.HasPrefix(contents, value) {
		return nil
	}

	return errors.Errorf("File %s does not contain a license header", path)
}

func validateFileWithPrecedingComments(path string, value []byte) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Errorf("Failed to open %s: %s", path, err)
	}
	defer f.Close() //nolint:errcheck

	var contents []byte

	// read lines until the first non-comment line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		comment := line == ""
		comment = comment || strings.HasPrefix(line, "//")
		comment = comment || strings.HasPrefix(line, "#")

		if !comment {
			break
		}

		contents = append(contents, scanner.Bytes()...)
		contents = append(contents, '\n')
	}

	if err := scanner.Err(); err != nil {
		return errors.Errorf("Failed to check file %s: %s", path, err)
	}

	if bytes.Contains(contents, value) {
		return nil
	}

	return errors.Errorf("File %s does not contain a license header", path)
}
