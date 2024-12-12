package memz

import (
	"cmp"
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

// Ternary returns "vT" if "cond" is true, "vF" otherwise.
func Ternary[V any](cond bool, vT V, vF V) V {
	if cond {
		return vT
	}
	return vF
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
