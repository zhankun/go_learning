package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "hello world"
}

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	s := p.WriteHelloWorld()
	t.Log(s)
}
