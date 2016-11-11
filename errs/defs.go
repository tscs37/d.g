package errs

import "errors"

var (
	// ErrHandlerNameDuplicate occurs when a handler is attempted to
	// be registered but another handler with the same name already
	// exists
	ErrHandlerNameDuplicate = errors.New("Duplicate Handler Name!")
	// ErrHandlerNotSupported indicates that it was attempted to add
	// a handler to a mux that doesn't support the specific handler type.
	ErrHandlerNotSupported = errors.New("Handler not supported")
	// ErrNotImplemnted indicates a particular function has not been implemented
	// yet.
	ErrNotImplemnted = errors.New("Function not implemented yet")
)
