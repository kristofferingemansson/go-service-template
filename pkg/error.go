package pkg

// Error ..
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrNotImplemented ..
	ErrNotImplemented = Error("ErrNotImplemented")
)
