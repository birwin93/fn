package fn

// Returns shared items between two arrays
func Intersect[T comparable](arr1 []T, arr2 []T) []T {
	inArr1 := map[T]bool{}
	for _, item := range arr1 {
		inArr1[item] = true
	}

	inBoth := []T{}
	for _, item := range arr2 {
		if inArr1[item] {
			inBoth = append(inBoth, item)
		}
	}

	return inBoth
}
