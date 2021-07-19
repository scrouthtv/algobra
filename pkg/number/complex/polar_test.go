package complex

import (
	"math"
	"testing"

	"github.com/scrouthtv/algobra/pkg/number/def"
)

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
