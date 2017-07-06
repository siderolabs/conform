package commit

import (
	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/pipeline"
	"github.com/autonomy/conform/conform/policy"
	"github.com/autonomy/conform/conform/policy/commit/conventional"
	"github.com/autonomy/conform/conform/stage"
	"github.com/mitchellh/mapstructure"
)

type Commit struct {
	Type string      `yaml:"type"`
	Spec interface{} `yaml:"spec"`
}

func (c *Commit) Compliance(metadata *metadata.Metadata, options ...policy.Option) policy.Report {
	report := policy.Report{}
	switch c.Type {
	case "conventional":
		conventional := conventional.Conventional{}
		err := mapstructure.Decode(c.Spec, &conventional)
		report.Errors = append(report.Errors, err)
		report = conventional.Compliance(metadata)
	}

	return report
}

func (c *Commit) Pipelines(map[string]*pipeline.Pipeline) policy.Option {
	return func(args *policy.Options) {}
}

func (c *Commit) Stages(map[string]*stage.Stage) policy.Option {
	return func(args *policy.Options) {}
}
