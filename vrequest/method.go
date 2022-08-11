package vrequest

import (
	"fmt"
	"net/http"
)

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
