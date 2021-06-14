package number

import "fmt"

type ErrDivideByZero struct {
}

func (e *ErrDivideByZero) Error() string {
	return "can't divide by 0"
}

type ErrUnsupportedRoot struct {
	n uint
}

func (e *ErrUnsupportedRoot) Error() string {
	return fmt.Sprintf("cant' calculate the %d-th root", e.n)
}
