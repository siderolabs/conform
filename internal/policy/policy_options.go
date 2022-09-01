// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package policy

// Option is a functional option used to pass in arguments to a Policy.
type Option func(*Options)

// Options defines the set of options available to a Policy.
type Options struct {
	CommitMsgFile *string
	CommitRef     string
	RevisionRange string
}

// WithCommitMsgFile sets the path to the commit message file.
func WithCommitMsgFile(o *string) Option {
	return func(args *Options) {
		args.CommitMsgFile = o
	}
}

// WithCommitRef sets the ref to compare git policies against.
func WithCommitRef(o string) Option {
	return func(args *Options) {
		args.CommitRef = o
	}
}

// WithRevisionRange sets the revision range to compare git policies against.
func WithRevisionRange(o string) Option {
	return func(args *Options) {
		args.RevisionRange = o
	}
}

// NewDefaultOptions initializes a Options struct with default values.
func NewDefaultOptions(setters ...Option) *Options {
	opts := &Options{
		CommitMsgFile: nil,
		CommitRef:     "refs/heads/master",
	}

	for _, setter := range setters {
		setter(opts)
	}

	return opts
}
