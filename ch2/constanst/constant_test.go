package constanst

import "testing"

const (
	Monday = iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable)
}
