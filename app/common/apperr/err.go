package apperr

type errType struct {
	errType string
}

func (et *errType) Error() string {
	return et.errType
}

func (et *errType) New(msg string) *Error {
	return &Error{
		errType: et.errType,
		message: msg,
	}
}

func (et *errType) Type() *errType {
	return et
}

func NewType(t string) *errType {
	return &errType{
		errType: t,
	}
}

type Error struct {
	errType string
	message string
}

func (e *Error) Error() string {
	return e.errType
}

func (e *Error) Type() string {
	return e.errType
}
