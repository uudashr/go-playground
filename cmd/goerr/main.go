package main

import (
	"errors"
	"fmt"
)

type ErrDuplicateRequest struct {
	cause error
}

func (e *ErrDuplicateRequest) Unwrap() error {
	return e.cause
}

func (e *ErrDuplicateRequest) Error() string {
	if e.cause == nil {
		return "duplicate request"
	}

	return fmt.Sprintf("duplicate request: %v", e.cause)
}

var (
	ErrOperationFailed = errors.New("duplicate request")
	ErrValidation      = errors.New("validation error")
)

func main() {
	err := fmt.Errorf("have error %w %w", ErrOperationFailed, ErrValidation)
	fmt.Println(err)

	fmt.Println("A duplicate request:", errors.Is(err, ErrOperationFailed))
	fmt.Println("An operation failure:", errors.Is(err, ErrValidation))

	fmt.Println(unwrapErrors(err))
	// It difficult to recognize the cause
}

func unwrapErrors(err error) []error {
	errorsWrapper, ok := err.(UnwrapErrors)
	if !ok {
		return []error{err}
	}

	return errorsWrapper.Unwrap()
}

type UnwrapErrors interface {
	Unwrap() []error
}
