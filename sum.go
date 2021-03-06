package fn

// Adds up all the contents of array
func Sum[T Num](arr []T) T {
	return Reduce(arr, 0, func(total T, val T) T {
		return total + val
	})
}
