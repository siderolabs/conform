package conform

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/pipeline"
	"github.com/autonomy/conform/conform/policy"
	"github.com/autonomy/conform/conform/policy/branch"
	"github.com/autonomy/conform/conform/policy/commit"
	"github.com/autonomy/conform/conform/stage"
	"github.com/mitchellh/mapstructure"

	yaml "gopkg.in/yaml.v2"
)

// Conform is a struct that conform.yaml gets decoded into.
type Conform struct {
	Metadata  *metadata.Metadata            `yaml:"metadata"`
	Policies  map[string]interface{}        `yaml:"policies"`
	Pipelines map[string]*pipeline.Pipeline `yaml:"pipelines"`
	Stages    map[string]*stage.Stage       `yaml:"stages"`
}

// policies defines the set of policies allowed within Conform.
var policies = map[string]policy.Policy{
	"commit": &commit.Commit{},
	"branch": &branch.Branch{},
	// "tag":    &tag.Tag{},
}

var policyOrder = []string{"commit", "branch"}

// New loads the conform.yaml file and unmarshals it into a Conform struct.
func New() (*Conform, error) {
	configBytes, err := ioutil.ReadFile("conform.yaml")
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
	for _, p := range policyOrder {
		err := c.enforce(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Conform) enforce(n string) error {
	// TODO: Check if the policy is a valid type
	// if policy, ok := p[n]; ok {
	//
	// }
	err := mapstructure.Decode(c.Policies[n], policies[n])
	if err != nil {
		return err
	}
	valid, errs := c.validate(policies[n])
	if valid {
		return nil
	}
	if len(errs) != 0 {
		fmt.Printf("Policy %q does not conform:\n", n)
		for _, err := range errs {
			fmt.Printf("\t%v\n", err)
		}
		os.Exit(1)
	}

	return nil
}

func (c *Conform) validate(p policy.Policy) (bool, []error) {
	report := p.Compliance(c.Metadata, p.Pipelines(c.Pipelines), p.Stages(c.Stages))
	if report.Valid {
		return true, nil
	}

	return false, report.Errors
}
