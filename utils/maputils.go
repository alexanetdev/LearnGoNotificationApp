package utils

func FilterMap[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if fn(v) {
			result[k] = v
		}
	}
	return result
}
