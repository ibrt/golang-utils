package errorz

import (
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
