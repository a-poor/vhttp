package vhttp

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-multierror"
)

// ValidateRequest validates the request against the given validators.
func ValidateRequest(req *http.Request, vs ...RequestValidator) error {
	// Check that the request is not nil
	if req == nil {
		return fmt.Errorf("request is nil")
	}

	// Iterate through the request validators.
	var merr *multierror.Error
	for _, v := range vs {
		err := v.ValidateRequest(req)
		if err != nil {
			merr = multierror.Append(merr, err)
		}
	}

	// Return the multi-error as an error (or nil if there are no errors).
	return merr.ErrorOrNil()
}

// ValidateResponse validates the response against the given validators.
func ValidateResponse(res *http.Response, vs ...ResponseValidator) error {
	// Check that the response is not nil
	if res == nil {
		return fmt.Errorf("request is nil")
	}

	// Iterate through the response validators.
	var merr *multierror.Error
	for _, v := range vs {
		err := v.ValidateResponse(res)
		if err != nil {
			merr = multierror.Append(merr, err)
		}
	}

	// Return the multi-error as an error (or nil if there are no errors).
	return merr.ErrorOrNil()
}

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

// ResponseValidator is a validator that validates an http.Response.
type ResponseValidator interface {
	ValidateResponse(res *http.Response) error
}

// URLValidator is a validator function that validates an http.Request's
// URL field.
type URLValidator func(*url.URL) error

func (v URLValidator) ValidateRequest(req *http.Request) error {
	return v(req.URL)
}

// MethodValidator is a validator that validates an http.Request's method.
type MethodValidator func(string) error

func (v MethodValidator) ValidateRequest(req *http.Request) error {
	return v(req.Method)
}

// BodyValidator is a validator that validates an http.Request's body.
//
// Note that this expects the body to be fully read as a byte slice.
// If more than one BodyValidator is being used, you should use a
// CachedBodyValidator instead – which will read the body once and
// pass the resulting byte slice to all of the BodyValidators.
type BodyValidator func(b []byte) error

func (v BodyValidator) ValidateRequest(req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return InternalErr(fmt.Errorf("failed to read request body: %s", err))
	}
	return v(b)
}

func (v BodyValidator) ValidateResponse(res *http.Response) error {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return InternalErr(fmt.Errorf("failed to read response body: %s", err))
	}
	return v(b)
}

// CachedBodyValidator is a RequestValidator/ResponseValidator that reads the
// Request or Response body once and passes the byte slice to each of it's
// BodyValidators (rather than calling their ValidateRequest or ValidateResponse
// methods each time).
//
// This helps avoid duplicated reads of the body and prevents issues with
// attemps to read the body after it has been closed.
type CachedBodyValidator struct {
	vs []BodyValidator
}

// CacheBody creates a new CachedBodyValidator
func CacheBody(vs ...BodyValidator) CachedBodyValidator {
	return CachedBodyValidator{vs}
}

func (v CachedBodyValidator) ValidateRequest(req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return InternalErr(fmt.Errorf("failed to read request body: %s", err))
	}

	var merr *multierror.Error
	for _, v := range v.vs {
		if err := v(b); err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	return merr.ErrorOrNil()
}

func (v CachedBodyValidator) ValidateResponse(res *http.Response) error {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return InternalErr(fmt.Errorf("failed to read response body: %s", err))
	}

	var merr *multierror.Error
	for _, v := range v.vs {
		if err := v(b); err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	return merr.ErrorOrNil()
}

// HeaderValidator is a validator that validates an http.Request or http.Response
// object's headers.
type HeaderValidator func(http.Header) error

func (v HeaderValidator) ValidateRequest(req *http.Request) error {
	return v(req.Header)
}

func (v HeaderValidator) ValidateResponse(res *http.Response) error {
	return v(res.Header)
}

// TLSValidator is a validator that validates an http.Request or http.Response
// object's TLS connection.
type TLSValidator func(*tls.ConnectionState) error

func (v TLSValidator) ValidateRequest(req *http.Request) error {
	return v(req.TLS)
}

func (v TLSValidator) ValidateResponse(res *http.Response) error {
	return v(res.TLS)
}

// ProtoValidator is a validator that validates an http.Request or http.Response's
// Proto field.
type ProtoValidator func(string) error

func (v ProtoValidator) ValidateRequest(req *http.Request) error {
	return v(req.Proto)
}

func (v ProtoValidator) ValidateResponse(res *http.Response) error {
	return v(res.Proto)
}
