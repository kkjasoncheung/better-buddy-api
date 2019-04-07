package errors

// invalidPasswordError is thrown when a user enters an incorrect password.
type invalidPasswordError struct {
	baseError
}

// NewInvalidPasswordError is a constructor for a new invalidPasswordError that sets it's default fields.
func NewInvalidPasswordError() invalidPasswordError {
	err := invalidPasswordError{
		baseError{
			Message: "Invalid Password Entered. Please try again.",
			Code:    InvalidPasswordError,
		}}
	return err
}

func (e invalidPasswordError) Error() string {
	return e.Message
}
