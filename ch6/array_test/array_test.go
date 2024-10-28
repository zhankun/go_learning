package array_test

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 5, 6, 7}
	// 任意长度的数组
	arr2 := [...]int{1, 6, 9, 0, 1}
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2)
}

func TestArrayIterator(t *testing.T) {
	arr2 := [...]int{1, 6, 9, 0, 1}
	for index, e := range arr2 {
		t.Log(index, e)
	}
}
