package complex

import "github.com/scrouthtv/algobra/pkg/number"

type Complex struct {
	realpart number.Real
	imgpart number.Real
}

func New(realpart, imgpart number.Real) *Complex {
	return &Complex{realpart, imgpart}
}

func (c *Complex) String() string {
	return fmt.Sprintf("%s + %si", c.realpart.String(), c.imgpart.String())
}
