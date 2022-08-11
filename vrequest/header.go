package vrequest

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/a-poor/vhttp"
)

// HasHeader creates a request validator that checks that the header h
// is present in the request object.
//
// Note that this function will convert the header key to canonical form
// using the vhttp.CanonicalHeaderKey function. If this behavior needs to
// be changed, either create a custom validator or change the value of
// the function.
func HasHeader(h string) RequestValidator {
	return func(r *http.Request) error {
		// Convert the header key to canonical form.
		h := vhttp.CanonicalHeaderKey(h)

		// Check if the header is present.
		if _, ok := r.Header[h]; !ok {
			return fmt.Errorf("header %q not found", h)
		}

		// Found!
		return nil
	}
}

// HasHeaderContentType creates a request validator that checks that the
// "Content-Type" header is present in the request object.
func HasHeaderContentType(ct string) RequestValidator {
	return HasHeader("Content-Type")
}

// HasHeaderAccept creates a request validator that checks that the header
// "Accept" is present in the request object.
func HasHeaderAccept() RequestValidator {
	return HasHeader("Accept")
}

// HasHeaderAuthorization creates a request validator that checks that the
// "Authorization" header is present in the request object.
func HasHeaderAuthorization() RequestValidator {
	return HasHeader("Authorization")
}

// HeaderIs creates a request validator that checks that at least one of
// the values for header h is equal to v.
//
// Note that this function will convert the header key to canonical form
// using the vhttp.CanonicalHeaderKey function. If this behavior needs to
// be changed, either create a custom validator or change the value of
// the function.
func HeaderIs(h, v string) RequestValidator {
	return func(req *http.Request) error {
		// Convert the header key to canonical form.
		h := vhttp.CanonicalHeaderKey(h)

		// Get the header values
		vs, ok := req.Header[h]
		if !ok {
			return fmt.Errorf("header %q not found", h)
		}

		// Check if the header is present.
		for _, s := range vs {
			if s == v {
				return nil // Found!
			}
		}

		// Not found.
		return fmt.Errorf("expected header %q to have value %q", h, v)
	}
}

// HeaderAuthorizationIs creates a request validator that checks that at least
// one of the "Authorization" header values are equal to t.
func HeaderAuthorizationIs(t string) RequestValidator {
	return HeaderIs("Authorization", t)
}

// HeaderContentTypeIs creates a request validator that checks that at least
// one of the "Content-Type" header values are equal to t.
func HeaderContentTypeIs(ct string) RequestValidator {
	return HeaderIs("Content-Type", ct)
}

// HeaderContentTypeJSON creates a request validator that checks that at
// least one of the "Content-Type" header values are equal to "application/json".
func HeaderContentTypeJSON() RequestValidator {
	return HeaderIs("Content-Type", "application/json")
}

// HeaderContentTypeXML creates a request validator that checks that at
// least one of the "Content-Type" header values are equal to "application/xml".
func HeaderContentTypeXML() RequestValidator {
	return HeaderIs("Content-Type", "application/xml")
}

// HeaderMatches creates a request validator that checks that at least one
// of the request header values matche the given regular expression.
//
// Note that this function will convert the header key to canonical form
// using the vhttp.CanonicalHeaderKey function. If this behavior needs to
// be changed, either create a custom validator or change the value of
// the function.
func HeaderMatches(h string, re *regexp.Regexp) RequestValidator {
	return func(req *http.Request) error {
		// Convert the header key to canonical form.
		h := vhttp.CanonicalHeaderKey(h)

		// Get the header values
		vs, ok := req.Header[h]
		if !ok {
			return fmt.Errorf("header %q not found", h)
		}

		// Check if the header is present.
		for _, s := range vs {
			if re.MatchString(s) {
				return nil // Found!
			}
		}

		// Not found.
		return fmt.Errorf("expected header %q to match %q", h, re)
	}
}

// HeaderAuthorizationMatchesBasic creates a request validator that checks that
// the request header value for the Authorization header matches the regular
// expression for a basic authentication header.
//
// The regular expression is defined in the variable vhttp.BasicAuthMatch
// as `^Basic .*`.
func HeaderAuthorizationMatchesBasic() RequestValidator {
	return HeaderMatches("Authorization", vhttp.BasicAuthMatch)
}

// HeaderAuthorizationMatchesBearer creates a request validator that checks that
// the request header value for the Authorization header matches the regular
// expression for a bearer authentication token.
//
// The regular expression is defined in the variable vhttp.BearerAuthMatch
// as `^Bearer .*`.
func HeaderAuthorizationMatchesBearer() RequestValidator {
	return HeaderMatches("Authorization", vhttp.BearerAuthMatch)
}
