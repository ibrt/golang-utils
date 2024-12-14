package errorz

import (
	"net/http"
	"reflect"
)

// ErrorName can be implemented by errors to return a name different from their Go type.
type ErrorName interface {
	GetErrorName() string
}

// ErrorHTTPStatus can be implemented by errors to attach an HTTP status to themselves.
type ErrorHTTPStatus interface {
	GetErrorHTTPStatus() int
}

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

// GetName attempts to get a meaningful, stable name for the given error, defaulting to "error".
func GetName(err error) string {
	if err == nil {
		return "<nil>"
	}

	for _, e := range Flatten(err) {
		if n, ok := e.(ErrorName); ok {
			return n.GetErrorName()
		}

		switch n := reflect.TypeOf(e).String(); n {
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

// GetHTTPStatus attempts to get a meaningful, stable HTTP status for the given error, defaulting to 500.
func GetHTTPStatus(err error) int {
	f := Flatten(err)

	for i := len(f) - 1; i >= 0; i-- {
		if h, ok := f[i].(ErrorHTTPStatus); ok {
			return h.GetErrorHTTPStatus()
		}
	}

	return http.StatusInternalServerError
}
