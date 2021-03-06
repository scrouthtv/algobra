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
	Sqrt() Number

	Abs() Real
	Equal(Number) bool
}

// Real represents a real number.
// Real numbers can be returned as an (approximated) float number.
type Real interface {
	Number

	AddReal(Real) Real
	SubtractReal(Real) Real
	MultiplyReal(Real) Real
	DivideReal(Real) (Real, error)

	AsFloat() float64
}

type Complex interface {
	Number

	RealPart() Real
	ImgPart() Real

	// Polar returns the polar coordinates (r, theta)
	// of the complex number.
	Polar() (Real, Real)
}
