package fn

// Grabs last item in array
// Will return empty value and false if array is empty
func Last[T any](arr []T) (T, bool) {
	if len(arr) == 0 {
		var zero T
		return zero, false
	}

	val := arr[len(arr)-1]
	return val, true
}
