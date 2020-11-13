package errors

import (
	"fmt"
	"runtime/debug"
)

type InternalError struct {
	err   error
	stack []byte
}

func (i *InternalError) Error() string {
	return fmt.Sprintf("%s\n ^s", i.err.Error(), string(i.stack))
}

func NewInternalError(err error) *InternalError {
	return &InternalError{
		err:   err,
		stack: debug.Stack(),
	}
}
