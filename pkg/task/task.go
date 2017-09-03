package task

// Task defines a stage that can be used within a pipeline.
type Task struct {
	Template string `yaml:"template"`
	Rendered string
}
