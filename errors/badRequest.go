package errors

// badRequestError is thrown when a bad request is sent.
// Use the constructor, NewBadRequestError() to create this error.
type badRequestError struct {
	baseError
}

// NewBadRequestError is a constructor for a new badRequestError that sets it's default fields.
func NewBadRequestError() badRequestError {
	err := badRequestError{
		baseError{
			Message: BadRequestErrMsg,
			Code:    BadRequestErrCode,
		}}
	return err
}

func (e badRequestError) Error() string {
	return e.Message
}
