package fastreal

import "github.com/scrouthtv/algobra/pkg/number"

// Fastreal implements a real number using Go's float64.
type Fastreal float64

func New(f float64) Fastreal {
	return Fastreal(f)
}

func (f Fastreal) Add(n number.Number) number.Number {
	f2, ok := n.(Fastreal)
	if ok {
		return f + f2
	}

	return n.Add(f2)
}

func (f Fastreal) Subtract(n number.Number) number.Number {
	f2, ok := n.(Fastreal)
	if ok {
		return f - f2
	}

	return n.Subtract(f2)
}

func (f Fastreal) Multiply(n number.Number) number.Number {
	f2, ok := n.(Fastreal)
	if ok {
		return f * f2
	}

	return n.Multiply(f2)
}

func (f Fastreal) Divide(n number.Number) (number.Number, error) {
	f2, ok := n.(Fastreal)
	if f2 == 0 {
		return nil, &number.ErrDivideByZero{}
	}
	if ok {
		return f / f2, nil
	}

	return n.Divide(f2)
}

func (f Fastreal) Abs() number.Real {
	if f < 0 {
		abs := New(-float64(f))
		return &abs, nil
	}

	return f
}

func (f Fastreal) AsFloat() float64 {
	return float64(f)
}

func (f Fastreal) Equal(n number.Number) bool {
	f2, ok := n.(Fastreal)
	if ok {
		return float64(f) == float64(f2)
	}

	return n.Equal(f)
}
