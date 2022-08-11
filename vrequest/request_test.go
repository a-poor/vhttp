package vrequest_test

import (
	"net/http"
	"testing"

	"github.com/a-poor/vhttp/vrequest"
)

func TestValidateRequest_Smoke(t *testing.T) {
	req := &http.Request{
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
	err := vrequest.ValidateRequest(
		req,
		vrequest.HasHeader("Content-Type"),
	)
	if err != nil {
		t.Errorf("expected no error, got: %s", err)
	}
}
