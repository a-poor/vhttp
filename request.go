package resttest

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CanonicallHeaderKey converts the given string to a canonical form.
//
// This function is used in the Header function, on the keys of the
// HeaderValidators parameter. By copying the function from the net/http
// package, this can be replaced with a custom implementation, if that
// functionality needs to be changed (eg to an identity function).
var CanonicalHeaderKey func(string) string = http.CanonicalHeaderKey

// RequestValidator is a function that returns a validation struct
// field validator for the given request.
type RequestValidator func(*http.Request) *validation.FieldRules

// HeaderValidators is a map from header keys to a slice of
// validators for the key.
type HeaderValidators map[string][]validation.Rule

// ValidateRequest validates the given HTTP request object
// against the given rules.
func ValidateRequest(req *http.Request, v ...RequestValidator) error {
	rs := make([]*validation.FieldRules, len(v))
	for i, r := range v {
		rs[i] = r(req)
	}
	return validation.ValidateStruct(req, rs...)
}

// MethodIs validates that the request's HTTP method is not empty and
// matches one of the given methods.
func MethodIs(m ...string) RequestValidator {
	ms := toAnySlice(m)
	return func(r *http.Request) *validation.FieldRules {
		return validation.Field(&r.Method, validation.Required, validation.In(ms...))
	}
}

// MethodIsGet validates that the request's HTTP method is GET.
func MethodIsGet() RequestValidator {
	return MethodIs(http.MethodGet)
}

// MethodIsPost validates that the request's HTTP method is POST.
func MethodIsPost() RequestValidator {
	return MethodIs(http.MethodPost)
}

// MethodIsPut validates that the request's HTTP method is PUT.
func MethodIsPut() RequestValidator {
	return MethodIs(http.MethodPut)
}

// MethodIsDelete validates that the request's HTTP method is DELETE.
func MethodIsDelete() RequestValidator {
	return MethodIs(http.MethodDelete)
}

// MethodIsPatch validates that the request's HTTP method is PATCH.
func MethodIsPatch() RequestValidator {
	return MethodIs(http.MethodPatch)
}

// MethodIsOptions validates that the request's HTTP method is OPTIONS.
func MethodIsOptions() RequestValidator {
	return MethodIs(http.MethodOptions)
}

// MethodIsnt validates that the request's HTTP method is not empty and is
// not one of the given methods.
func MethodIsnt(m ...string) RequestValidator {
	ms := toAnySlice(m)
	return func(r *http.Request) *validation.FieldRules {
		return validation.Field(&r.Method, validation.Required, validation.NotIn(ms...))
	}
}

func Demo() RequestValidator {
	// Create a header validator
	hv := Header(HeaderValidators{
		HeaderContentType:   {validation.Required, validation.In(MimeJSON)},
		HeaderAuthorization: {validation.Required},
	})
	return hv
}

// Header creates a validator that makes assertions about
// the request's Header field.
//
//   hv := Header(HeaderValidators{
// 	   HeaderContentType:   {validation.Required, validation.In(MimeJSON)},
//     HeaderAuthorization: {validation.Required},
//   })
//
// Note that keys will be converted to the canonical form
// using the CanonicalHeaderKey function of this package
// (copied from the net/http package).
func Header(rules HeaderValidators) RequestValidator {
	// Slice of any specified key validators
	var hr []*validation.KeyRules
	for k, v := range rules {
		// Make sure the header is in canonical form
		h := CanonicalHeaderKey(k)

		// Apply the rule to each header value
		rs := validation.Each(v...)

		// Add the key rule to the slice
		hr = append(hr, validation.Key(any(h), rs))
	}

	// Return a function that applies the rules from to the request
	return func(r *http.Request) *validation.FieldRules {
		return validation.Field(&r.Header, validation.Map(hr...))
	}
}
