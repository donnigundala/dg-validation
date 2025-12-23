package dgvalidation

// Error is a custom error type that holds field-specific validation errors.
type Error struct {
	Errors map[string]map[string]string
}

// Error implements the error interface.
func (v *Error) Error() string {
	return "Validation failed"
}
