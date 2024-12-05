package memz

import (
	"cmp"
	"sort"
)

// MergeMaps returns a new map built by setting all key/value pairs from the given maps in order.
func MergeMaps[K comparable, V any](mm ...map[K]V) map[K]V {
	length := 0

	for _, m := range mm {
		length += len(m)
	}

	out := make(map[K]V, length)

	for _, m := range mm {
		for k, v := range m {
			out[k] = v
		}
	}

	return out
}

// ShallowCopyMap makes a shallow copy of a map.
func ShallowCopyMap[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	out := make(map[K]V, len(m))

	for k, v := range m {
		out[k] = v
	}

	return out
}

// FilterMap copies a map including only (k, v) pairs for which the predicate returns true.
func FilterMap[K comparable, V any](m map[K]V, f func(k K, v V) bool) map[K]V {
	if m == nil {
		return nil
	}

	out := make(map[K]V, len(m))

	for k, v := range m {
		if f(k, v) {
			out[k] = v
		}
	}

	return out
}

// TransformMapValues returns a new map built by passing all values through the given function, while the keys remain stable.
func TransformMapValues[K comparable, V1 any, V2 any](m map[K]V1, f func(k K, v V1) V2) map[K]V2 {
	if m == nil {
		return nil
	}

	out := make(map[K]V2, len(m))

	for k, v := range m {
		out[k] = f(k, v)
	}

	return out
}

// GetSortedMapKeys returns a slice built by appending all keys in the map and sorting them.
func GetSortedMapKeys[K cmp.Ordered, V any](m map[K]V, less func(K, K) bool) []K {
	ks := make([]K, 0)

	for k := range m {
		ks = append(ks, k)
	}

	sort.Slice(ks, func(i, j int) bool {
		return less(ks[i], ks[j])
	})

	return ks
}
