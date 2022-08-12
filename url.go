package vhttp

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

// URLValidator is a validator function that validates an http.Request's
// URL field.
type URLValidator func(*url.URL) error

func (v URLValidator) ValidateRequest(req *http.Request) error {
	return v(req.URL)
}

// URLIs creates a url validator that checks that the
// URL exactly matches the given string s.
func URLIs(s string) URLValidator {
	return func(u *url.URL) error {
		if u.String() != s {
			return fmt.Errorf("expected URL %q, found %q", s, u.String())
		}
		return nil
	}
}

// URLSchemeIs creates a URLValidator that checks that the request
// URL's scheme matches the given scheme s.
func URLSchemeIs(s string) URLValidator {
	return func(u *url.URL) error {
		if u.Scheme != s {
			return fmt.Errorf("expected URL scheme %q, found %q", s, u.Scheme)
		}
		return nil
	}
}

// URLSchemeIsHTTP creates a URLValidator that checks that the request
// URL's scheme is "http".
func URLSchemeIsHTTP() URLValidator {
	return URLSchemeIs("http")
}

// URLSchemeIsHTTPS creates a URLValidator that checks that the request
// URL's scheme is "https".
func URLSchemeIsHTTPS() URLValidator {
	return URLSchemeIs("https")
}

// URLPathIs creates a URLValidator that checks that the request URL's
// path is equal to the given path p.
func URLPathIs(p string) URLValidator {
	return func(u *url.URL) error {
		if u.Path != p {
			return fmt.Errorf("expected URL path %q, found %q", p, u.Path)
		}
		return nil
	}
}

// URLUserinfoIs creates a url validator that checks that the
// URL's userinfo matches the given userinfo ui.
//
// The user info will be in the form of "username[:password]".
func URLUserinfoIs(ui string) URLValidator {
	return func(u *url.URL) error {
		if u.User.String() != ui {
			return fmt.Errorf("expected URL userinfo %q, found %q", ui, u.User.String())
		}
		return nil
	}
}

// URLHostIs creates a URL validator that checks that the
// URL's Host field matches h.
func URLHostIs(h string) URLValidator {
	return func(u *url.URL) error {
		if u.Host != h {
			return fmt.Errorf("expected URL host %q, found %q", h, u.Host)
		}
		return nil
	}
}

// URLPathGlob creates a URL validator that checks that the
// URL's path matches the given glob pattern p.
//
// Uses path.Match to match the glob pattern.
func URLPathGlob(p string) URLValidator {
	return func(u *url.URL) error {
		// Match the pattern against the path (& check for error)
		m, err := path.Match(p, u.Path)
		if err != nil {
			return InternalErr(fmt.Errorf("error matching path against pattern %q: %v", p, err))
		}

		// If the path does not match, return an error
		if !m {
			return fmt.Errorf("path %q does not match pattern %q", u.Path, p)
		}
		return nil
	}
}

// URLQueryHas creates a URLValidator that checks that the request
// URL's query parameters contain the given key k.
func URLQueryHas(k string) URLValidator {
	return func(u *url.URL) error {
		if !u.Query().Has(k) {
			return fmt.Errorf("expected value for URL query key %q to be present", k)
		}
		return nil
	}
}

// URLQueryIs creates a URLValidator that checks that the request
// URL's query parameters contain the given key k with the value v.
func URLQueryIs(k, v string) URLValidator {
	return func(u *url.URL) error {
		// Get the list of values for the given key
		vs := u.Query()[k]

		// For each value...
		for _, s := range vs {
			// If the value is the one we're looking for, success!
			if s == v {
				return nil
			}
		}

		// Not found...
		return fmt.Errorf("expected at least one value for URL query %q to be %q", k, v)
	}
}

// URLQueryValueValidator creates a URLValidator that applies the validator
// function vfn to the first value for the given key in the request URL's query.
func URLQueryValueValidator(k string, vfn func(string) error) URLValidator {
	return func(u *url.URL) error {
		// Get the list of values for the given key
		v := u.Query().Get(k)

		// Run the validator function
		err := vfn(v)
		if err != nil {
			return fmt.Errorf("error validating URL query %q=%q: %v", k, v, err)
		}
		return nil
	}
}
