package fastreal

import "testing"

import "github.com/scrouthtv/algobra/pkg/number"

func TestBasicMaths(t *testing.T) {
	var n1, n2 number.Number = New(15), New(5)
	t.Log(n1.Add(n2))
	t.Log(n1.Subtract(n2))
}
