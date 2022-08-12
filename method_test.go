package vhttp_test

import (
	"net/http"
	"testing"

	"github.com/a-poor/vhttp"
)

func TestMethodValidator(t *testing.T) {
	var called bool
	v := vhttp.MethodValidator(func(m string) error {
		called = true
		return nil
	})
	v.ValidateRequest((&http.Request{}))

	if !called {
		t.Errorf("expected method validator to be called")
	}
}

func TestMethodIs(t *testing.T) {
	cases := []struct {
		name   string // Name of the test case
		method string // Request method
		expect string // Expected method (passed to fn)
		valid  bool   // No error returned?
	}{
		{
			name:   "GET-good",
			method: http.MethodGet,
			expect: http.MethodGet,
			valid:  true,
		},
		{
			name:   "GET-bad",
			method: http.MethodPost,
			expect: http.MethodGet,
			valid:  false,
		},
		{
			name:   "POST-good",
			method: http.MethodPost,
			expect: http.MethodPost,
			valid:  true,
		},
		{
			name:   "custom-good",
			method: "HELLO-WORLD",
			expect: "HELLO-WORLD",
			valid:  true,
		},
		{
			name:   "empty-bad",
			method: http.MethodPut,
			expect: "",
			valid:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Create the validator function
			fn := vhttp.MethodIs(c.expect)

			// Create the request
			req := &http.Request{Method: c.method}

			// Validate the request
			err := fn.ValidateRequest(req)

			// Parse the response
			valid := err == nil

			// Check the error
			if valid != c.valid {
				t.Errorf("valid? %t!=%t. Error: %s", valid, c.valid, err)
			}
		})
	}
}

func TestMethodIsNot(t *testing.T) {
	cases := []struct {
		name   string // Name of the test case
		method string // Request method
		expect string // Expected method (passed to fn)
		valid  bool   // No error returned?
	}{
		{
			name:   "GET-bad",
			method: http.MethodGet,
			expect: http.MethodGet,
			valid:  false,
		},
		{
			name:   "GET-good",
			method: http.MethodPost,
			expect: http.MethodGet,
			valid:  true,
		},
		{
			name:   "POST-bad",
			method: http.MethodPost,
			expect: http.MethodPost,
			valid:  false,
		},
		{
			name:   "custom-bad",
			method: "HELLO-WORLD",
			expect: "HELLO-WORLD",
			valid:  false,
		},
		{
			name:   "empty-good",
			method: http.MethodPut,
			expect: "",
			valid:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Create the validator function
			fn := vhttp.MethodIsNot(c.expect)

			// Create the request
			req := &http.Request{Method: c.method}

			// Validate the request
			err := fn.ValidateRequest(req)

			// Parse the response
			valid := err == nil

			// Check the error
			if valid != c.valid {
				t.Errorf("valid? %t!=%t. Error: %s", valid, c.valid, err)
			}
		})
	}
}

func TestMethodIsGet(t *testing.T) {
	// Test IS get
	r := &http.Request{Method: http.MethodDelete}
	err := vhttp.MethodIsDelete().ValidateRequest(r)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	// Test is NOT get
	r = &http.Request{Method: http.MethodGet}
	err = vhttp.MethodIsDelete().ValidateRequest(r)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestMethodIsPost(t *testing.T) {
	// Test IS post
	r := &http.Request{Method: http.MethodPost}
	err := vhttp.MethodIsPost().ValidateRequest(r)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	// Test is NOT post
	r = &http.Request{Method: http.MethodGet}
	err = vhttp.MethodIsPost().ValidateRequest(r)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestMethodIsPut(t *testing.T) {
	// Test IS put
	r := &http.Request{Method: http.MethodPut}
	err := vhttp.MethodIsPut().ValidateRequest(r)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	// Test is NOT put
	r = &http.Request{Method: http.MethodGet}
	err = vhttp.MethodIsPut().ValidateRequest(r)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestMethodIsPatch(t *testing.T) {
	// Test IS patch
	r := &http.Request{Method: http.MethodPatch}
	err := vhttp.MethodIsPatch().ValidateRequest(r)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	// Test is NOT patch
	r = &http.Request{Method: http.MethodPut}
	err = vhttp.MethodIsPatch().ValidateRequest(r)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestMethodIsOptions(t *testing.T) {
	// Test IS options
	r := &http.Request{Method: http.MethodOptions}
	err := vhttp.MethodIsOptions().ValidateRequest(r)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	// Test is NOT options
	r = &http.Request{Method: http.MethodDelete}
	err = vhttp.MethodIsOptions().ValidateRequest(r)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestMethodIsDelete(t *testing.T) {
	// Test IS delete
	r := &http.Request{Method: http.MethodDelete}
	err := vhttp.MethodIsDelete().ValidateRequest(r)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	// Test is NOT delete
	r = &http.Request{Method: http.MethodGet}
	err = vhttp.MethodIsDelete().ValidateRequest(r)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
