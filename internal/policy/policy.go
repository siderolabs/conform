/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package policy

// Report summarizes the compliance of a policy.
type Report struct {
	checks []Check
}

// Check defines a policy check.
type Check interface {
	Name() string
	Message() string
	Errors() []error
}

// Policy is an interface that policies must implement.
type Policy interface {
	Compliance(*Options) (*Report, error)
}

// Valid checks if a report is valid.
func (r *Report) Valid() bool {
	for _, check := range r.checks {
		if len(check.Errors()) != 0 {
			return false
		}
	}
	return true
}

// Checks returns the checks executed by a policy.
func (r *Report) Checks() []Check {
	return r.checks
}

// AddCheck adds a check to the policy report.
func (r *Report) AddCheck(c Check) {
	r.checks = append(r.checks, c)
}
