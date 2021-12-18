package fn

// Converts all items in array to specific output of func
func Map[T any, N any](arr []T, f func(T) N) []N {
	newList := make([]N, len(arr))
	for i, val := range arr {
		newList[i] = f(val)
	}
	return newList
}
