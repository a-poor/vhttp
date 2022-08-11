package vhttp

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
