package fastreal

import "github.com/scrouthtv/algobra/pkg/number"

// Fastreal implements a real number using Go's float64.
type Fastreal float64

func New(f float64) Fastreal {
	return Fastreal(f)
}

func (f Fastreal) Add(n number.Number) (number.Number, error) {
	f2, ok := n.(Fastreal)
	if ok {
		return f + f2, nil
	}

	return n.Add(f2)
}

func (f Fastreal) Subtract(n number.Number) (number.Number, error) {
	f2, ok := n.(Fastreal)
	if ok {
		return f - f2, nil
	}

	return n.Subtract(f2)
}

func (f Fastreal) Multiply(n number.Number) (number.Number, error) {
	f2, ok := n.(Fastreal)
	if ok {
		return f * f2, nil
	}

	return n.Multiply(f2)
}

func (f Fastreal) Divide(n number.Number) (number.Number, error) {
	f2, ok := n.(Fastreal)
	if f2 == 0 {
		return nil, number.ErrDivideByZero{}
	}
	if ok {
		return f / f2, nil
	}

	return n.Multiply(f2)
}

func (f Fastreal) Abs() (number.Realnumber, error) {
	if f < 0 {
		return New(-float64(f)), nil
	} else {
		return f, nil
	}
}
