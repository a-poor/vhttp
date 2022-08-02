package resttest

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	// A regex rule that matches a basic bearer "Authorization" header.
	// For example: "Bearer <token>"
	BearerAuthRegexRule = validation.Match(regexp.MustCompile(`^Bearer\s+(.+)$`))

	// A regex rule that matches a basic basic "Authorization" header.
	// For example: "Basic <token>"
	BasicAuthRegexRule = validation.Match(regexp.MustCompile(`^Basic\s+(.+)$`))
)
