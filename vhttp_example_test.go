package vhttp_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/a-poor/vhttp"
)

func asReadCloser(b []byte) io.ReadCloser {
	return io.NopCloser(bytes.NewReader(b))
}

func ExampleValidateRequest() {
	// Create a sample request...
	u, _ := url.Parse("https://example.com/api/v1/users")
	req := &http.Request{
		Method: http.MethodPost,
		Header: http.Header{
			"Content-Type":  []string{"application/json"},
			"Authorization": []string{"Basic abcde12345"},
		},
		Body: asReadCloser([]byte(`{{{{`)),
		URL:  u,
	}

	// Validate the request...
	err := vhttp.ValidateRequest(req,
		// Is a "GET" request
		vhttp.MethodIsGet(),

		// Calling json.Valid() on the body returns true
		vhttp.BodyIsValidJSON(),

		// Has the header "Content-Type" and it's equal to "application/json"
		vhttp.HeaderContentTypeJSON(),

		// The header "Authorization" matches the regular expression ^Bearer .+$
		vhttp.HeaderAuthorizationMatchesBearer(),

		// Has the URL path "/api/v2/posts"
		vhttp.URLPathIs("/api/v2/posts"),
	)

	// Print the output...
	fmt.Println(err)
	// Output:
	// 4 errors occurred:
	// 	* expected method "GET", found "POST"
	// 	* body is not valid JSON
	// 	* expected header "Authorization" to match "^Bearer .+$"
	// 	* expected URL path "/api/v2/posts", found "/api/v1/users"
	//
}
