package vhttp

import (
	"fmt"
	"net/http"
)

// StatusCodeValidator is a function that validates an http.Response's status code.
type StatusCodeValidator func(int) error

func (v StatusCodeValidator) ValidateResponse(res *http.Response) error {
	return v(res.StatusCode)
}

// StatusIs checks that the status code is equal to the given code.
func StatusIs(code int) StatusCodeValidator {
	return func(c int) error {
		if c != code {
			return fmt.Errorf("expected status code is %d, got %d", code, c)
		}
		return nil
	}
}

// StatusIsNot checks that the status code is not equal to the given code.
func StatusIsNot(code int) StatusCodeValidator {
	return func(c int) error {
		if c == code {
			return fmt.Errorf("expected status code to not be %d", code)
		}
		return nil
	}
}

// StatusIsOK checks that the status code is 200.
func StatusIsOK() StatusCodeValidator {
	return StatusIs(http.StatusOK)
}

// StatusInRange checks that the status code is in the given range: [min, max).
func StatusInRange(min, max int) StatusCodeValidator {
	return func(c int) error {
		if c < min || c >= max {
			return fmt.Errorf("expected status code to be in range [%d, %d), got %d", min, max, c)
		}
		return nil
	}
}

// StatusNotInRange checks that the status code is not in the given range: [min, max).
func StatusNotInRange(min, max int) StatusCodeValidator {
	return func(c int) error {
		if c >= min && c < max {
			return fmt.Errorf("expected status code to not be in range [%d, %d)", min, max)
		}
		return nil
	}
}

// StatusIs1XX checks that the status code is in the range [100, 200).
func StatusIs1XX() StatusCodeValidator {
	return StatusInRange(100, 200)
}

// StatusIs2XX checks that the status code is in the range [200, 300).
func StatusIs2XX() StatusCodeValidator {
	return StatusInRange(200, 300)
}

// StatusIs3XX checks that the status code is in the range [300, 400).
func StatusIs3XX() StatusCodeValidator {
	return StatusInRange(300, 400)
}

// StatusIs4XX checks that the status code is in the range [400, 500).
func StatusIs4XX() StatusCodeValidator {
	return StatusInRange(400, 500)
}

// StatusIs5XX checks that the status code is in the range [500, 600).
func StatusIs5XX() StatusCodeValidator {
	return StatusInRange(500, 600)
}

// StatusIsError checks that the status code is not in the range [400, 600).
func StatusNotError() StatusCodeValidator {
	return StatusNotInRange(400, 600)
}
