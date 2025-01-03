package errorz

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"sync"
)

// ErrorName can be implemented by errors to return a name different from their Go type.
type ErrorName interface {
	// GetErrorName returns a human-readable name describing an error.
	GetErrorName() string
}

// ErrorHTTPStatus can be implemented by errors to attach an HTTP status to themselves.
type ErrorHTTPStatus interface {
	// GetErrorHTTPStatus returns an HTTP status associated with an error.
	GetErrorHTTPStatus() int
}

// ErrorDetails can be implemented by errors to export some additional, human-readable information about the error.
type ErrorDetails interface {
	// GetErrorDetails returns some human-readable information about an error.
	GetErrorDetails() map[string]any
}

// UnwrapSingle describes a method which returns a single error.
type UnwrapSingle interface {
	// Unwrap returns a single wrapped error, if any, nil otherwise.
	Unwrap() error
}

// UnwrapMulti describes a method which returns multiple errors.
type UnwrapMulti interface {
	// Unwrap returns one or more wrapped errors, if any, nil otherwise.
	Unwrap() []error
}

var (
	_ error        = (*valueError)(nil)
	_ ErrorName    = (*valueError)(nil)
	_ ErrorDetails = (*valueError)(nil)
	_ UnwrapSingle = (*valueError)(nil)
)

type valueError struct {
	value any
}

// GetErrorName implements the [ErrorName] interface.
func (*valueError) GetErrorName() string {
	return "value-error"
}

// GetErrorDetails implements the [ErrorDetails] interface.
func (e *valueError) GetErrorDetails() map[string]any {
	return map[string]any{
		"value": e.value,
	}
}

// Unwrap implements the [UnwrapSingle] interface.
func (e *valueError) Unwrap() error {
	if e == nil {
		return nil
	}

	if ee, ok := e.value.(error); ok {
		return ee
	}
	return nil
}

var (
	_ error       = (*wrappedError)(nil)
	_ UnwrapMulti = (*wrappedError)(nil)
)

// Error implements the error interface.
func (e *valueError) Error() string {
	return fmt.Sprintf("%v", e.value)
}

type wrappedError struct {
	m        *sync.Mutex
	errs     []error
	frames   Frames
	metadata map[any]any
}

// Error implements the error interface.
func (e *wrappedError) Error() string {
	e.m.Lock()
	defer e.m.Unlock()

	w := &strings.Builder{}

	for i := len(e.errs) - 1; i >= 0; i-- {
		if i < len(e.errs)-1 {
			_, _ = w.WriteString(": ")
		}
		_, _ = w.WriteString(e.errs[i].Error())
	}

	return w.String()
}

// Unwrap implements the [UnwrapMulti] interface.
func (e *wrappedError) Unwrap() []error {
	if e == nil {
		return nil
	}

	e.m.Lock()
	defer e.m.Unlock()

	errs := slices.Clone(e.errs)
	return errs
}

func (e *wrappedError) setMetadata(k, v any) {
	e.m.Lock()
	defer e.m.Unlock()

	e.metadata[k] = v
}

func (e *wrappedError) getMetadata(k any) (any, bool) {
	e.m.Lock()
	defer e.m.Unlock()

	v, ok := e.metadata[k]
	return v, ok
}

var (
	genericErrorsErrorString     = reflect.TypeOf(fmt.Errorf("e"))
	genericErrorsErrorStringName = genericErrorsErrorString.String()
	genericErrorsJoinError       = reflect.TypeOf(errors.Join(fmt.Errorf("e")))
	genericErrorsJoinErrorName   = genericErrorsJoinError.String()
	genericFmtWrapError          = reflect.TypeOf(fmt.Errorf("%w", fmt.Errorf("e")))
	genericFmtWrapErrorName      = genericFmtWrapError.String()
	genericFmtWrapErrors         = reflect.TypeOf(fmt.Errorf("%w%w", fmt.Errorf("e"), fmt.Errorf("e")))
	genericFmtWrapErrorsName     = genericFmtWrapErrors.String()
)

func isGenericError(err error) bool {
	if err == nil {
		return false
	}

	switch reflect.TypeOf(err) {
	case genericErrorsErrorString, genericErrorsJoinError, genericFmtWrapError, genericFmtWrapErrors:
		return true
	default:
		return false
	}
}

func isJoinError(err error) bool {
	if err == nil {
		return false
	}

	return reflect.TypeOf(err) == genericErrorsJoinError
}

func isWrapError(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*wrappedError) //nolint:errorlint
	return ok
}
