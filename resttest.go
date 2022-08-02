package resttest

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

// toAnySlice is a helper function to convert a generic
// slice to a slice of any type.
func toAnySlice[T any](vs []T) []any {
	as := make([]any, len(vs))
	for i, v := range vs {
		as[i] = any(v)
	}
	return as
}
