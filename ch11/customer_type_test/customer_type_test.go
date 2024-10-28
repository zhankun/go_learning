package customer_type_test

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeCompute(inner IntConv) IntConv {
	return func(key int) int {
		start := time.Now()
		fmt.Println("time spent:", time.Since(start).Seconds())
		return inner(key)
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * time.Duration(op))
	return op
}

func TestCustomerType(t *testing.T) {
	compute := timeCompute(slowFunc)
	i := compute(2)
	t.Log(i)
}
