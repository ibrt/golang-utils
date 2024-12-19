//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -source ./fixtures.go -destination ./mocks.gen.go -package terrorz

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
	_ TestDetailedError = (*TestDetailedErrorImpl)(nil)
)

// TestDetailedErrorImpl is a simple implementation of [TestDetailedError] for test purposes.
type TestDetailedErrorImpl struct {
	ErrorMessage string
	Name         string
	HTTPStatus   int
	Details      map[string]any
}

// Error implements the [TestDetailedError] interface.
func (e *TestDetailedErrorImpl) Error() string {
	return e.ErrorMessage
}

// GetErrorName implements the [TestDetailedError] interface.
func (e *TestDetailedErrorImpl) GetErrorName() string {
	return e.Name
}

// GetErrorHTTPStatus implements the [TestDetailedError] interface.
func (e *TestDetailedErrorImpl) GetErrorHTTPStatus() int {
	return e.HTTPStatus
}

// GetErrorDetails implements the [TestDetailedError] interface.
func (e *TestDetailedErrorImpl) GetErrorDetails() map[string]any {
	return e.Details
}

// TestDetailedErrorUnwrapSingle is a mock blueprint.
type TestDetailedErrorUnwrapSingle interface {
	TestDetailedError
	errorz.UnwrapSingle
}

var (
	_ TestDetailedErrorUnwrapSingle = (*TestDetailedErrorUnwrapSingleImpl)(nil)
)

// TestDetailedErrorUnwrapSingleImpl is a simple implementation of [TestDetailedErrorUnwrapSingle] for test purposes.
type TestDetailedErrorUnwrapSingleImpl struct {
	*TestDetailedErrorImpl
	UnwrapSingle error
}

// Unwrap implements the [TestDetailedErrorUnwrapSingle] interface.
func (e *TestDetailedErrorUnwrapSingleImpl) Unwrap() error {
	return e.UnwrapSingle
}

// TestDetailedErrorUnwrapMulti is a mock blueprint.
type TestDetailedErrorUnwrapMulti interface {
	TestDetailedError
	errorz.UnwrapMulti
}

var (
	_ TestDetailedErrorUnwrapMulti = (*TestDetailedErrorUnwrapMultiImpl)(nil)
)

// TestDetailedErrorUnwrapMultiImpl is a simple implementation of [TestDetailedErrorUnwrapMulti] for test purposes.
type TestDetailedErrorUnwrapMultiImpl struct {
	*TestDetailedErrorImpl
	UnwrapMulti []error
}

// Unwrap implements the [TestDetailedErrorUnwrapMulti] interface.
func (e *TestDetailedErrorUnwrapMultiImpl) Unwrap() []error {
	return e.UnwrapMulti
}

// TestStringError is a test error.
type TestStringError string

// Error implements the [error] interface.
func (s TestStringError) Error() string {
	return string(s)
}
