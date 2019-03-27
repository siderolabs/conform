/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package license

import (
	"bytes"
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
func (l *License) Compliance(options *policy.Options) (report policy.Report) {
	var err error

	report = policy.Report{}
	if l.Header == "" {
		report.Errors = append(report.Errors, errors.New("Header is not defined"))
		return report
	}
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
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
						report.Errors = append(report.Errors, errors.Errorf("Failed to open %s", path))
						return nil
					}
					ValidateLicenseHeader(&report, info.Name(), contents, []byte(l.Header))
					return nil
				}
			}
		}
		return nil
	})
	if err != nil {
		report.Errors = append(report.Errors, errors.Errorf("Failed to walk directory: %v", err))
	}

	return report
}

// ValidateLicenseHeader checks the header of a file and ensures it contains the
// provided value.
func ValidateLicenseHeader(report *policy.Report, name string, contents, value []byte) {
	if bytes.HasPrefix(contents, value) {
		return
	}
	report.Errors = append(report.Errors, errors.Errorf("File %s does not contain a license header", name))
}
