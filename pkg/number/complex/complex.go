package complex

import (
	"fmt"

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
		real := c.realpart.MultiplyReal(v.RealPart())                 // ac
		real = real.SubtractReal(c.imgpart.MultiplyReal(v.ImgPart())) // -bd
		img := c.realpart.MultiplyReal(v.ImgPart())                   // ad
		img = img.AddReal(c.imgpart.MultiplyReal(v.RealPart()))       // +bc

		return New(real, img)
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
		return New(c.realpart.AddReal(v.RealPart()),
			c.imgpart.AddReal(v.ImgPart())), nil
	default:
		return n.Divide(c)
	}
}

func (c *Complex) Pow(n number.Real) number.Number {
	panic("not impl")
}

func (c *Complex) Sqrt() number.Number {
	panic("not impl")
}

func (c *Complex) Nthrt(n uint) (number.Number, error) {
	panic("not impl")
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
	panic("not impl")
}
