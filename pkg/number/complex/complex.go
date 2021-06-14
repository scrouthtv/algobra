package complex

import "fmt"

import "github.com/scrouthtv/algobra/pkg/number"

type Complex struct {
	realpart number.Real
	imgpart number.Real
}

func New(realpart, imgpart number.Real) *Complex {
	return &Complex{realpart, imgpart}
}

func (c *Complex) Add(n number.Number) (number.Number) {
	switch (n.(type)) {
	case number.Real:
		return New(c.realpart.Add(n.(number.Real)), c.imgpart)
	case number.Complex:
		return New(c.realpart.Add(n.(*number.Complex).realpart),
			c.imgpart.Add(n.(*number.Complex).imgpart))
	}
}

func (c *Complex) String() string {
	return fmt.Sprintf("%s + %si", c.realpart.String(), c.imgpart.String())
}
