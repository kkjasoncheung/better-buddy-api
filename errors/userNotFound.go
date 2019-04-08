package errors

// UserNotFoundError is thrown when a user enters an incorrect password.
// Use the constructor, UserNotFoundError() to create this error.
type UserNotFoundError struct {
	baseError
}

// NewUserNotFoundError is a constructor for a new invalidPasswordError that sets it's default fields.
func NewUserNotFoundError() UserNotFoundError {
	err := UserNotFoundError{
		baseError{
			Message: UserNotFoundErrMsg,
			Code:    UserNotFoundErrCode,
		}}
	return err
}

func (e UserNotFoundError) Error() string {
	return e.Message
}
