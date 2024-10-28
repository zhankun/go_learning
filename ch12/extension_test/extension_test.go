package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) Action() {
	fmt.Println("开始动起来")
	p.Speak()
}

type Dog struct {
	//p *Pet
	Pet
}

func TestExtension(t *testing.T) {
	d := new(Dog)
	//d.p.Action()
	d.Action()
}
