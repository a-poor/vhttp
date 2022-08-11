package vhttp

import (
	"fmt"
	"net/http"
)

// MethodValidator is a validator that validates an http.Request's method.
type MethodValidator func(string) error

func (v MethodValidator) ValidateRequest(req *http.Request) error {
	return v(req.Method)
}

// MethodIs creates a request validator that checks that the request method
// is equal to the given method m.
func MethodIs(s string) MethodValidator {
	return func(m string) error {
		if m != s {
			return fmt.Errorf("expected method %q, found %q", s, m)
		}
		return nil
	}
}

// MethodIs creates a request validator that checks that the request method
// is NOT equal to the given method m.
func MethodIsNot(s string) MethodValidator {
	return func(m string) error {
		if m == s {
			return fmt.Errorf("expected method %q, found %q", s, m)
		}
		return nil
	}
}

// MethodIsGet creates a request validator that checks that the request
// is a GET request.
func MethodIsGet() MethodValidator {
	return MethodIs(http.MethodGet)
}

// MethodIsPost creates a request validator that checks that the request
// is a POST request.
func MethodIsPost() MethodValidator {
	return MethodIs(http.MethodPost)
}

// MethodIsPut creates a request validator that checks that the request
// is a PUT request.
func MethodIsPut() MethodValidator {
	return MethodIs(http.MethodPut)
}

// MethodIsDelete creates a request validator that checks that the request
// is a DELETE request.
func MethodIsDelete() MethodValidator {
	return MethodIs(http.MethodDelete)
}

// MethodIsOptions creates a request validator that checks that the request
// is a OPTIONS request.
func MethodIsOptions() MethodValidator {
	return MethodIs(http.MethodOptions)
}

// MethodIsPatch creates a request validator that checks that the request
// is a PATCH request.
func MethodIsPatch() MethodValidator {
	return MethodIs(http.MethodPatch)
}
