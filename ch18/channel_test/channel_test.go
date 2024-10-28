package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 关闭Channel
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				fmt.Println("Channel关闭了")
				break
			}
		}
		wg.Done()
	}()
}

func TestChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	// 传递的参数，通道本身就是引用类型,传递会直接引用原始通道
	// 传递wg需要指针，是为了确保函数中队WaitGroup的操作直接作用于原始对象
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
