package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	mut := sync.Mutex{}
	count := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() { mut.Unlock() }()
			mut.Lock()
			count++
		}()
	}
	time.Sleep(time.Second * 2)
	t.Logf("counter = %d", count)
}

func TestCounterWait(t *testing.T) {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	var count = 0
	for i := 0; i < 5000; i++ {
		// 协程等待
		wg.Add(1)
		go func() {
			defer func() { mu.Unlock() }()
			mu.Lock()
			count++
			// 当前协程结束
			wg.Done()
		}()
	}
	// 等待所有协程结束,会阻塞,直到所有任务结束
	wg.Wait()
	t.Logf("coun = %d", count)
}
