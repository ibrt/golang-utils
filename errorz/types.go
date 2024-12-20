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
	GetErrorName() string
}

// ErrorHTTPStatus can be implemented by errors to attach an HTTP status to themselves.
type ErrorHTTPStatus interface {
	GetErrorHTTPStatus() int
}

// ErrorDetails can be implemented by errors to export some additional, human-readable information about the error.
type ErrorDetails interface {
	GetErrorDetails() map[string]any
}

// UnwrapSingle describes a method which returns a single error.
type UnwrapSingle interface {
	Unwrap() error
}

// UnwrapMulti describes a method which returns multiple errors.
type UnwrapMulti interface {
	Unwrap() []error
}

var (
	_ error        = (*valueError)(nil)
	_ UnwrapSingle = (*valueError)(nil)
)

type valueError struct {
	value any
}

// Error implements the error interface.
func (e *valueError) Error() string {
	return fmt.Sprintf("%v", e.value)
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
