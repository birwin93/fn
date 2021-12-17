package fn

func Contains[T comparable](arr []T, lookup T) bool {
	for _, val := range arr {
		if val == lookup {
			return true
		}
	}
	return false
}
