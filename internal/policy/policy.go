package policy

// Report summarizes the compliance of a policy.
type Report struct {
	Errors []error
}

// Policy is an interface that policies must implement.
type Policy interface {
	Compliance(*Options) Report
}

// Valid checks if a report is valid.
func (r Report) Valid() bool {
	return len(r.Errors) == 0
}
