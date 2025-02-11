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

// TransformMapValues returns a new map built by passing all values through the given function, while the keys remain
// stable.
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
	if m == nil {
		return nil
	}

	out := make([]K, len(m))
	i := 0

	for k := range m {
		out[i] = k
		i++
	}

	sort.Slice(out, func(i, j int) bool {
		return less(out[i], out[j])
	})

	return out
}

// GetMapValuesSortedByKey returns a slice built by appending all values in the map, sorted by their key.
func GetMapValuesSortedByKey[K cmp.Ordered, V any](m map[K]V, less func(K, K) bool) []V {
	if m == nil {
		return nil
	}

	out := make([]V, len(m))

	for i, k := range GetSortedMapKeys(m, less) {
		out[i] = m[k]
	}

	return out
}

// GetSortedMapValues returns a slice built by appending all values in the map and sorting them.
func GetSortedMapValues[K, V cmp.Ordered](m map[K]V, less func(V, V) bool) []V {
	if m == nil {
		return nil
	}

	out := make([]V, len(m))
	i := 0

	for _, v := range m {
		out[i] = v
		i++
	}

	sort.Slice(out, func(i, j int) bool {
		return less(out[i], out[j])
	})

	return out
}

// MapEntry describes a (key, value) pair.
type MapEntry[K, V any] struct {
	Key   K
	Value V
}

// GetMapEntriesSortedByKey returns a slice built by appending all (key, value) in the map, sorted by key.
func GetMapEntriesSortedByKey[K cmp.Ordered, V any](m map[K]V, less func(K, K) bool) []*MapEntry[K, V] {
	if m == nil {
		return nil
	}

	out := make([]*MapEntry[K, V], len(m))
	i := 0

	for k, v := range m {
		out[i] = &MapEntry[K, V]{
			Key:   k,
			Value: v,
		}
		i++
	}

	sort.Slice(out, func(i, j int) bool {
		return less(out[i].Key, out[j].Key)
	})

	return out
}

// GetMapEntriesSortedByValue returns a slice built by appending all (key, value) in the map, sorted by value.
func GetMapEntriesSortedByValue[K, V cmp.Ordered](m map[K]V, less func(V, V) bool) []*MapEntry[K, V] {
	if m == nil {
		return nil
	}

	out := make([]*MapEntry[K, V], len(m))
	i := 0

	for k, v := range m {
		out[i] = &MapEntry[K, V]{
			Key:   k,
			Value: v,
		}
		i++
	}

	sort.Slice(out, func(i, j int) bool {
		return less(out[i].Value, out[j].Value)
	})

	return out
}
