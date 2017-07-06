package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config represents the YAML.
type Config struct {
	Debug     bool
	Default   *string           `yaml:"default"`
	Metadata  *Metadata         `yaml:"metadata"`
	Policies  *Policies         `yaml:"policies"`
	Scripts   map[string]string `yaml:"scripts"`
	Templates map[string]string `yaml:"templates"`
	Rules     map[string]*Rule  `yaml:"rules"`
}

// Metadata contains metadata.
type Metadata struct {
	Repository *string `yaml:"repository"`
	Registry   *string `yaml:"registry"`
}

// Policies contains policies that are enforced.
type Policies struct {
	Git *Git `yaml:"git"`
}

// Git contains git specific policies.
type Git struct {
	Types  []string `yaml:"types"`
	Scopes []string `yaml:"scopes"`
}

// Rule contains rules.
type Rule struct {
	Templates []string `yaml:"templates"`
	Artifacts []string `yaml:"artifacts"`
	Before    []string `yaml:"before"`
	After     []string `yaml:"after"`
}

// NewConfig instantiates and returns a config.
func NewConfig() (*Config, error) {
	rBytes, err := ioutil.ReadFile("conform.yaml")
	if err != nil {
		return nil, fmt.Errorf("Unable to load conform.yaml: %v", err)
	}
	c := Config{}
	err = yaml.Unmarshal(rBytes, &c)
	if err != nil {
		return nil, fmt.Errorf("Unable to load conform.yaml: %v", err)
	}

	return &c, nil
}
