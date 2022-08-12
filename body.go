package vhttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-multierror"
)

// BodyValidator is a validator that validates an http.Request's body.
//
// Note that this expects the body to be fully read as a byte slice.
// If more than one BodyValidator is being used, you should use a
// CachedBodyValidator instead – which will read the body once and
// pass the resulting byte slice to all of the BodyValidators.
type BodyValidator func(b []byte) error

func (v BodyValidator) ValidateRequest(req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return InternalErr(fmt.Errorf("failed to read request body: %s", err))
	}
	return v(b)
}

func (v BodyValidator) ValidateResponse(res *http.Response) error {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return InternalErr(fmt.Errorf("failed to read response body: %s", err))
	}
	return v(b)
}

// CachedBodyValidator is a RequestValidator/ResponseValidator that reads the
// Request or Response body once and passes the byte slice to each of it's
// BodyValidators (rather than calling their ValidateRequest or ValidateResponse
// methods each time).
//
// This helps avoid duplicated reads of the body and prevents issues with
// attempts to read the body after it has been closed.
type CachedBodyValidator struct {
	vs []BodyValidator
}

// CacheBody creates a new CachedBodyValidator
func CacheBody(vs ...BodyValidator) CachedBodyValidator {
	return CachedBodyValidator{vs}
}

func (v CachedBodyValidator) ValidateRequest(req *http.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return InternalErr(fmt.Errorf("failed to read request body: %s", err))
	}

	var merr *multierror.Error
	for _, v := range v.vs {
		if err := v(b); err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	return merr.ErrorOrNil()
}

func (v CachedBodyValidator) ValidateResponse(res *http.Response) error {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return InternalErr(fmt.Errorf("failed to read response body: %s", err))
	}

	var merr *multierror.Error
	for _, v := range v.vs {
		if err := v(b); err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	return merr.ErrorOrNil()
}

// BodyIs validates that the body is equal to the given byte slice.
func BodyIs(b []byte) BodyValidator {
	return func(b2 []byte) error {
		if !bytes.Equal(b, b2) {
			return fmt.Errorf("body is not equal")
		}
		return nil
	}
}

// BodyIsString validates that the body is equal to the given string.
func BodyIsString(s string) BodyValidator {
	return func(b []byte) error {
		if string(b) != s {
			return fmt.Errorf("body is not equal")
		}
		return nil
	}
}

// BodyIsValidJSON uses the json.Valid function (from the encoding/json) to test
// if the body is valid JSON.
func BodyIsValidJSON() BodyValidator {
	return func(b []byte) error {
		if !json.Valid(b) {
			return fmt.Errorf("body is not valid JSON")
		}
		return nil
	}
}

// BodyLengthIs validates that the body has the given length n.
func BodyLengthIs(n int) BodyValidator {
	return func(b []byte) error {
		if m := len(b); m != n {
			return fmt.Errorf("expected body length to be %d, got %d", n, m)
		}
		return nil
	}
}

// BodyIsNil validates that the body is nil.
func BodyIsNil() BodyValidator {
	return func(b []byte) error {
		if b != nil {
			return fmt.Errorf("body is not nil")
		}
		return nil
	}
}

// BodyDetectedTypeIs uses the http.DetectContentType function to guess the content
// type of the body and returns an error if it does not match the expected type t.
func BodyDetectedTypeIs(t string) BodyValidator {
	return func(b []byte) error {
		if res := http.DetectContentType(b); res != t {
			return fmt.Errorf("body detected type is not %s", t)
		}
		return nil
	}
}

// BodyJSONUnmarshalsAs attmepts to uses the json.Unmarshal function to unmarshal
// the body. If it fails, an error is returned.
func BodyJSONUnmarshalsAs(v any) BodyValidator {
	return func(b []byte) error {
		if err := json.Unmarshal(b, v); err != nil {
			return fmt.Errorf("body JSON unmarshal failed: %s", err)
		}
		return nil
	}
}

// BodyXMLUnmarshalsAs attmepts to uses the xml.Unmarshal function to unmarshal
// the body. If it fails, an error is returned.
func BodyXMLUnmarshalsAs(v any) BodyValidator {
	return func(b []byte) error {
		if err := xml.Unmarshal(b, v); err != nil {
			return fmt.Errorf("body JSON unmarshal failed: %s", err)
		}
		return nil
	}
}
