package vhttp_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/a-poor/vhttp"
)

func TestURLIs(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		// Test a working case
		s := "https://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLIs(s).
			ValidateRequest(req)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})
	t.Run("bad", func(t *testing.T) {
		// Test a working case
		s := "https://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLIs("http://example.com").
			ValidateRequest(req)
		if err == nil {
			t.Errorf("expected error but none returned")
		}
	})
}

func TestURLSchemeIs(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		// Test a working case
		s := "https://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}
		in := "https"

		err := vhttp.
			URLSchemeIs(in).
			ValidateRequest(req)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})
	t.Run("bad", func(t *testing.T) {
		// Test a working case
		s := "http://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}
		in := "https"

		err := vhttp.
			URLSchemeIs(in).
			ValidateRequest(req)
		if err == nil {
			t.Errorf("expected error but none returned")
		}
	})
}

func TestURLSchemeIsHTTP(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		// Test a working case
		s := "http://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLSchemeIsHTTP().
			ValidateRequest(req)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})
	t.Run("bad", func(t *testing.T) {
		// Test a working case
		s := "https://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLSchemeIsHTTP().
			ValidateRequest(req)
		if err == nil {
			t.Errorf("expected error but none returned")
		}
	})
}

func TestURLSchemeIsHTTPS(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		// Test a working case
		s := "HTTPS://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLSchemeIsHTTPS().
			ValidateRequest(req)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})
	t.Run("bad", func(t *testing.T) {
		// Test a working case
		s := "http://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLSchemeIsHTTPS().
			ValidateRequest(req)
		if err == nil {
			t.Errorf("expected error but none returned")
		}
	})
}

func TestURLPathIs(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		// Test a working case
		s := "HTTPS://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLPathIs("/api/v1/users").
			ValidateRequest(req)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})
	t.Run("bad", func(t *testing.T) {
		// Test a working case
		s := "http://example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLPathIs("/").
			ValidateRequest(req)
		if err == nil {
			t.Errorf("expected error but none returned")
		}
	})
}

func TestURLUserinfoIs(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		// Test a working case
		s := "https://user:pass@example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLUserinfoIs("user:pass").
			ValidateRequest(req)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})
	t.Run("just-user", func(t *testing.T) {
		// Test a working case
		s := "http://user@example.com/api/v1/users"
		u, _ := url.Parse(s)
		req := &http.Request{URL: u}

		err := vhttp.
			URLUserinfoIs("user").
			ValidateRequest(req)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})
}

func TestURLHostIs(t *testing.T) {}

func TestURLPathGlob(t *testing.T) {}

func TestURLQueryHas(t *testing.T) {}

func TestURLQueryIs(t *testing.T) {}

func TestURLQueryValueValidator(tt *testing.T) {}
