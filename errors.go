package fn

// Error returned when an operator requires a non empty array
type EmptyArrayError struct{}

// EmptyArrayError conformance to error
func (e *EmptyArrayError) Error() string {
	return "array cannot be empty"
}
