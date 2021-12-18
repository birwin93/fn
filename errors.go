package fn

type EmptyArrayError struct{}

func (e *EmptyArrayError) Error() string {
	return "array cannot be empty"
}
