package number

type ErrDivideByZero struct {
}

func (e *ErrDivideByZero) Error() string {
	return "can't divide by 0"
}
