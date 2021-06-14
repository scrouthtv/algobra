package fastreal

import "testing"

import "github.com/scrouthtv/algobra/pkg/number"

func TestBasicMaths(t *testing.T) {
	var n1, n2 number.Number = New(15), New(5)
	if (!n1.Add(n2).Equal(New(20))) {
		t.Errorf("%s + %s = %s, should be %d", n1, n2, n1.Add(n2), 20)
	}

	if (!n1.Subtract(n2).Equal(New(10))) {
		t.Errorf("%s - %s = %s, should be %d", n1, n2, n1.Subtract(n2), 10)
	}

	if (!n1.Multiply(n2).Equal(New(75))) {
		t.Errorf("%s * %s = %s, should be %d", n1, n2, n1.Multiply(n2), 75)
	}

	ans, err := n1.Divide(n2)
	if (err != nil) {
		t.Errorf("%s / %s gave error: %s", n1, n2, err)
	} else if (!ans.Equal(New(3))) {
		t.Errorf("%s / %s = %s, should be %d", n1, n2, ans, 3)
	}
}
