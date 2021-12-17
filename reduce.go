package fn

func Reduce[T any, N any](arr []T, initial N, f func(N, T) N) N {
	reduced := initial
	for _, val := range arr {
		reduced = f(reduced, val)
	}
	return reduced
}