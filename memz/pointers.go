package memz

// Ptr returns a pointer to the given value.
func Ptr[T any](v T) *T {
	return &v
}

// PtrIfTrue returns a pointer to the given value of cond is true, nil otherwise.
func PtrIfTrue[T any](cond bool, v T) *T {
	if cond {
		return &v
	}
	return nil
}

// PtrZeroToNil returns a pointer to the given value if different from the zero-value, nil otherwise.
func PtrZeroToNil[T comparable](v T) *T {
	var z T
	if v == z {
		return nil
	}

	return &v
}

// PtrZeroToNilIfTrue returns a pointer to the given value if different from the zero-value and cond is true, nil otherwise.
func PtrZeroToNilIfTrue[T comparable](cond bool, v T) *T {
	var z T
	if !cond || v == z {
		return nil
	}

	return &v
}

// ValNilToZero returns the value of the given pointer if found, a zero-value of the same type otherwise.
func ValNilToZero[T any](v *T) T {
	if v == nil {
		var z T
		return z
	}

	return *v
}

// ValNilToDef returns the value of the given pointer if found, the given default value otherwise.
func ValNilToDef[T any](v *T, d T) T {
	if v == nil {
		return d
	}

	return *v
}

// SlicePtr returns a copy of the given slice with each element passed through Ptr.
func SlicePtr[T any](s []T) []*T {
	if s == nil {
		return nil
	}

	out := make([]*T, len(s), cap(s))

	for i, v := range s {
		out[i] = Ptr(v)
	}

	return out
}

// SlicePtrZeroToNil returns a copy of the given slice with each element passed through PtrZeroToNil.
func SlicePtrZeroToNil[T comparable](s []T) []*T {
	if s == nil {
		return nil
	}

	out := make([]*T, len(s), cap(s))

	for i, v := range s {
		out[i] = PtrZeroToNil(v)
	}

	return out
}

// SliceValNilToZero returns a copy of the given slice with each element passed through ValNilToZero.
func SliceValNilToZero[T any](s []*T) []T {
	if s == nil {
		return nil
	}

	out := make([]T, len(s), cap(s))

	for i, v := range s {
		out[i] = ValNilToZero(v)
	}

	return out
}

// SliceValNilToDef returns a copy of the given slice with each element passed through ValNilToDef.
func SliceValNilToDef[T any](s []*T, d T) []T {
	if s == nil {
		return nil
	}

	out := make([]T, len(s), cap(s))

	for i, v := range s {
		out[i] = ValNilToDef(v, d)
	}

	return out
}

// MapPtr returns a copy of the given map with each value passed through Ptr.
func MapPtr[K comparable, V any](m map[K]V) map[K]*V {
	if m == nil {
		return nil
	}

	out := make(map[K]*V, len(m))

	for k, v := range m {
		out[k] = Ptr(v)
	}

	return out
}

// MapPtrZeroToNil returns a copy of the given map with each value passed through PtrZeroToNil.
func MapPtrZeroToNil[K comparable, V comparable](m map[K]V) map[K]*V {
	if m == nil {
		return nil
	}

	out := make(map[K]*V, len(m))

	for k, v := range m {
		out[k] = PtrZeroToNil(v)
	}

	return out
}

// MapValNilToZero returns a copy of the given map with each value passed through ValNilToZero.
func MapValNilToZero[K comparable, V any](m map[K]*V) map[K]V {
	if m == nil {
		return nil
	}

	out := make(map[K]V, len(m))

	for k, v := range m {
		out[k] = ValNilToZero(v)
	}

	return out
}

// MapValNilToDef returns a copy of the given map with each value passed through ValNilToDef.
func MapValNilToDef[K comparable, V any](m map[K]*V, d V) map[K]V {
	if m == nil {
		return nil
	}

	out := make(map[K]V, len(m))

	for k, v := range m {
		out[k] = ValNilToDef(v, d)
	}

	return out
}
