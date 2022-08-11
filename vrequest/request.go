package vrequest

import (
	"fmt"
	"net/http"

	"github.com/a-poor/vhttp"
	"github.com/hashicorp/go-multierror"
)

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
