package vhttp_test

import (
	"net/http"
	"testing"

	"github.com/a-poor/vhttp"
)

func TestHasHeader(t *testing.T) {
	cases := []struct {
		name    string      // Case name
		key     string      // Header to search for
		headers http.Header // Request's headers
		isErr   bool        // Should an error be returned
	}{
		{
			name: "success-and-normalize-key",
			key:  "content-type",
			headers: http.Header{
				vhttp.HeaderContentType: []string{"application/json"},
			},
			isErr: false,
		},
		{
			name:    "missing-error",
			key:     "content-type",
			headers: http.Header{},
			isErr:   true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Create the validator function
			vf := vhttp.HasHeader(c.key)

			// Validate and check for an error
			err := vf(c.headers)

			// Unexpected error?
			if err != nil && !c.isErr {
				t.Errorf("unexpected error returned for error %q: %s", c.key, err)
				return
			}

			// Unexpected success?
			if err == nil && c.isErr {
				t.Errorf("expected an error to be returned searching for %q", c.key)
				return
			}

			// Expected error?
			if err != nil && c.isErr {
				t.Log("successfully returned an error!")
				return
			}

			// Expected success?
			if err == nil && !c.isErr {
				t.Log("successfully returned no error!")
				return
			}
		})
	}
}

func TestHasHeaderContentType(t *testing.T) {
	cases := []struct {
		name    string      // Case name
		headers http.Header // Request's headers
		isErr   bool        // Should an error be returned
	}{
		{
			name: "success-and-normalize-key",
			headers: http.Header{
				vhttp.HeaderContentType: []string{"application/json"},
			},
			isErr: false,
		},
		{
			name:    "missing-error",
			headers: http.Header{},
			isErr:   true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Create the validator function
			vf := vhttp.HasHeaderContentType()

			// Validate and check for an error
			err := vf(c.headers)

			// Unexpected error?
			if err != nil && !c.isErr {
				t.Errorf("unexpected error returned for error %q: %s", "Content-Type", err)
				return
			}

			// Unexpected success?
			if err == nil && c.isErr {
				t.Errorf("expected an error to be returned searching for %q", "Content-Type")
				return
			}

			// Expected error?
			if err != nil && c.isErr {
				t.Log("successfully returned an error!")
				return
			}

			// Expected success?
			if err == nil && !c.isErr {
				t.Log("successfully returned no error!")
				return
			}
		})
	}
}

func TestHasHeaderAccept(t *testing.T) {
	cases := []struct {
		name    string      // Case name
		headers http.Header // Request's headers
		isErr   bool        // Should an error be returned
	}{
		{
			name: "success-and-normalize-key",
			headers: http.Header{
				vhttp.HeaderAccept: []string{"application/json"},
			},
			isErr: false,
		},
		{
			name:    "missing-error",
			headers: http.Header{},
			isErr:   true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Create the validator function
			vf := vhttp.HasHeaderAccept()

			// Validate and check for an error
			err := vf(c.headers)

			// Unexpected error?
			if err != nil && !c.isErr {
				t.Errorf("unexpected error returned for error %q: %s", "Accept", err)
				return
			}

			// Unexpected success?
			if err == nil && c.isErr {
				t.Errorf("expected an error to be returned searching for %q", "Accept")
				return
			}

			// Expected error?
			if err != nil && c.isErr {
				t.Log("successfully returned an error!")
				return
			}

			// Expected success?
			if err == nil && !c.isErr {
				t.Log("successfully returned no error!")
				return
			}
		})
	}
}

func TestHasHeaderAuthorization(t *testing.T) {
	cases := []struct {
		name    string      // Case name
		headers http.Header // Request's headers
		isErr   bool        // Should an error be returned
	}{
		{
			name: "success-and-normalize-key",
			headers: http.Header{
				vhttp.HeaderAuthorization: []string{"Bearer abc123"},
			},
			isErr: false,
		},
		{
			name:    "missing-error",
			headers: http.Header{},
			isErr:   true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Create the validator function
			vf := vhttp.HasHeaderAuthorization()

			// Validate and check for an error
			err := vf(c.headers)

			// Unexpected error?
			if err != nil && !c.isErr {
				t.Errorf("unexpected error returned for error %q: %s", "Authorization", err)
				return
			}

			// Unexpected success?
			if err == nil && c.isErr {
				t.Errorf("expected an error to be returned searching for %q", "Authorization")
				return
			}

			// Expected error?
			if err != nil && c.isErr {
				t.Log("successfully returned an error!")
				return
			}

			// Expected success?
			if err == nil && !c.isErr {
				t.Log("successfully returned no error!")
				return
			}
		})
	}
}

func TestHeaderIs(t *testing.T) {
	t.Errorf("not implemented")
}

func TestHeaderAuthorizationIs(t *testing.T) {
	t.Errorf("not implemented")
}

func TestHeaderContentTypeIs(t *testing.T) {
	t.Errorf("not implemented")
}

func TestHeaderContentTypeJSON(t *testing.T) {
	t.Errorf("not implemented")
}

func TestHeaderContentTypeXML(t *testing.T) {
	t.Errorf("not implemented")
}

func TestHeaderMatches(t *testing.T) {
	t.Errorf("not implemented")
}

func TestHeaderAuthorizationMatchesBasic(t *testing.T) {
	t.Errorf("not implemented")
}

func TestHeaderAuthorizationMatchesBearer(t *testing.T) {
	t.Errorf("not implemented")
}
