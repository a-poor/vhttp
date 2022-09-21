package vhttp

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/hashicorp/go-multierror"
)

// Common headers, used for convenience
const (
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "Accept"
	HeaderHost          = "Host"
	HeaderAuthorization = "Authorization"
	HeaderConnection    = "Connection"
)

// Common Content-Type / MIME type values
const (
	MimePlain          = "text/plain"
	MimeHTML           = "text/html"
	MimeCSS            = "text/css"
	MimeTextJavascript = "text/javascript"

	MimeJSON = "application/json"
	MimeXML  = "application/xml"

	MimeImageAPNG   = "image/apng"
	MimeImageAVIF   = "image/avif"
	MimeImageGIF    = "image/gif"
	MimeImageJPEG   = "image/jpeg"
	MimeImagePNG    = "image/png"
	MimeImageSVGXML = "image/svg+xml"
	MimeImageWEBP   = "image/webp"
)

// Regular expressions for matching against (simplified)
// Authentication HTTP headers.
var (
	BasicAuthMatch  = regexp.MustCompile(`^Basic .+$`)
	BearerAuthMatch = regexp.MustCompile(`^Bearer .+$`)
)

// CanonicallHeaderKey converts the given string to a canonical form.
//
// This function is used in Header validation function, on the keys of the
// HeaderValidators parameter. By copying the function from the net/http
// package, this can be replaced with a custom implementation, if that
// functionality needs to be changed.
var CanonicalHeaderKey func(string) string = http.CanonicalHeaderKey

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

// ValidateRequest validates the request against the given validators.
//
//	err := vhttp.ValidateRequest(req,
//		vhttp.MethodIsGet(),                      // Method == GET
//		vhttp.BodyIsValidJSON(),                  // Body can be parsed as JSON
//		vhttp.HeaderContentTypeJSON(),            // Body has JSON content-type header
//		vhttp.HeaderAuthorizationMatchesBearer(), // "Authorization" header matches regex `^Bearer .+$`
//		vhttp.URLPathIs("/users/all"),            // URL Path == "/users/all"
//	)
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
	if merr != nil {
		return multierror.Flatten(merr)
	}
	return nil
}

// ValidateRequestFF validates the request against the given validators
// but fails fast.
//
// This is a version of the ValidateRequest function that stops at the
// first validation error returned (if any).
func ValidateRequestFF(req *http.Request, vs ...RequestValidator) error {
	// Check that the request is not nil
	if req == nil {
		return fmt.Errorf("request is nil")
	}

	// Iterate through the request validators.
	for _, v := range vs {
		if err := v.ValidateRequest(req); err != nil {
			return err
		}
	}

	// Success!
	return nil
}

// ValidateResponse validates the response against the given validators.
func ValidateResponse(res *http.Response, vs ...ResponseValidator) error {
	// Check that the response is not nil
	if res == nil {
		return fmt.Errorf("response is nil")
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
	if merr != nil {
		return multierror.Flatten(merr)
	}
	return nil
}

// ValidateResponseFF validates the response against the given validators
// but fails fast.
//
// This is a version of the ValidateResponse function that stops at the
// first validation error returned (if any).
func ValidateResponseFF(res *http.Response, vs ...ResponseValidator) error {
	// Check that the response is not nil
	if res == nil {
		return fmt.Errorf("response is nil")
	}

	// Iterate through the response validators.
	for _, v := range vs {
		if err := v.ValidateResponse(res); err != nil {
			return err
		}
	}

	// Success!
	return nil
}
