package def

import (
	"github.com/scrouthtv/algobra/pkg/number"
	"github.com/scrouthtv/algobra/pkg/number/fastreal"
)

var (
	Zero = NewReal(0)
	One  = NewReal(1)
)

// NewReal creates a new real number that uses the current default
// implementation for real numbers.
func NewReal(f float64) number.Real {
	return fastreal.New(f)
}
