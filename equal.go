package fn

// Checks if both slices have the same elements in the same order
func Equal[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
