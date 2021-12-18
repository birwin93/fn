package fn

// Maps items to an array of another type and then flattens into one array
func FlatMap[T any, N any](arr []T, f func(T) []N) []N {
	newList := make([]N, 0)
	for _, val := range arr {
		newVals := f(val)
		newList = append(newList, newVals...)
	}
	return newList
}
