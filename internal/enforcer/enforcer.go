package enforcer

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/autonomy/conform/internal/policy"
	"github.com/autonomy/conform/internal/policy/commit"
	"github.com/autonomy/conform/internal/policy/conventionalcommit"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"

	yaml "gopkg.in/yaml.v2"
)

// Conform is a struct that conform.yaml gets decoded into.
type Conform struct {
	Policies []*PolicyDeclaration `yaml:"policies"`
}

// PolicyDeclaration allows a user to declare an arbitrary type along with a
// spec that will be decoded into the appropriate concrete type.
type PolicyDeclaration struct {
	Type string      `yaml:"type"`
	Spec interface{} `yaml:"spec"`
}

// policyMap defines the set of policies allowed within Conform.
var policyMap = map[string]policy.Policy{
	"commit":             &commit.Commit{},
	"conventionalCommit": &conventionalcommit.Conventional{},
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

	return c, nil
}

// Enforce enforces all policies defined in the conform.yaml file.
func (c *Conform) Enforce(setters ...policy.Option) error {
	opts := policy.NewDefaultOptions(setters...)

	for _, p := range c.Policies {
		fmt.Printf("Enforcing policy %q ... ", p.Type)
		err := c.enforce(p, opts)
		if err != nil {
			fmt.Printf("FAILED\n")
			return err
		}
		fmt.Printf("PASS\n")
	}

	return nil
}

func (c *Conform) enforce(declaration *PolicyDeclaration, opts *policy.Options) error {
	if _, ok := policyMap[declaration.Type]; !ok {
		return errors.Errorf("Policy %q is not defined", declaration.Type)
	}
	p := policyMap[declaration.Type]
	err := mapstructure.Decode(declaration.Spec, p)
	if err != nil {
		return err
	}

	report := p.Compliance(opts)

	if !report.Valid() {
		fmt.Printf("Violation of policy %q:\n", declaration.Type)
		for i, err := range report.Errors {
			fmt.Printf("\tViolation %d: %v\n", i+1, err)
		}
		os.Exit(1)
	}

	return nil
}
