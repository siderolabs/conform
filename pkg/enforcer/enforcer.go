package enforcer

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/autonomy/conform/pkg/metadata"
	"github.com/autonomy/conform/pkg/pipeline"
	"github.com/autonomy/conform/pkg/policy"
	"github.com/autonomy/conform/pkg/policy/conventionalcommit"
	"github.com/autonomy/conform/pkg/script"
	"github.com/autonomy/conform/pkg/stage"
	"github.com/autonomy/conform/pkg/task"
	"github.com/mitchellh/mapstructure"

	yaml "gopkg.in/yaml.v2"
)

// Conform is a struct that conform.yaml gets decoded into.
type Conform struct {
	Metadata *metadata.Metadata      `yaml:"metadata"`
	Policies []*PolicyDeclaration    `yaml:"policies"`
	Pipeline *pipeline.Pipeline      `yaml:"pipeline"`
	Stages   map[string]*stage.Stage `yaml:"stages"`
	Tasks    map[string]*task.Task   `yaml:"tasks"`
	Script   *script.Script          `yaml:"script"`
}

// PolicyDeclaration allows a user to declare an arbitrary type along with a
// spec that will be decoded into the appropriate concrete type.
type PolicyDeclaration struct {
	Type string      `yaml:"type"`
	Spec interface{} `yaml:"spec"`
}

// policyMap defines the set of policies allowed within Conform.
var policyMap = map[string]policy.Policy{
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
func (c *Conform) Enforce() error {
	for _, p := range c.Policies {
		fmt.Printf("Enforcing policy %q: ", p.Type)
		err := c.enforce(p)
		if err != nil {
			fmt.Printf("failed\n")
			return err
		}
		fmt.Printf("passed\n")
	}

	return nil
}

func (c *Conform) enforce(p *PolicyDeclaration) error {
	if _, ok := policyMap[p.Type]; !ok {
		return fmt.Errorf("Policy %q is not defined", p.Type)
	}
	policy := policyMap[p.Type]
	err := mapstructure.Decode(p.Spec, policy)
	if err != nil {
		return err
	}

	report := policy.Compliance(
		c.Metadata,
		policy.Pipeline(c.Pipeline),
		policy.Tasks(c.Tasks),
	)

	if !report.Valid() {
		fmt.Printf("Violation of policy %q:\n", p.Type)
		for i, err := range report.Errors {
			fmt.Printf("\tViolation %d: %v\n", i, err)
		}
		os.Exit(1)
	}

	return nil
}
