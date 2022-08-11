package vhttp

import "net/http"

// ProtoValidator is a validator that validates an http.Request or http.Response's
// Proto field.
type ProtoValidator func(string) error

func (v ProtoValidator) ValidateRequest(req *http.Request) error {
	return v(req.Proto)
}

func (v ProtoValidator) ValidateResponse(res *http.Response) error {
	return v(res.Proto)
}
