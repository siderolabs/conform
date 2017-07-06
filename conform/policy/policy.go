package policy

// Report summarizes the compliance of a policy.
type Report struct {
	Valid  bool
	Errors []error
}

// Policy is an interface used for enforcing policies.
type Policy interface {
	Compliance(interface{}) (report *Report, err error)
}
