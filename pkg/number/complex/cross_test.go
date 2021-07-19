package complex

import (
	"testing"

	"github.com/scrouthtv/algobra/pkg/number/def"
)

func TestBasicCross(t *testing.T) {
	a := New(def.NewReal(8), def.NewReal(12))
	n := def.NewReal(3)

	x := New(def.NewReal(24), def.NewReal(36))
	if !n.Multiply(a).Equal(x) {
		t.Errorf("Multiply failed: %s * %s = %s, should be %s", n, a, n.Multiply(a), x)
	}
}
