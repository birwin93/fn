package fn

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
