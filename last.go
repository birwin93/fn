package fn

func Last[T any](arr []T) T {
	return arr[len(arr)-1]
}
