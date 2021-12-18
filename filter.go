package fn

// Removes items from array that does not pass predicate
func Filter[T any](arr []T, predicate func(T) bool) []T {
	newList := make([]T, 0)
	for _, val := range arr {
		if predicate(val) {
			newList = append(newList, val)
		}
	}
	return newList
}
