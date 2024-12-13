package errorz

import (
	"reflect"
)

// MaybeSetMetadata sets the given metadata (k, v) on the error if it has been wrapped, does nothing otherwise.
func MaybeSetMetadata(err error, k, v any) {
	if e, ok := err.(*wrappedError); ok {
		e.setMetadata(k, v)
	}
}

// MustGetMetadata gets the given metadata key from the error, panics if not found or wrong type.
func MustGetMetadata[T any](err error, k any) T {
	return err.(*wrappedError).metadata[k].(T)
}

// MaybeGetMetadata tries to get the given metadata key from the error.
func MaybeGetMetadata[T any](err error, k any) (T, bool) {
	if e, ok := err.(*wrappedError); ok {
		if m, ok := e.getMetadata(k); ok {
			if v, ok := m.(T); ok {
				return v, true
			}
		}
	}

	var v T
	return v, false
}

// GetName attempts to get a meaningful name for the given error.
func GetName(err error) string {
	if err == nil {
		return "<nil>"
	}

	errs := make([]error, 0)
	errs = append(errs, err)

	for i := 0; i < len(errs); i++ {
		switch e := errs[i].(type) {
		case *wrappedError:
			for j := len(e.errs) - 1; j >= 0; j-- {
				errs = append(errs, e.errs[j])
			}
		case UnwrapMulti:
			es := e.Unwrap()
			for j := len(es) - 1; j >= 0; j-- {
				errs = append(errs, es[j])
			}
		case UnwrapSingle:
			if ee := e.Unwrap(); ee != nil {
				errs = append(errs, ee)
			}
		}
	}

	for i := len(errs) - 1; i >= 0; i-- {
		switch n := reflect.TypeOf(errs[i]).String(); n {
		case
			"*errors.errorString",
			"*errors.joinError",
			"*errorz.wrappedError",
			"*fmt.wrapError",
			"*fmt.wrapErrors":
			continue
		default:
			return n
		}
	}

	return "error"
}
