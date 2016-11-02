package errs

import "errors"

var (
	// ErrHandlerNameDuplicate occurs when a handler is attempted to
	// be registered but another handler with the same name already
	// exists
	ErrHandlerNameDuplicate = errors.New("Duplicate Handler Name!")
)
