package errors

type BussinessError struct {
	Msg string
}

func (e *BussinessError) Error() string {
	return e.Msg
}

func New(msg string) *BussinessError {
	return &BussinessError{
		Msg: msg,
	}
}
