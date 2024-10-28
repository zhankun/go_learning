package condition

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			fmt.Printf("%d 是偶数", i)
		case 1, 3:
			fmt.Printf("%d 是奇数", i)
		default:
			fmt.Printf("%d 不在0-3范围内", i)
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("偶数")
		case i%2 == 1:
			t.Log("奇数")
		default:
			t.Log("不知道什么数")

		}
	}
}
