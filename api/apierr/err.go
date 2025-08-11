package apierr

type Error struct {
	Type    string
	Message string
	Status  int
}

func (e *Error) Error() string {
	return e.Message
}

func Wrap(err error, message string, status int) *Error {
	return &Error{
		Type:    err.Error(),
		Message: message,
		Status:  status,
	}
}
