package complex

import "testing"

import "github.com/scrouthtv/algobra/pkg/number"
import "github.com/scrouthtv/algobra/pkg/number/impl"

func TestComplexMaths(t *testing.T) {
	var z1 number.Number = New(impl.NewReal(1), impl.NewReal(0))
	var z2 number.Number = New(impl.NewReal(0), impl.NewReal(-2))
}
