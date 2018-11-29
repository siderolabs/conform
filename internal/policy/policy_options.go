package policy

// Option is a functional option used to pass in arguments to a Policy.
type Option func(*Options)

// Options defines the set of options available to a Policy.
type Options struct {
	CommitMsgFile *string
}

// WithCommitMsgFile sets the path to the commit message file.
func WithCommitMsgFile(o *string) Option {
	return func(args *Options) {
		args.CommitMsgFile = o
	}
}

// NewDefaultOptions initializes a Options struct with default values.
func NewDefaultOptions(setters ...Option) *Options {
	opts := &Options{
		CommitMsgFile: nil,
	}

	for _, setter := range setters {
		setter(opts)
	}

	return opts
}
