package vhttp

import (
	"net/http"
	"regexp"
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

// Validator is a function that validates a value of type T.
type Validator[T any] func(T) error

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
