package vrequest

import (
	"fmt"
	"net/http"
	"path"
)

// URLIs creates a request validator that checks that the request
// URL exactly matches the given URL u.
func URLIs(u string) RequestValidator {
	return func(req *http.Request) error {
		if req.URL.String() != u {
			return fmt.Errorf("expected URL %q, found %q", u, req.URL.String())
		}
		return nil
	}
}

func URLSchemeIs(s string) RequestValidator {
	return func(req *http.Request) error {
		if req.URL.Scheme != s {
			return fmt.Errorf("expected URL scheme %q, found %q", s, req.URL.Scheme)
		}
		return nil
	}
}

func URLSchemeIsHTTP() RequestValidator {
	return URLSchemeIs("http")
}

func URLSchemeIsHTTPS() RequestValidator {
	return URLSchemeIs("https")
}

func URLPathIs(p string) RequestValidator {
	return func(req *http.Request) error {
		if req.URL.Path != p {
			return fmt.Errorf("expected URL path %q, found %q", p, req.URL.Path)
		}
		return nil
	}
}

// URLUserinfoIs creates a request validator that checks that the request
// URL's userinfo matches the given userinfo u.
//
// The user info will be in the form of "username[:password]".
func URLUserinfoIs(u string) RequestValidator {
	return func(req *http.Request) error {
		if req.URL.User.String() != u {
			return fmt.Errorf("expected URL userinfo %q, found %q", u, req.URL.User.String())
		}
		return nil
	}
}

func URLHostIs(h string) RequestValidator {
	return func(req *http.Request) error {
		if req.URL.Host != h {
			return fmt.Errorf("expected URL host %q, found %q", h, req.URL.Host)
		}
		return nil
	}
}

// URLPathGlob creates a request validator that checks that the request
// URL's path matches the given glob pattern p.
//
// Uses path.Match to match the glob pattern.
func URLPathGlob(p string) RequestValidator {
	return func(req *http.Request) error {
		// Match the pattern against the path (& check for error)
		m, err := path.Match(p, req.URL.Path)
		if err != nil {
			return fmt.Errorf("error matching path against pattern %q: %v", p, err)
		}

		// If the path does not match, return an error
		if !m {
			return fmt.Errorf("path %q does not match pattern %q", req.URL.Path, p)
		}
		return nil
	}
}

func URLQueryHas(k string) RequestValidator {
	return func(req *http.Request) error {
		if !req.URL.Query().Has(k) {
			return fmt.Errorf("expected URL query %q, found none", k)
		}
		return nil
	}
}

func URLQueryIs(k, v string) RequestValidator {
	return func(req *http.Request) error {
		// Get the list of values for the given key
		vs := req.URL.Query()[k]

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

// URLQueryValueValidator creates a request validator that applies the validator
// function vfn to the first value for the given key in the request URL's query.
func URLQueryValueValidator(k string, vfn func(string) error) RequestValidator {
	return func(req *http.Request) error {
		// Get the list of values for the given key
		v := req.URL.Query().Get(k)

		// Run the validator function
		err := vfn(v)
		if err != nil {
			return fmt.Errorf("error validating URL query %q=%q: %v", k, v, err)
		}
		return nil
	}
}
