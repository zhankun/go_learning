package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[string]int{"name": 20, "age": 122}
	t.Log(m1)
	t.Log(m1["name"])
	m2 := map[string]string{}
	m2["name"] = "张三"
	m2["address"] = "杭州"
	t.Log(m2)
	t.Log(m2["fadf"])
	t.Log("--------------")
	m3 := make(map[int]int)
	for i := 0; i < 10; i++ {
		m3[i] = i
		t.Logf("%d,len=%d", i, len(m3))
	}
	if v, ok := m3[19]; ok {
		t.Logf("key存在，对应的value=%d", v)
	} else {
		t.Logf("key不存在,默认值=%d", v)
	}
}

func TestMapForeach(t *testing.T) {
	m1 := map[string]int{"name": 20, "age": 122}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
