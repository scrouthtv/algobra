package complex

import (
	"fmt"
	"math"

	"github.com/scrouthtv/algobra/pkg/number"
	"github.com/scrouthtv/algobra/pkg/number/def"
)

type Complex struct {
	realpart number.Real
	imgpart  number.Real
}

func New(realpart, imgpart number.Real) *Complex {
	return &Complex{realpart, imgpart}
}

func NewFromPolar(r, theta number.Real) *Complex {
	a := def.NewReal(math.Cos(theta.AsFloat())).MultiplyReal(r)
	b := def.NewReal(math.Sin(theta.AsFloat())).MultiplyReal(r)

	return New(a, b)
}

func (c *Complex) Add(n number.Number) number.Number {
	switch v := n.(type) {
	case number.Real:
		return New(c.realpart.AddReal(v), c.imgpart)
	case number.Complex:
		return New(c.realpart.AddReal(v.RealPart()),
			c.imgpart.AddReal(v.ImgPart()))
	default:
		return n.Add(c)
	}
}

func (c *Complex) Subtract(n number.Number) number.Number {
	switch v := n.(type) {
	case number.Real:
		return New(c.realpart.SubtractReal(v), c.imgpart)
	case number.Complex:
		return New(c.realpart.SubtractReal(v.RealPart()),
			c.imgpart.SubtractReal(v.ImgPart()))
	default:
		return n.Subtract(c)
	}
}

func (c *Complex) Multiply(n number.Number) number.Number {
	switch v := n.(type) {
	case number.Real:
		return New(c.realpart.MultiplyReal(v), c.imgpart.MultiplyReal(v))
	case number.Complex:
		//   - c -      - v -
		// (a + bi) * (c + di) = (ac - bd) + (ad + bc)i
		r := c.realpart.MultiplyReal(v.RealPart())              // ac
		r = r.SubtractReal(c.imgpart.MultiplyReal(v.ImgPart())) // -bd
		i := c.realpart.MultiplyReal(v.ImgPart())               // ad
		i = i.AddReal(c.imgpart.MultiplyReal(v.RealPart()))     // +bc

		return New(r, i)
	default:
		return n.Multiply(c)
	}
}

func (c *Complex) Divide(n number.Number) (number.Number, error) {
	if n.Equal(def.Zero) {
		return nil, &number.ErrDivideByZero{}
	}

	switch v := n.(type) {
	case number.Real:
		r, _ := c.realpart.DivideReal(v)
		i, _ := c.imgpart.DivideReal(v)

		return New(r, i), nil
	case number.Complex:
		// a + bi   a + bi   c - di   ac - adi + cbi + bd   ac + bd   bc - ad
		// ------ = ------ x ------ = ------------------- = ------- + ------- i
		// c + di   c + di   c - di   cc - cdi + cdi + dd   cc + dd   cc + dd

		// a = c.RealPart()
		// b = c.ImgPart()
		// c = v.RealPart()
		// d = v.ImgPart()

		den := v.RealPart().MultiplyReal(v.RealPart())
		den = den.AddReal(v.ImgPart().MultiplyReal(v.ImgPart()))

		r := c.realpart.MultiplyReal(v.RealPart())
		r = r.AddReal(c.imgpart.MultiplyReal(v.ImgPart()))

		i := c.imgpart.MultiplyReal(v.RealPart())
		i = i.SubtractReal(c.realpart.MultiplyReal(v.ImgPart()))

		// den != 0 because cc + dd != 0 because cc >= 0 and dd > 0
		r, _ = r.DivideReal(den)
		i, _ = i.DivideReal(den)

		return New(r, i), nil
	default:
		return n.Divide(c)
	}
}

func (c *Complex) Pow(n number.Real) number.Number {
	r, theta := c.Polar()

	r, _ = r.Pow(n).(number.Real)

	theta = theta.MultiplyReal(n)

	return NewFromPolar(r, theta)
}

func (c *Complex) Sqrt() number.Number {
	// https://math.stackexchange.com/a/44500/622559

	r := c.Abs()

	if r.Equal(def.Zero) {
		return def.Zero
	}

	v := c.Add(r)

	ans := r.Sqrt()
	ans, _ = ans.Multiply(v).Divide(v.Abs())

	return ans
}

func (c *Complex) Abs() number.Real {
	inner := c.realpart.Pow(def.NewReal(2)).Add(c.imgpart.Pow(def.NewReal(2)))
	return inner.Sqrt().(number.Real)
}

func (c *Complex) Equal(other number.Number) bool {
	switch v := other.(type) {
	case number.Real:
		return c.realpart.Equal(v) && c.imgpart.Equal(def.Zero)
	case number.Complex:
		return c.realpart.Equal(v.RealPart()) && c.imgpart.Equal(v.ImgPart())
	default:
		return other.Equal(c)
	}
}

func (c *Complex) String() string {
	return fmt.Sprintf("%s + %si", c.realpart.String(), c.imgpart.String())
}

func (c *Complex) RealPart() number.Real {
	return c.realpart
}

func (c *Complex) ImgPart() number.Real {
	return c.imgpart
}

func (c *Complex) Polar() (number.Real, number.Real) {
	r := c.Abs()
	if r.Equal(def.Zero) {
		return def.Zero, def.Zero
	}

	realn, _ := c.realpart.DivideReal(r)
	theta := math.Acos(realn.AsFloat())

	if c.imgpart.AsFloat() < 0 {
		theta = -theta
	}

	return r, def.NewReal(theta)
}
