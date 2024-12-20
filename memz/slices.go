package memz

// SafeSliceIndexZero indexes a slice, supports negative indexes (from end), and returns a zero-value instead of panic.
func SafeSliceIndexZero[T any](s []T, i int) T {
	if i < 0 {
		i = len(s) + i
	}

	if i < 0 || i >= len(s) {
		var z T
		return z
	}

	return s[i]
}

// SafeSliceIndexDef indexes a slice, supports negative indexes (from end), and returns a default value instead of
// panic.
func SafeSliceIndexDef[T any](s []T, i int, d T) T {
	if i < 0 {
		i = len(s) + i
	}

	if i < 0 || i >= len(s) {
		return d
	}

	return s[i]
}

// SafeSliceIndexPtr indexes a slice, supports negative indexes (from end), and returns pointer to value instead of
// panic.
func SafeSliceIndexPtr[T any](s []T, i int) *T {
	if i < 0 {
		i = len(s) + i
	}

	if i < 0 || i >= len(s) {
		return nil
	}

	return &s[i]
}

// ConcatSlices returns a new slice built by appending all values from the given slices in order.
func ConcatSlices[T any](ss ...[]T) []T {
	length := 0

	for _, s := range ss {
		length += len(s)
	}

	out := make([]T, 0, length)

	for _, s := range ss {
		out = append(out, s...)
	}

	return out
}

// ShallowCopySlice makes a shallow copy of a slice.
func ShallowCopySlice[T any](s []T) []T {
	if s == nil {
		return nil
	}

	out := make([]T, len(s), cap(s))
	copy(out, s)

	return out
}

// FilterSlice makes a shallow copy of a slice including only elements for which the predicate returns true.
func FilterSlice[T any](s []T, f func(t T) bool) []T {
	if s == nil {
		return nil
	}

	out := make([]T, 0, cap(s))

	for _, t := range s {
		if f(t) {
			out = append(out, t)
		}
	}

	return out
}

// BatchSlice splits a slice in batches.
func BatchSlice[T any](s []T, batchSize int) [][]T {
	out := make([][]T, 0)

	for i := 0; i < len(s); i += batchSize {
		out = append(out, s[i:Min(i+batchSize, len(s))])
	}

	return out
}

// TransformSlice returns a new slice built by passing all elements through the given function.
func TransformSlice[I any, O any](s []I, f func(i int, v I) O) []O {
	if s == nil {
		return nil
	}

	out := make([]O, 0, cap(s))

	for i, v := range s {
		out = append(out, f(i, v))
	}

	return out
}

// SliceToStructMap converts a slice to struct map.
func SliceToStructMap[T comparable](s []T) map[T]struct{} {
	if s == nil {
		return nil
	}

	out := make(map[T]struct{}, len(s))

	for _, v := range s {
		out[v] = struct{}{}
	}

	return out
}
