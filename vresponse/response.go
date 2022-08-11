package vresponse

import (
	"net/http"

	"github.com/a-poor/vhttp"
)

// ResponseValidator is a function that validates an HTTP response object.
type ResponseValidator vhttp.Validator[*http.Response]
