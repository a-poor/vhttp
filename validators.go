package vhttp

import (
	"fmt"
	"net/http"
)

// InternalError is used to signal that an error returned from a validator
// is not a validation error but an internal error.
//
// An InternalError is meant to wrap a real error returned in the process
// of validating a request or response.
type InternalError struct {
	err error
}

// InternalErr creates a new InternalError
func InternalErr(err error) error {
	return InternalError{err}
}

func (e InternalError) Error() string {
	return e.err.Error()
}

func (e InternalError) Unwrap() error {
	return e.err
}

// Wrap returns a new InternalError where the underlying error is the
// result of wrapping e's underlying error (using fmt.Errorf) with
// the message msg prepended to the error message chain.
//
// Equivalent to:
//
//	err := vhttp.InternalErr(fmt.Errorf("%s: %w", msg, ierr.Unwrap()))
func (e InternalError) Wrap(msg string) error {
	return InternalErr(fmt.Errorf("%s: %w", msg, e.err))
}

// RequestValidator is a validator that validates an http.Request.
type RequestValidator interface {
	ValidateRequest(req *http.Request) error
}

// RequestFunc is a function that validates an http.Request and
// can act as a RequestValidator.
type RequestFunc func(req *http.Request) error

func (v RequestFunc) ValidateRequest(req *http.Request) error {
	return v(req)
}

// ResponseValidator is a validator that validates an http.Response.
type ResponseValidator interface {
	ValidateResponse(res *http.Response) error
}

// ResponseFunc is a function that validates an http.Response and
// can act as a ResponseValidator.
type ResponseFunc func(res *http.Response) error

func (v ResponseFunc) ValidateResponse(res *http.Response) error {
	return v(res)
}
