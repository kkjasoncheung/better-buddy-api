package errors

// InvalidPasswordError is thrown when a user enters an incorrect password.
// Use the constructor, NewInvalidPasswordError() to create this error.
type invalidPasswordError struct {
	baseError
}

// NewInvalidPasswordError is a constructor for a new invalidPasswordError that sets it's default fields.
func NewInvalidPasswordError() invalidPasswordError {
	err := invalidPasswordError{
		baseError{
			Message: "Invalid Password Entered. Please try again.",
			Code:    InvalidPasswordErrCode,
		}}
	return err
}

func (e invalidPasswordError) Error() string {
	return e.Message
}
