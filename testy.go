package resty

import (
	"net/http"
)

type RequestCheck func(*http.Request) error

type ResponseCheck func(*http.Response) error
