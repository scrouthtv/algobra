package number

type Number interface {
	Add(Number) (Number, error)
	Subtract(Number) (Number, error)
	Multiply(Number) (Number, error)
	Divide(Number) (Number, error)

	Abs() (RealNumber, error)
}

// Real represents a real number.
// Real numbers can be returned as an (approximated) float number
type Real interface {
	Number
	AsFloat() float64
}
