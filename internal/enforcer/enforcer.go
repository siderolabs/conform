/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package enforcer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"

	"github.com/autonomy/conform/internal/policy"
	"github.com/autonomy/conform/internal/policy/commit"
	"github.com/autonomy/conform/internal/policy/license"
	"github.com/autonomy/conform/internal/summarizer"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"

	yaml "gopkg.in/yaml.v2"
)

// Conform is a struct that conform.yaml gets decoded into.
type Conform struct {
	Policies   []*PolicyDeclaration `yaml:"policies"`
	summarizer summarizer.Summarizer
}

// PolicyDeclaration allows a user to declare an arbitrary type along with a
// spec that will be decoded into the appropriate concrete type.
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
func New() (*Conform, error) {
	configBytes, err := ioutil.ReadFile(".conform.yaml")
	if err != nil {
		return nil, err
	}
	c := &Conform{}
	err = yaml.Unmarshal(configBytes, c)
	if err != nil {
		return nil, err
	}

	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if ok {
		s, err := summarizer.NewGitHubSummarizer(token)
		if err != nil {
			return nil, err
		}
		c.summarizer = s
	} else {
		c.summarizer = &summarizer.Noop{}
	}

	return c, nil
}

// Enforce enforces all policies defined in the conform.yaml file.
func (c *Conform) Enforce(setters ...policy.Option) {
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
				if err := c.summarizer.SetStatus("failure", p.Type, check.Name(), check.Message()); err != nil {
					log.Printf("WARNING: summary failed: %+v", err)
				}
				pass = false
			} else {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", p.Type, check.Name(), "PASS", "<none>")
				if err := c.summarizer.SetStatus("success", p.Type, check.Name(), check.Message()); err != nil {
					log.Printf("WARNING: summary failed: %+v", err)
				}
			}
		}
	}

	// nolint: errcheck
	w.Flush()

	if !pass {
		os.Exit(1)
	}
}

func (c *Conform) enforce(declaration *PolicyDeclaration, opts *policy.Options) (*policy.Report, error) {
	if _, ok := policyMap[declaration.Type]; !ok {
		return nil, errors.Errorf("Policy %q is not defined", declaration.Type)
	}

	p := policyMap[declaration.Type]

	err := mapstructure.Decode(declaration.Spec, p)
	if err != nil {
		return nil, errors.Errorf("Internal error: %v", err)
	}

	return p.Compliance(opts)
}
