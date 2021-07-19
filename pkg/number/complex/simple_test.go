package complex

import (
	"math"
	"testing"

	"github.com/scrouthtv/algobra/pkg/number"
	"github.com/scrouthtv/algobra/pkg/number/def"
)

func TestComplexMaths(t *testing.T) {
	var z1 number.Complex = New(def.NewReal(1), def.NewReal(3))
	z2 := New(def.NewReal(0), def.NewReal(-2))
	/* z1 + z2 = */ z3 := New(def.NewReal(1), def.NewReal(1))
	/* z1 - z2 = */ z4 := New(def.NewReal(1), def.NewReal(5))
	/* z1 * z2 = */ z5 := New(def.NewReal(6), def.NewReal(-2))
	/* z1 / z2 = */ z6 := New(def.NewReal(-1.5), def.NewReal(.5))

	if !z1.Add(z2).Equal(z3) {
		t.Errorf("Add failed: %s + %s = %s, should be %s", z1, z2, z1.Add(z2), z3)
	}

	if !z1.Subtract(z2).Equal(z4) {
		t.Errorf("Subtract failed: %s - %s = %s, should be %s", z1, z2, z1.Subtract(z2), z4)
	}

	if !z1.Multiply(z2).Equal(z5) {
		t.Errorf("Multiply failed: %s * %s = %s, should be %s", z1, z2, z1.Multiply(z2), z5)
	}

	ans, err := z1.Divide(z2)
	if err != nil {
		t.Errorf("Divide failed: %s / %s failed with error %ss, should be %s", z1, z2, err.Error(), z6)
	}
	if !ans.Equal(z6) {
		t.Errorf("Divide failed: %s / %s = %s, should be %s", z1, z2, ans, z6)
	}

	z7 := New(def.One, def.Zero)

	if !z7.RealPart().Equal(def.One) {
		t.Errorf("RealPart failed: real(%s) = %s, should be %d", z1, z1.RealPart(), 1)
	}

	if !z7.ImgPart().Equal(def.Zero) {
		t.Errorf("ImgPart failed: img(%s) = %s, should be %d", z1, z1.ImgPart(), 0)
	}
}

func TestComplexPow(t *testing.T) {
	z1 := New(def.NewReal(3), def.NewReal(-2))
	n1 := def.NewReal(6)
	/* z1 ^ n1 = */ z2 := New(def.NewReal(-2035), def.NewReal(828))

	diff := z1.Pow(n1).Subtract(z2)
	if diff.Abs().AsFloat() > 1e-6 {
		t.Errorf("Pow failed: %s ^ %s = %s, should be %s", z1, n1, z1.Pow(n1), z2)
		t.Log("Difference:", diff)
	}
}

func TestComplexSqrt(t *testing.T) {
	z1 := New(def.NewReal(-5), def.NewReal(12))
	/* sqrt(z1) = */ z2 := New(def.NewReal(2), def.NewReal(3))

	ans := z1.Sqrt()
	diff := ans.Subtract(z2)
	if diff.Abs().AsFloat() > 1e-6 {
		t.Errorf("Sqrt failed: sqrt(%s) = %s, should be %s", z1, ans, z2)
		t.Log("Difference:", diff.Abs())
	}
}

func TestAbs(t *testing.T) {
	z1 := New(def.NewReal(-5), def.NewReal(12))
	z2 := New(def.NewReal(5), def.NewReal(0))
	z3 := New(def.NewReal(0), def.NewReal(-3))

	if !z1.Abs().Equal(def.NewReal(math.Sqrt(25 + 144))) {
		t.Errorf("Abs 1 failed: %s, should be %s", z1.Abs(), def.NewReal(math.Sqrt(25+144)))
	}

	if !z2.Abs().Equal(def.NewReal(5)) {
		t.Errorf("Abs 2 failed: %s, should be %d", z2.Abs(), 5)
	}

	if !z3.Abs().Equal(def.NewReal(3)) {
		t.Errorf("Abs 3 failed: %s, should be %d", z3.Abs(), 3)
	}
}

func TestEqual(t *testing.T) {
	z1 := New(def.NewReal(1), def.NewReal(2))
	z2 := New(def.NewReal(3), def.NewReal(-1))
	z3 := New(def.NewReal(1), def.NewReal(-2))
	z4 := New(def.NewReal(-1), def.NewReal(2))

	z5 := New(def.NewReal(1), def.NewReal(0))
	z6 := def.NewReal(1)

	if !z1.Equal(z1) {
		t.Errorf("Equal failed: %s != %s, should be true", z1, z1)
	}

	if z1.Equal(z2) {
		t.Errorf("Equal failed: %s == %s, should be false", z1, z2)
	}

	if z1.Equal(z3) {
		t.Errorf("Equal failed: %s == %s, should be false", z1, z3)
	}

	if z1.Equal(z4) {
		t.Errorf("Equal failed: %s == %s, should be false", z1, z4)
	}

	if z1.Equal(z5) {
		t.Errorf("Equal failed: %s == %s, should be false", z1, z5)
	}

	if !z5.Equal(z6) {
		t.Errorf("Equal failed: %s != %s, should be true", z5, z6)
	}
}

func TestConv(t *testing.T) {
	z := New(def.NewReal(12.517829), def.NewReal(0))

	if z.String() != "12.517829 + 0i" {
		t.Errorf("%s.String() = %s, should be %s", z, z.String(), "12.517829 + 0i")
	}
}
