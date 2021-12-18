package fn

// Removes duplicates from array
//   arr := []int{1, 2, 3, 3}
//   arr = fn.Dedupe(arr) // outputs [1, 2, 3]
func Dedupe[T comparable](arr []T) []T {
	exists := map[T]bool{}
	return Filter(arr, func(item T) bool {
		if exists[item] {
			return false
		} else {
			exists[item] = true
			return true
		}
	})
}
