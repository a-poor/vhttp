package vrequest

/*

Request Rules:
- Status Code
  - Is x
  - IsNot x
  - IsIn [x]
  - IsNotIn [x]
  - InRange x-y
  - NotInRange x-y
  - IsSuccess
  - IsError
  - IsRedirect
  - IsClientError
  - IsServerError
- Method
  - Is x
  - IsNot x
  - IsIn [x]
  - IsNotIn [x]
- URL
  - Is x
  - IsNot x
- Headers
- Body (as bytes)
- TCP
- Form Data

Response Rules:
-

*/

import (
	"fmt"
	"net/http"

	"github.com/a-poor/vhttp"
	"github.com/hashicorp/go-multierror"
)

// CanonicallHeaderKey converts the given string to a canonical form.
//
// This function is used in the Header function, on the keys of the
// HeaderValidators parameter. By copying the function from the net/http
// package, this can be replaced with a custom implementation, if that
// functionality needs to be changed.
var CanonicalHeaderKey func(string) string = http.CanonicalHeaderKey

// RequestValidator is a function that validates an HTTP request object.
type RequestValidator vhttp.Validator[*http.Request]

// ValidateRequest is a function that validates an HTTP request object
// with a series of validator functions.
func ValidateRequest(req *http.Request, validators ...RequestValidator) error {
	// Check that the request is not nil
	if req == nil {
		return multierror.Append(fmt.Errorf("request is nil")).ErrorOrNil()
	}

	// Iterate through the request validators.
	var merr *multierror.Error
	for _, v := range validators {
		err := v(req)
		if err != nil {
			merr = multierror.Append(merr, err)
		}
	}

	// Return the multi-error as an error (or nil if there are no errors).
	return merr.ErrorOrNil()
}

// MethodIs creates a request validator that checks that the request method
// is equal to the given method m.
func MethodIs(m string) RequestValidator {
	return func(req *http.Request) error {
		if req.Method != m {
			return fmt.Errorf("expected method %q, found %q", m, req.Method)
		}
		return nil
	}
}

// MethodIs creates a request validator that checks that the request method
// is NOT equal to the given method m.
func MethodIsNot(m string) RequestValidator {
	return func(req *http.Request) error {
		if req.Method != m {
			return fmt.Errorf("expected method to not be %q, found %q", m, req.Method)
		}
		return nil
	}
}

// MethodIsGet creates a request validator that checks that the request
// is a GET request.
func MethodIsGet() RequestValidator {
	return MethodIs(http.MethodGet)
}

// MethodIsPost creates a request validator that checks that the request
// is a POST request.
func MethodIsPost() RequestValidator {
	return MethodIs(http.MethodPost)
}

// MethodIsPut creates a request validator that checks that the request
// is a PUT request.
func MethodIsPut() RequestValidator {
	return MethodIs(http.MethodPut)
}

// MethodIsDelete creates a request validator that checks that the request
// is a DELETE request.
func MethodIsDelete() RequestValidator {
	return MethodIs(http.MethodDelete)
}

// MethodIsOptions creates a request validator that checks that the request
// is a OPTIONS request.
func MethodIsOptions() RequestValidator {
	return MethodIs(http.MethodOptions)
}

// MethodIsPatch creates a request validator that checks that the request
// is a PATCH request.
func MethodIsPatch() RequestValidator {
	return MethodIs(http.MethodPatch)
}

func HasHeader(h string) RequestValidator {
	return func(r *http.Request) error {
		// Convert the header key to canonical form.
		h := CanonicalHeaderKey(h)

		// Check if the header is present.
		if _, ok := r.Header[h]; !ok {
			return fmt.Errorf("header %q not found", h)
		}

		// Found!
		return nil
	}
}

func HasHeaderContentType(ct string) RequestValidator {
	return HasHeader("Content-Type")
}

func RequestHasHeaderAccept() RequestValidator {
	return HasHeader("Accept")
}

func HasHeaderAuthorization() RequestValidator {
	return HasHeader("Authorization")
}

func HeaderIs(h, v string) RequestValidator {
	return func(req *http.Request) error {
		// Convert the header key to canonical form.
		h := CanonicalHeaderKey(h)

		// Get the header values
		vs, ok := req.Header[h]
		if !ok {
			return fmt.Errorf("header %q not found", h)
		}

		// Check if the header is present.
		for _, s := range vs {
			if s == v {
				return nil // Found!
			}
		}

		// Not found.
		return fmt.Errorf("expected header %q to have value %q", h, v)
	}
}

func HeaderContentTypeIs(ct string) RequestValidator {
	return HeaderIs("Content-Type", ct)
}

func HeaderContentTypeJSON() RequestValidator {
	return HeaderIs("Content-Type", "application/json")
}

func HeaderContentTypeXML() RequestValidator {
	return HeaderIs("Content-Type", "application/xml")
}
