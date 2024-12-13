package errorz

import (
	"errors"
	"reflect"
	"sync"
)

// Wrap wraps the given errors.
func Wrap(err error, outerErrs ...error) error {
	if err == nil {
		MustErrorf("err is nil")
	}

	wErr, ok := err.(*wrappedError)
	if !ok {
		wErr = &wrappedError{
			m:      &sync.Mutex{},
			errs:   []error{err},
			frames: GetFrames(nil),
		}
	}

	wErr.m.Lock()
	defer wErr.m.Unlock()

	for _, outerErr := range outerErrs {
		if outerErr != nil {
			wErr.errs = append(wErr.errs, outerErr)
		}
	}

	return wErr
}

// MaybeWrap is like [Wrap], but returns nil if called with a nil error.
func MaybeWrap(err error, outerErrs ...error) error {
	if err != nil {
		return Wrap(err, outerErrs...)
	}

	return nil
}

// MustWrap is like [Wrap], but panics with the wrapped error instead of returning it.
func MustWrap(innerErr error, outerErrs ...error) {
	panic(Wrap(innerErr, outerErrs...))
}

// MaybeMustWrap is like [MustWrap], but does nothing if called with a nil error.
func MaybeMustWrap(err error, outerErrs ...error) {
	if err != nil {
		MustWrap(err, outerErrs...)
	}
}

// WrapRecover takes a recovered value and converts it to a wrapped error.
func WrapRecover(r any, outerErrs ...error) error {
	if isNil(r) {
		MustErrorf("r is nil")
	}

	err, ok := r.(error)
	if !ok {
		err = &valueError{
			Value: r,
		}
	}

	return Wrap(err, outerErrs...)
}

// MaybeWrapRecover is like [WrapRecover] but returns nil if called with a nil value.
func MaybeWrapRecover(r any, outerErrs ...error) error {
	if !isNil(r) {
		return WrapRecover(r, outerErrs...)
	}

	return nil
}

// As provides a more handy implementation of [errors.As] using generics.
func As[T any](err error) (T, bool) {
	var t T
	return t, errors.As(err, &t)
}

func isNil(x any) bool {
	if x == nil {
		return true
	}

	switch v := reflect.ValueOf(x); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
