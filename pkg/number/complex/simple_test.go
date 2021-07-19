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
}

func testPair(t *testing.T, id string, a, b, r, theta float64) {
	t.Helper()

	z := New(def.NewReal(a), def.NewReal(b))
	zfrompolar := NewFromPolar(def.NewReal(r), def.NewReal(theta))
	rfromz, thetafromz := z.Polar()

	diff := zfrompolar.Subtract(z)
	if diff.Abs().AsFloat() > 1e-6 {
		t.Errorf("%s Created a wrong complex from polar coordinates:", id)
		t.Logf("%s, should be %s", zfrompolar, z)
		t.Log("Differs by", diff.Abs())
	}

	diffr := rfromz.Subtract(def.NewReal(r))
	difft := thetafromz.Subtract(def.NewReal(theta))
	if diffr.Abs().AsFloat() > 1e-6 || difft.Abs().AsFloat() > 0.1648 { // 0.1648 = 12.25°
		t.Errorf("%s Created wrong polar coordinates from complex number:", id)
		t.Logf("%s / %s, should be %f / %f", rfromz, thetafromz, r, theta)
		t.Logf("Differs by %s/%s", diffr, difft)
	}
}

func TestComplexPolar(t *testing.T) {
	// first quadrant
	testPair(t, "1.", 1, 1, math.Sqrt(2), math.Pi/4.0)

	// second quadrant
	// r = sqrt(13) ~ 3.60555
	// theta = 150° = 5/6 pi
	testPair(t, "2.", -4, 4, 4*math.Sqrt(2), 3.0/4.0*math.Pi)

	// third quadrant
	testPair(t, "3.", -3.122499, -1.8027756, math.Sqrt(13), -5.0/6.0*math.Pi)

	// fourth quadrant
	testPair(t, "4.", 0.5, -0.5, math.Sqrt(0.5), -1.0/4.0*math.Pi)
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
