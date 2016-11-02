package errs

import "fmt"

type HandlerError struct {
	innerError  error
	handlerName string
}

func (h HandlerError) Error() string {
	return fmt.Sprintf("[%s]: %s", h.handlerName, h.innerError)
}

func NewHandlerError(iErr error, hName string) error {
	return HandlerError{
		innerError:  iErr,
		handlerName: hName,
	}
}
