package policy

import (
	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/pipeline"
	"github.com/autonomy/conform/conform/task"
)

// Option is a functional option used to pass in arguments to a Policy.
type Option func(*Options)

// Options defines the set of options available to a Policy.
type Options struct {
	Pipeline *pipeline.Pipeline
	Tasks    map[string]*task.Task
}

// Report summarizes the compliance of a policy.
type Report struct {
	Errors []error
}

// Policy is an interface that policies must implement.
type Policy interface {
	Compliance(*metadata.Metadata, ...Option) Report
	Pipeline(*pipeline.Pipeline) Option
	Tasks(map[string]*task.Task) Option
}

// Valid checks if a report is valid.
func (r Report) Valid() bool {
	return len(r.Errors) == 0
}
