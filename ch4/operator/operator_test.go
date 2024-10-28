package operator

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	//b := [...]int{1, 2, 3, 4,5}
	c := [...]int{1, 2, 3, 8}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == c)
	t.Log(a == d)
}
