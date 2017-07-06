package branch

import (
	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/pipeline"
	"github.com/autonomy/conform/conform/policy"
	"github.com/autonomy/conform/conform/stage"
)

// Branch implements the policy.Policy interface.
type Branch struct {
	Spec map[string]struct {
		Pipelines []*struct {
			Name string `yaml:"name"`
		} `yaml:"pipelines"`
	} `yaml:"spec"`
}

// Compliance implements the policy.Policy.Compliance function.
func (b *Branch) Compliance(metadata *metadata.Metadata, options ...policy.Option) policy.Report {
	args := &policy.Options{}
	for _, setter := range options {
		setter(args)
	}
	report := policy.Report{}
	branch := metadata.Git.Branch
	if _, ok := b.Spec[branch]; ok {
		for _, pipeline := range b.Spec[branch].Pipelines {
			_ = args.Pipelines[pipeline.Name].Build(metadata, args.Stages)
		}
	} else if spec, ok := b.Spec["default"]; ok {
		for _, pipeline := range spec.Pipelines {
			_ = args.Pipelines[pipeline.Name].Build(metadata, args.Stages)
		}
	}

	return report
}

// Pipelines implements the policy.Policy.Pipelines function.
func (b *Branch) Pipelines(pipelines map[string]*pipeline.Pipeline) policy.Option {
	return func(args *policy.Options) {
		args.Pipelines = pipelines
	}
}

// Stages implements the policy.Policy.Stages function.
func (b *Branch) Stages(stages map[string]*stage.Stage) policy.Option {
	return func(args *policy.Options) {
		args.Stages = stages
	}
}
