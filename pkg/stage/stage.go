package stage

// Stage defines a stage within a pipeline.
type Stage struct {
	Artifacts []*Artifact `yaml:"artifacts"`
	Tasks     []string    `yaml:"tasks"`
}

// Artifact is a struct that represents an artifact.
type Artifact struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}
