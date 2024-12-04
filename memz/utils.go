package memz

import (
	"cmp"
	"fmt"
	"reflect"
)

// Min returns the lowest value.
func Min[T cmp.Ordered](v1 T, vs ...T) T {
	for _, v := range vs {
		if cmp.Less(v, v1) {
			v1 = v
		}
	}

	return v1
}

// Max returns the highest value.
func Max[T cmp.Ordered](v1 T, vs ...T) T {
	for _, v := range vs {
		if cmp.Less(v1, v) {
			v1 = v
		}
	}

	return v1
}

// IsAnyNil returns true if the given value is a "nil" of any type.
// Note that this method uses reflection.
func IsAnyNil(x any) bool {
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

// PredicateIsZeroValue returns true if v is the zero-value of its type.
func PredicateIsZeroValue[T comparable](v T) bool {
	var z T
	return v == z
}

// TransformSprintf stringifies values using fmt.Sprintf("%v").
func TransformSprintf[V any](v V) string {
	return fmt.Sprintf("%v", v)
}
