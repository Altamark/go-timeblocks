package timeblocks

type ErrorNoLocation struct{}

func (e *ErrorNoLocation) Error() string {
	return "Error: location is empty"
}
