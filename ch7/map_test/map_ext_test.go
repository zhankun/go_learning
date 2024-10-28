package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	m := map[int]bool{}
	m[1] = true
	key := 4
	if m[key] {
		t.Logf("%d 存在", key)
	} else {
		t.Logf("%d 不存在", key)
	}
	// 从map中删除元素
	delete(m, 1)
	delete(m, 2)
	t.Log(m)
}
