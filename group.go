package fn

func Group[T any, N comparable](arr []T, f func(T) N) map[N][]T {
	groups := make(map[N][]T)
	for _, val := range arr {
		key := f(val)
		members, ok := groups[key]
		if !ok {
			members = make([]T, 0)
		}
		members = append(members, val)
		groups[key] = members
	}
	return groups
}