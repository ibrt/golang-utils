//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -write_package_comment=false -source ./fixtures.go -destination ./mocks.gen.go -package terrorz

// Package terrorz provides test fixtures for the "errorz" package.
package terrorz

import (
	"github.com/ibrt/golang-utils/errorz"
)

// TestDetailedError is a mock blueprint.
type TestDetailedError interface {
	error
	errorz.ErrorName
	errorz.ErrorHTTPStatus
	errorz.ErrorDetails
}

var (
	_ TestDetailedError = (*SimpleMockTestDetailedError)(nil)
)

// SimpleMockTestDetailedError is a simple implementation of [TestDetailedError] for test purposes.
type SimpleMockTestDetailedError struct {
	ErrorMessage string
	Name         string
	HTTPStatus   int
	Details      map[string]any
}

// Error implements the [TestDetailedError] interface.
func (e *SimpleMockTestDetailedError) Error() string {
	return e.ErrorMessage
}

// GetErrorName implements the [TestDetailedError] interface.
func (e *SimpleMockTestDetailedError) GetErrorName() string {
	return e.Name
}

// GetErrorHTTPStatus implements the [TestDetailedError] interface.
func (e *SimpleMockTestDetailedError) GetErrorHTTPStatus() int {
	return e.HTTPStatus
}

// GetErrorDetails implements the [TestDetailedError] interface.
func (e *SimpleMockTestDetailedError) GetErrorDetails() map[string]any {
	return e.Details
}

// TestDetailedUnwrapSingleError is a mock blueprint.
type TestDetailedUnwrapSingleError interface {
	TestDetailedError
	errorz.UnwrapSingle
}

var (
	_ TestDetailedUnwrapSingleError = (*SimpleMockTestDetailedUnwrapSingleError)(nil)
)

// SimpleMockTestDetailedUnwrapSingleError is a simple implementation of [TestDetailedUnwrapSingleError] for test purposes.
type SimpleMockTestDetailedUnwrapSingleError struct {
	*SimpleMockTestDetailedError
	UnwrapSingle error
}

// Unwrap implements the [TestDetailedUnwrapSingleError] interface.
func (e *SimpleMockTestDetailedUnwrapSingleError) Unwrap() error {
	return e.UnwrapSingle
}

// TestDetailedUnwrapMultiError is a mock blueprint.
type TestDetailedUnwrapMultiError interface {
	TestDetailedError
	errorz.UnwrapMulti
}

var (
	_ TestDetailedUnwrapMultiError = (*SimpleMockTestDetailedUnwrapMultiError)(nil)
)

// SimpleMockTestDetailedUnwrapMultiError is a simple implementation of [TestDetailedUnwrapMultiError] for test purposes.
type SimpleMockTestDetailedUnwrapMultiError struct {
	*SimpleMockTestDetailedError
	UnwrapMulti []error
}

// Unwrap implements the [TestDetailedUnwrapMultiError] interface.
func (e *SimpleMockTestDetailedUnwrapMultiError) Unwrap() []error {
	return e.UnwrapMulti
}

// TestStringError is a test error.
type TestStringError string

// Error implements the [error] interface.
func (s TestStringError) Error() string {
	return string(s)
}
