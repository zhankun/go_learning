package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `
{
	"basic_info": {
		"name": "杰克",
		"age": 24
	},
	"job_info": {
		"skills": ["Java", "Go", "Python"]
	}
}
`

func TestJson(t *testing.T) {
	e := Employee{}
	err := json.Unmarshal([]byte(jsonStr), &e)
	if err != nil {
		t.Log("转为对象错误")
	}
	if marshal, err := json.Marshal(e); err != nil {
		t.Log("对象转为字符串失败")
	} else {
		t.Log(string(marshal))
	}
}

func TestEasyJson(t *testing.T) {
	e := Employee{}
	err := e.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		t.Log("json解析失败")
	}
	if eJson, err := e.MarshalJSON(); err != nil {
		t.Log("对象转为json失败")
	} else {
		t.Log(string(eJson))
	}
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			fmt.Println("json解析失败")
		}
		if _, err = e.MarshalJSON(); err != nil {
			fmt.Println("对象转json失败")
		}
	}
}

func BenchmarkJson(b *testing.B) {
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), &e)
		if err != nil {
			fmt.Println("json解析失败")
		}
		if _, err := json.Marshal(e); err != nil {
			fmt.Println("对象转json失败")
		}
	}
}
