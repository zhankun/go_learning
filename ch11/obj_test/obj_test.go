package obj_test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateObj(t *testing.T) {
	e1 := Employee{"1", "行的", 24}
	e2 := Employee{Id: "2", Name: "李思", Age: 30}
	e3 := new(Employee)
	e3.Id = "3"
	e3.Name = "dfaer"
	e3.Age = 22
	t.Log(e1)
	t.Log(e2)
	t.Log(e3)
	t.Log("--------")
	e1.Name = "test"
	t.Log(e1)

	t.Log(e1.string())
}

func (e *Employee) string() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
