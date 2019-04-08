package errors

// invalidPasswordError is thrown when a user enters an incorrect password.
type badRequestError struct {
	baseError
}

// NewBadRequestError is a constructor for a new badRequestError that sets it's default fields.
func NewBadRequestError() badRequestError {
	err := badRequestError{
		baseError{
			Message: "Bad Request.",
			Code:    BadRequestError,
		}}
	return err
}

func (e badRequestError) Error() string {
	return e.Message
}
