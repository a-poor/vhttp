// Package vhttp contains helper functions for testing aspects of http.Request or
// http.Response objects from the net/http package.
//
// The core functions used are ValidateRequest and ValidateResponse whose arguments
// (in addition to the http.Request or http.Response) implement the RequestValidator
// ResponseValidator interfaces, respectively.
//
// This package also includes helper types for validating specific aspects of the
// requests or responses, such as the MethodValidator, HeaderValidator, or the
// BodyValidator.
//
// Some helper functions exist for creating validator functions but others can be
// created manually.
package vhttp
