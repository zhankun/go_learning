package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	// 切片声明,和数组的区别就是没有指定长度,因为切片是动态的
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 2, 4)
	t.Log(s0, len(s0), cap(s0))
	s0[0] = 9
	t.Log(s0)

	t.Log("-------------------")

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	// 会报错，数组越界，因为该切片只有3个元素
	/*t.Log(s2[0], s2[1], s2[2], s2[3])*/

	s2 = append(s2, -1)
	t.Log(s2, len(s2), cap(s2))
	s2 = append(s2, -1)
	t.Log(s2, len(s2), cap(s2))
	s2 = append(s2, -1)
	t.Log(s2, len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"}
	q2 := year[3:6]
	t.Log(q2, len(q2), cap(q2))
	q2[0] = "不知几月"
	t.Log(q2)
	t.Log(year)
	q2 = append(q2, "test")
	t.Log(q2, len(q2), cap(q2))
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	t.Log(a, b)
	// 语法错误，无法対切片进行比较
	//t.Log(a == b)
}
