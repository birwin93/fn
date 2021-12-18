package fn

// Gets the value for the key in the map, or returns provided fallback
func GetOr[K comparable, V any](dict map[K]V, key K, fallback V) V {
	if val, ok := dict[key]; ok {
		return val
	} else {
		return fallback
	}
}
