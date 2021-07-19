package complex

import (
	"testing"

	"github.com/scrouthtv/algobra/pkg/number/def"
)

func TestComplexMaths(t *testing.T) {
	z1 := New(def.NewReal(1), def.NewReal(0))
	z2 := New(def.NewReal(0), def.NewReal(-2))
	z3 := New(def.NewReal(1), def.NewReal(-2))

	if !z1.Add(z2).Equal(z3) {
		t.Error("Add failed")
	}

}
