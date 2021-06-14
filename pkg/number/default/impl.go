package number

import "github.com/scrouthtv/algobra/pkg/number"
import "github.com/scrouthtv/algobra/pkg/number/fastreal"

// NewReal creates a new real number that uses the current default
// implementation for real numbers.
func NewReal(f float64) number.Real {
	return fastreal.New(f)
}
