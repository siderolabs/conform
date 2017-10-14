package stage

import "github.com/autonomy/conform/pkg/service"

// Stage defines a stage within a pipeline.
type Stage struct {
	Artifacts []*Artifact        `yaml:"artifacts"`
	Services  []*service.Service `yaml:"services"`
	Tasks     []string           `yaml:"tasks"`
}

// Artifact is a struct that represents an artifact.
type Artifact struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}
