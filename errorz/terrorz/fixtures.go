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
	_ TestDetailedError = (*testDetailedError)(nil)
)

type testDetailedError struct {
	errorMessage string
	name         string
	httpStatus   int
	details      map[string]any
}

// NewTestDetailedError initializes a new [TestDetailedError].
func NewTestDetailedError(
	errorMessage string,
	name string,
	httpStatus int,
	details map[string]any,
) TestDetailedError {
	return &testDetailedError{
		errorMessage: errorMessage,
		name:         name,
		httpStatus:   httpStatus,
		details:      details,
	}
}

// Error implements the [TestDetailedError] interface.
func (e *testDetailedError) Error() string {
	return e.errorMessage
}

// GetErrorName implements the [TestDetailedError] interface.
func (e *testDetailedError) GetErrorName() string {
	return e.name
}

// GetErrorHTTPStatus implements the [TestDetailedError] interface.
func (e *testDetailedError) GetErrorHTTPStatus() int {
	return e.httpStatus
}

// GetErrorDetails implements the [TestDetailedError] interface.
func (e *testDetailedError) GetErrorDetails() map[string]any {
	return e.details
}

// TestDetailedErrorUnwrapSingle is a mock blueprint.
type TestDetailedErrorUnwrapSingle interface {
	error
	errorz.ErrorName
	errorz.ErrorHTTPStatus
	errorz.ErrorDetails
	errorz.UnwrapSingle
}

var (
	_ TestDetailedErrorUnwrapSingle = (*testDetailedErrorUnwrapSingle)(nil)
)

type testDetailedErrorUnwrapSingle struct {
	TestDetailedError
	unwrapSingle error
}

// NewTestDetailedErrorUnwrapSingle initializes a new [TestDetailedErrorUnwrapSingle].
func NewTestDetailedErrorUnwrapSingle(
	errorMessage string,
	name string,
	httpStatus int,
	details map[string]any,
	unwrapSingle error,
) TestDetailedErrorUnwrapSingle {
	return &testDetailedErrorUnwrapSingle{
		TestDetailedError: NewTestDetailedError(errorMessage, name, httpStatus, details),
		unwrapSingle:      unwrapSingle,
	}
}

// Unwrap implements the [TestDetailedErrorUnwrapSingle] interface.
func (e *testDetailedErrorUnwrapSingle) Unwrap() error {
	return e.unwrapSingle
}

// TestDetailedErrorUnwrapMulti is a mock blueprint.
type TestDetailedErrorUnwrapMulti interface {
	error
	errorz.ErrorName
	errorz.ErrorHTTPStatus
	errorz.ErrorDetails
	errorz.UnwrapMulti
}

var (
	_ TestDetailedErrorUnwrapMulti = (*testDetailedErrorUnwrapMulti)(nil)
)

type testDetailedErrorUnwrapMulti struct {
	TestDetailedError
	unwrapMulti []error
}

// NewTestDetailedErrorUnwrapMulti initializes a new [TestDetailedErrorUnwrapMulti].
func NewTestDetailedErrorUnwrapMulti(
	errorMessage string,
	name string,
	httpStatus int,
	details map[string]any,
	unwrapMulti []error,
) TestDetailedErrorUnwrapMulti {
	return &testDetailedErrorUnwrapMulti{
		TestDetailedError: NewTestDetailedError(errorMessage, name, httpStatus, details),
		unwrapMulti:       unwrapMulti,
	}
}

// Unwrap implements the [TestDetailedErrorUnwrapMulti] interface.
func (e *testDetailedErrorUnwrapMulti) Unwrap() []error {
	return e.unwrapMulti
}

// TestStringError is a test error.
type TestStringError string

// Error implements the [error] interface.
func (s TestStringError) Error() string {
	return string(s)
}
