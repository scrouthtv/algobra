package number

type Number interface {
	Add(Number) Number
	Subtract(Number) Number
	Multiply(Number) Number
	Divide(Number) (Number, error)

	Abs() Real
	Equal(Number) bool
}

// Real represents a real number.
// Real numbers can be returned as an (approximated) float number
type Real interface {
	Number
	AsFloat() float64
}
