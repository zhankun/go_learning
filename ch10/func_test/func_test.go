package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	spent := timeSpent(slowFun)
	i := spent(5)
	t.Log(i)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * time.Duration(op))
	return op
}

func sum(ops ...int) int {
	result := 0
	for _, op := range ops {
		result += op
	}
	return result
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2))
	t.Log(sum(2, 3))
}

func clean() {
	fmt.Println("Clear resources...")
}

func TestDefer(t *testing.T) {
	defer clean()
	fmt.Println("start")
	panic("error")
}
