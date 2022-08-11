package vhttp

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

// TLSValidator is a validator that validates an http.Request or http.Response
// object's TLS connection.
type TLSValidator func(*tls.ConnectionState) error

func (v TLSValidator) ValidateRequest(req *http.Request) error {
	return v(req.TLS)
}

func (v TLSValidator) ValidateResponse(res *http.Response) error {
	return v(res.TLS)
}

func TLSIsNil() TLSValidator {
	return func(tls *tls.ConnectionState) error {
		if tls != nil {
			return fmt.Errorf("tls is not nil")
		}

		return nil
	}
}

func TLSIsNotNil() TLSValidator {
	return func(tls *tls.ConnectionState) error {
		if tls != nil {
			return fmt.Errorf("tls is nil")
		}

		return nil
	}
}

func TLSVersionIs(v uint16) TLSValidator {
	return func(tls *tls.ConnectionState) error {
		if tls.Version != v {
			return fmt.Errorf("tls version is not %d", v)
		}

		return nil
	}
}
