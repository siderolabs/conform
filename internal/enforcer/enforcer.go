// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package enforcer defines policy enforcement.
package enforcer

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v3"

	"github.com/siderolabs/conform/internal/policy"
	"github.com/siderolabs/conform/internal/policy/commit"
	"github.com/siderolabs/conform/internal/policy/license"
	"github.com/siderolabs/conform/internal/reporter"
)

// Conform is a struct that conform.yaml gets decoded into.
//
//nolint:govet
type Conform struct {
	Policies []*PolicyDeclaration `yaml:"policies"`
	reporter reporter.Reporter
}

// PolicyDeclaration allows a user to declare an arbitrary type along with a
// spec that will be decoded into the appropriate concrete type.
//
//nolint:govet
type PolicyDeclaration struct {
	Type string      `yaml:"type"`
	Spec interface{} `yaml:"spec"`
}

// policyMap defines the set of policies allowed within Conform.
var policyMap = map[string]policy.Policy{
	"commit":  &commit.Commit{},
	"license": &license.License{},
	// "version":    &version.Version{},
}

// New loads the conform.yaml file and unmarshals it into a Conform struct.
func New(r string) (*Conform, error) {
	c := &Conform{}

	switch r {
	case "github":
		s, err := reporter.NewGitHubReporter()
		if err != nil {
			return nil, err
		}

		c.reporter = s
	default:
		c.reporter = &reporter.Noop{}
	}

	configBytes, err := os.ReadFile(".conform.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configBytes, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Enforce enforces all policies defined in the conform.yaml file.
func (c *Conform) Enforce(setters ...policy.Option) error {
	opts := policy.NewDefaultOptions(setters...)

	const padding = 8
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, "POLICY\tCHECK\tSTATUS\tMESSAGE\t")

	pass := true

	for _, p := range c.Policies {
		report, err := c.enforce(p, opts)
		if err != nil {
			log.Fatal(err)
		}

		for _, check := range report.Checks() {
			if len(check.Errors()) != 0 {
				for _, err := range check.Errors() {
					fmt.Fprintf(w, "%s\t%s\t%s\t%v\t\n", p.Type, check.Name(), "FAILED", err)
				}

				if err := c.reporter.SetStatus("failure", p.Type, check.Name(), check.Message()); err != nil {
					log.Printf("WARNING: report failed: %+v", err)
				}

				pass = false
			} else {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", p.Type, check.Name(), "PASS", check.Message())

				if err := c.reporter.SetStatus("success", p.Type, check.Name(), check.Message()); err != nil {
					log.Printf("WARNING: report failed: %+v", err)
				}
			}
		}
	}

	w.Flush() //nolint:errcheck

	if !pass {
		return errors.New("1 or more policy failed")
	}

	return nil
}

func (c *Conform) enforce(declaration *PolicyDeclaration, opts *policy.Options) (*policy.Report, error) {
	if _, ok := policyMap[declaration.Type]; !ok {
		return nil, errors.Errorf("Policy %q is not defined", declaration.Type)
	}

	p := policyMap[declaration.Type]

	// backwards compatibility, convert `gpg: bool` into `gpg: required: bool`
	if declaration.Type == "commit" {
		if spec, ok := declaration.Spec.(map[interface{}]interface{}); ok {
			if gpg, ok := spec["gpg"]; ok {
				if val, ok := gpg.(bool); ok {
					spec["gpg"] = map[string]interface{}{
						"required": val,
					}
				}
			}
		}
	}

	err := mapstructure.Decode(declaration.Spec, p)
	if err != nil {
		return nil, errors.Errorf("Internal error: %v", err)
	}

	return p.Compliance(opts)
}
