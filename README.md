# vhttp

__created by Austin Poor_

A library for validating HTTP requests and responses from the `net/http` package.

## Quick Example

The following is a quick example testing that `req` (of type `*http.Request`)...
1. Is a "GET" request
2. Calling `json.Valid()` on the body returns true
3. Has the header "Content-Type" and it's equal to "application/json"
4. The header "Authorization" matches the regular expression `^Bearer .+$`
5. Has the URL path "/users/all"

```go
err := vhttp.ValidateRequest(req, 
    vhttp.MethodIsGet(),                      // 1
    vhttp.BodyIsValidJSON(),                  // 2
    vhttp.HeaderContentTypeJSON(),            // 3
    vhttp.HeaderAuthorizationMatchesBearer(), // 4
    vhttp.URLPathIs("/users/all"),            // 5
)
```

Read more and find more examples in the [go docs](https://pkg.go.dev/github.com/a-poor/vhttp)!

## To-Do

- Add lots of tests!
- Add examples!
- Form validators
- Check for `nil` pointers? (eg `*url.URL`)
