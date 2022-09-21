package vhttp_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/a-poor/vhttp"
)

// t.Run("", func(t *testing.T) {})

func TestValidateRequest(t *testing.T) {
	t.Run("non-nil-request", func(t *testing.T) {
		err := vhttp.ValidateRequest(&http.Request{})
		if err != nil {
			t.Errorf("Unexpected error validating non-nil request with no validators: %s", err)
		}
	})
	t.Run("nil-request", func(t *testing.T) {
		err := vhttp.ValidateRequest(nil)
		if err == nil {
			t.Error("Expected an error to be returned when validating a nil request")
		}
	})
	t.Run("doesnt-fail-fast", func(t *testing.T) {
		// Create a request
		req := &http.Request{}

		// Create three validator functions and corresponding flags
		// to track if they've run.
		//
		// The second function will return an error.
		//
		// Since this function doesn't fail fast, all three should run.
		runA, runB, runC := false, false, false
		fa := vhttp.RequestFunc(func(req *http.Request) error {
			runA = true
			return nil
		})
		fb := vhttp.RequestFunc(func(req *http.Request) error {
			runB = true
			return errors.New("an error!")
		})
		fc := vhttp.RequestFunc(func(req *http.Request) error {
			runC = true
			return nil
		})

		// Validate
		err := vhttp.ValidateRequest(req, fa, fb, fc)

		// An error should have been returned
		if err == nil {
			t.Error("expected an error to be returned")
		}

		// Check which functions have run
		if !runA || !runB || !runC {
			t.Errorf("Expected all three functions to have run. RunA=%t, RunB=%t, RunC=%t", runA, runB, runC)
		}
	})
}

func TestValidateRequestFF(t *testing.T) {
	t.Run("non-nil-request", func(t *testing.T) {
		err := vhttp.ValidateRequestFF(&http.Request{})
		if err != nil {
			t.Errorf("Unexpected error validating non-nil request with no validators: %s", err)
		}
	})
	t.Run("nil-request", func(t *testing.T) {
		err := vhttp.ValidateRequestFF(nil)
		if err == nil {
			t.Error("Expected an error to be returned when validating a nil request")
		}
	})
	t.Run("fails-fast", func(t *testing.T) {
		// Create a request
		req := &http.Request{}

		// Create three validator functions and corresponding flags
		// to track if they've run.
		//
		// The second function will return an error.
		//
		// Since this function fails fast, only the first two should run.
		runA, runB, runC := false, false, false
		fa := vhttp.RequestFunc(func(req *http.Request) error {
			runA = true
			return nil
		})
		fb := vhttp.RequestFunc(func(req *http.Request) error {
			runB = true
			return errors.New("an error!")
		})
		fc := vhttp.RequestFunc(func(req *http.Request) error {
			runC = true
			return nil
		})

		// Validate
		err := vhttp.ValidateRequestFF(req, fa, fb, fc)

		// An error should have been returned
		if err == nil {
			t.Error("expected an error to be returned")
		}

		// Check which functions have run
		if !runA || !runB {
			t.Errorf("Expected the first two functions to have run. RunA=%t, RunB=%t", runA, runB)
		}
		if runC {
			t.Error("Expected the third function to have been skipped.")
		}
	})
}

func TestValidateResponse(t *testing.T) {
	t.Run("non-nil-response", func(t *testing.T) {
		err := vhttp.ValidateResponse(&http.Response{})
		if err != nil {
			t.Errorf("Unexpected error validating non-nil request with no validators: %s", err)
		}
	})
	t.Run("nil-response", func(t *testing.T) {
		err := vhttp.ValidateResponse(nil)
		if err == nil {
			t.Error("Expected an error to be returned when validating a nil response")
		}
	})
	t.Run("doesnt-fail-fast", func(t *testing.T) {
		// Create a response
		req := &http.Response{}

		// Create three validator functions and corresponding flags
		// to track if they've run.
		//
		// The second function will return an error.
		//
		// Since this function doesn't fail fast, all three should run.
		runA, runB, runC := false, false, false
		fa := vhttp.ResponseFunc(func(req *http.Response) error {
			runA = true
			return nil
		})
		fb := vhttp.ResponseFunc(func(req *http.Response) error {
			runB = true
			return errors.New("an error!")
		})
		fc := vhttp.ResponseFunc(func(req *http.Response) error {
			runC = true
			return nil
		})

		// Validate
		err := vhttp.ValidateResponse(req, fa, fb, fc)

		// An error should have been returned
		if err == nil {
			t.Error("expected an error to be returned")
		}

		// Check which functions have run
		if !runA || !runB || !runC {
			t.Errorf("Expected all three functions to have run. RunA=%t, RunB=%t, RunC=%t", runA, runB, runC)
		}
	})
}

func TestValidateResponseFF(t *testing.T) {
	t.Run("non-nil-response", func(t *testing.T) {
		err := vhttp.ValidateResponseFF(&http.Response{})
		if err != nil {
			t.Errorf("Unexpected error validating non-nil request with no validators: %s", err)
		}
	})
	t.Run("nil-response", func(t *testing.T) {
		err := vhttp.ValidateResponseFF(nil)
		if err == nil {
			t.Error("Expected an error to be returned when validating a nil response")
		}
	})
	t.Run("fails-fast", func(t *testing.T) {
		// Create a response
		req := &http.Response{}

		// Create three validator functions and corresponding flags
		// to track if they've run.
		//
		// The second function will return an error.
		//
		// Since this function fails fast, only the first two should run.
		runA, runB, runC := false, false, false
		fa := vhttp.ResponseFunc(func(req *http.Response) error {
			runA = true
			return nil
		})
		fb := vhttp.ResponseFunc(func(req *http.Response) error {
			runB = true
			return errors.New("an error!")
		})
		fc := vhttp.ResponseFunc(func(req *http.Response) error {
			runC = true
			return nil
		})

		// Validate
		err := vhttp.ValidateResponseFF(req, fa, fb, fc)

		// An error should have been returned
		if err == nil {
			t.Error("expected an error to be returned")
		}

		// Check which functions have run
		if !runA || !runB {
			t.Errorf("Expected the first two functions to have run. RunA=%t, RunB=%t", runA, runB)
		}
		if runC {
			t.Error("Expected the third function to have been skipped.")
		}
	})
}
