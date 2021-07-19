package complex

import (
	"testing"
	"testing/quick"

	"github.com/scrouthtv/algobra/pkg/number/def"
)

func TestBasicCross(t *testing.T) {
	a := New(def.NewReal(9), def.NewReal(12))
	n := def.NewReal(3)

	x := New(def.NewReal(12), def.NewReal(12))
	if !n.Add(a).Equal(x) {
		t.Errorf("Add failed %s + %s = %s, should be %s", n, a, n.Add(a), x)
	}

	x = New(def.NewReal(6), def.NewReal(12))
	if !n.Subtract(a).Equal(x) {
		t.Errorf("Subtract failed %s - %s = %s, should be %s", n, a, n.Subtract(a), x)
	}

	x = New(def.NewReal(27), def.NewReal(36))
	if !n.Multiply(a).Equal(x) {
		t.Errorf("Multiply failed: %s * %s = %s, should be %s", n, a, n.Multiply(a), x)
	}

	x = New(def.NewReal(3), def.NewReal(4))
	ans, err := a.Divide(n)
	if err != nil {
		t.Errorf("Divide failed: %s / %s gave error %s, should be %s", a, n, err.Error(), x)
	}
	if !ans.Equal(x) {
		t.Errorf("Divide failed: %s / %s = %s, should be %s", a, n, ans, x)
	}
}

func TestBlindCross(t *testing.T) {
	cfg := &quick.Config{}
	quick.Check(testAllOps, cfg)
}

// testAllOps tests whether
//  - (a+bi) + c == c + (a+bi)
//  - (a+bi) - c == -(c + (a+bi))
//  - (a+bi) * c == c * (a+bi)
func testAllOps(a, b, c float64) bool {
	z := New(def.NewReal(a), def.NewReal(b))
	n := def.NewReal(c)

	ansa := z.Add(n)
	ansb := n.Add(z)
	if !ansa.Equal(ansb) {
		return false
	}

	ansa = z.Subtract(n)
	ansb = n.Subtract(z)
	if !ansa.Equal(ansb) {
		return false
	}

	ansa = z.Multiply(n)
	ansb = n.Multiply(z)
	if !ansa.Equal(ansb) {
		return false
	}

	return true
}
