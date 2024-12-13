package errorz

import (
	"fmt"
)

// Errorf creates an error and wraps it.
func Errorf(format string, a ...any) error {
	return Wrap(fmt.Errorf(format, a...))
}

// MustErrorf is like [Errorf] but panics with the wrapped error instead of returning it.
func MustErrorf(format string, a ...any) {
	panic(Errorf(format, a...))
}

// Assertf is like [MustErrorf] if cond is false, does nothing otherwise.
func Assertf(cond bool, format string, a ...any) {
	if !cond {
		MustErrorf(format, a...)
	}
}
