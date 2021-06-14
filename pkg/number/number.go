package number

import "fmt"

// Number is any number.
// Currently, real and complex numbers are implemented.
type Number interface {
	fmt.Stringer

	Add(Number) Number
	Subtract(Number) Number
	Multiply(Number) Number
	Divide(Number) (Number, error)
	Pow(Real) Number

	Abs() Real
	Equal(Number) bool
}

// Real represents a real number.
// Real numbers can be returned as an (approximated) float number
type Real interface {
	Number
	AsFloat() float64
}
