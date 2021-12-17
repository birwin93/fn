package fn

func Filter[T any](arr []T, f func(T) bool) []T {
	newList := make([]T, 0)
	for _, val := range arr {
		if f(val) {
			newList = append(newList, val)
		}
	}
	return newList
}
