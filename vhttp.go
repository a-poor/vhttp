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
	if merr != nil {
		return multierror.Flatten(merr)
	}
	return nil
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
	if merr != nil {
		return multierror.Flatten(merr)
	}
	return nil
}
