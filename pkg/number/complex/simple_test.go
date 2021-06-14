package complex

import "testing"

import "github.com/scrouthtv/algobra/pkg/number"
import "github.com/scrouthtv/algobra/pkg/number/fastreal"

func TestComplexMaths(t *testing.T) {
	var z1 number.Number = New(fastreal.New(1), fastreal.New(0))
	var z2 number.Number = New(fastreal.New(0), fastreal.New(-2))
}
