package type_test

import "testing"

func TestConverter(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}
