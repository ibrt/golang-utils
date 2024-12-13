package errorz

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"
)

// UnwrapSingle describes a method which returns a single error.
type UnwrapSingle interface {
	Unwrap() error
}

// UnwrapMulti describes a method which returns multiple errors.
type UnwrapMulti interface {
	Unwrap() []error
}

// IsHelper describes a method called by [errors.Is] to allow customizing its logic.
type IsHelper interface {
	Is(error) bool
}

// AsHelper describes a method called by [errors.As] to allow customizing its logic.
type AsHelper interface {
	As(any) bool
}

var (
	_ error       = (*valueError)(nil)
	_ UnwrapMulti = (*valueError)(nil)
)

type valueError struct {
	Value any
}

// Error implements the error interface.
func (e *valueError) Error() string {
	return fmt.Sprintf("%v", e.Value)
}

// Unwrap implements the [UnwrapMulti] interface.
func (e *valueError) Unwrap() []error {
	if e == nil {
		return nil
	}

	switch v := e.Value.(type) {
	case UnwrapMulti:
		return v.Unwrap()
	case UnwrapSingle:
		if err := v.Unwrap(); err != nil {
			return []error{err}
		}
		return nil
	default:
		return nil
	}
}

var (
	_ error       = (*wrappedError)(nil)
	_ IsHelper    = (*wrappedError)(nil)
	_ AsHelper    = (*wrappedError)(nil)
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

// Is provides interoperability with [errors.Is].
func (e *wrappedError) Is(target error) bool {
	if e == nil || target == nil {
		return e == target
	}

	e.m.Lock()
	defer e.m.Unlock()

	return errors.Is(e.errs[0], target)
}

// As provides interoperability with [errors.As].
func (e *wrappedError) As(target any) bool {
	if e == nil {
		return false
	}

	e.m.Lock()
	defer e.m.Unlock()

	return errors.As(e.errs[0], target)
}

// Unwrap implements the [UnwrapMulti] interface.
func (e *wrappedError) Unwrap() []error {
	if e == nil {
		return nil
	}

	e.m.Lock()
	defer e.m.Unlock()

	errs := slices.Clone(e.errs)
	return errs[1:]
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
