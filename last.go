package fn

func Last[T any](arr []T) (*T, error) {
	if len(arr) == 0 {
		return nil, &EmptyArrayError{}
	}

	val := arr[len(arr)-1]
	return &val, nil
}
