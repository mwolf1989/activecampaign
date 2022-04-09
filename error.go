package activecampaign

type Error struct {
	Op  string
	Err error
}

func (e *Error) Unwrap() error { return e.Err }
func (e *Error) Error() string { return e.Op + ": " + e.Err.Error() }
