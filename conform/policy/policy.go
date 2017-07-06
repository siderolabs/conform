package policy

import (
	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/pipeline"
	"github.com/autonomy/conform/conform/stage"
)

type Option func(*Options)

type Options struct {
	Pipelines map[string]*pipeline.Pipeline
	Stages    map[string]*stage.Stage
}

// Report summarizes the compliance of a policy.
type Report struct {
	Valid  bool
	Errors []error
}

// Policy is an interface that policies must implement.
type Policy interface {
	Compliance(*metadata.Metadata, ...Option) Report
	Pipelines(map[string]*pipeline.Pipeline) Option
	Stages(map[string]*stage.Stage) Option
}
