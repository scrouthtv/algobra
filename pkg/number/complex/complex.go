package complex

import "fmt"

import "github.com/scrouthtv/algobra/pkg/number"
import "github.com/scrouthtv/algobra/pkg/number/def"

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
		return New(c.realpart.Add(v), c.imgpart)
	case *number.Complex:
		return New(c.realpart.Add(v.realpart),
			c.imgpart.Add(v.imgpart))
	default:
		return n.Add(c)
	}
}

func (c *Complex) Abs() Real {
	inner := c.realpart.Pow(2).Add(c.imgpart.Pow(2))
	return def.NewReal(inner.Nthrt(2))
}

func (c *Complex) String() string {
	return fmt.Sprintf("%s + %si", c.realpart.String(), c.imgpart.String())
}
