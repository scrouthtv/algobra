package fastreal

import (
	"testing"

	"github.com/scrouthtv/algobra/pkg/number"
)

func TestRealMaths(t *testing.T) {
	n1, n2 := New(15), New(5)

	if !n1.AddReal(n2).Equal(New(20)) {
		t.Errorf("%s + %s = %s, should be %d", n1, n2, n1.AddReal(n2), 20)
	}

	if !n1.SubtractReal(n2).Equal(New(10)) {
		t.Errorf("%s - %s = %s, should be %d", n1, n2, n1.SubtractReal(n2), 10)
	}

	if !n1.MultiplyReal(n2).Equal(New(75)) {
		t.Errorf("%s * %s = %s, should be %d", n1, n2, n1.MultiplyReal(n2), 75)
	}

	ans, err := n1.DivideReal(n2)
	if err != nil {
		t.Errorf("%s / %s gave error: %s", n1, n2, err)
	} else if !ans.Equal(New(3)) {
		t.Errorf("%s / %s = %s, should be %d", n1, n2, ans, 3)
	}

	ans, err = n1.DivideReal(New(0))
	if err == nil {
		t.Errorf("%s / %s = %s, should be error", n1, New(0), ans)
	}
}

func TestBasicMaths(t *testing.T) {
	var n1, n2 number.Number = New(15), New(5)

	if !n1.Add(n2).Equal(New(20)) {
		t.Errorf("%s + %s = %s, should be %d", n1, n2, n1.Add(n2), 20)
	}

	if !n1.Subtract(n2).Equal(New(10)) {
		t.Errorf("%s - %s = %s, should be %d", n1, n2, n1.Subtract(n2), 10)
	}

	if !n1.Multiply(n2).Equal(New(75)) {
		t.Errorf("%s * %s = %s, should be %d", n1, n2, n1.Multiply(n2), 75)
	}

	ans, err := n1.Divide(n2)
	if err != nil {
		t.Errorf("%s / %s gave error: %s", n1, n2, err)
	} else if !ans.Equal(New(3)) {
		t.Errorf("%s / %s = %s, should be %d", n1, n2, ans, 3)
	}

	ans, err = n1.Divide(New(0))
	if err == nil {
		t.Errorf("%s / %s = %s, should be error", n1, New(0), ans)
	}
}

func TestPow(t *testing.T) {
	n1 := New(8)
	n2, n3 := New(2), New(3)

	if !n1.Pow(n2).Equal(New(64)) {
		t.Errorf("%s ^ %s = %s, should be %d", n1, n2, n1.Pow(n2), 64)
	}

	if !n1.Pow(n3).Equal(New(512)) {
		t.Errorf("%s ^ %s = %s, should be %d", n1, n3, n1.Pow(n3), 64)
	}
}

func TestRoot(t *testing.T) {

	n1 := New(64)

	ans := n1.Sqrt()
	if !ans.Equal(New(8)) {
		t.Errorf("sqrt(%s) = %d, should be %d", n1, ans, 8)
	}
}

func TestAbs(t *testing.T) {
	n1, n2, n3 := New(20), New(-12), New(12)

	if !n1.Abs().Equal(n1) {
		t.Errorf("|%s| = %d, should be %s", n1, n1.Abs(), n1)
	}

	if !n2.Abs().Equal(n3) {
		t.Errorf("|%s| = %d, should be %s", n2, n2.Abs(), n3)
	}
}

func TestConv(t *testing.T) {
	n := New(12.517829)

	if n.AsFloat() != 12.517829 {
		t.Errorf("%s.AsFloat() = %f, should be %f", n, n.AsFloat(), 12.517829)
	}

	if n.String() != "12.517829" {
		t.Errorf("%s.String() = %s, should be %s", n, n.String(), "12.517829")
	}
}
