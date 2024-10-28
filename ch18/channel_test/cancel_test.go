package channel

import (
	"fmt"
	"testing"
	"time"
)

func cancel(ch chan struct{}) {
	// 关闭Channel
	close(ch)
}

func isCancelled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		// 没有收到消息，被阻塞的时候返回
		return false
	}
}

func TestCancel(t *testing.T) {
	ch := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCancelled(ch) {
					break
				}
				time.Sleep(time.Second * 2)
			}
			fmt.Println(i, "cancel")
		}(i, ch)
	}
	cancel(ch)
	time.Sleep(time.Second * 10)
}
