package stage

// Stage defines a stage within a pipeline.
type Stage struct {
	Artifacts []string `yaml:"artifacts"`
	Tasks     []string `yaml:"tasks"`
}
