package async_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Second * 2)
	return "Done"
}

func otherTask() string {
	time.Sleep(time.Second * 3)
	return "Success"
}

func AsyncService() chan string {
	// 单独的Channel,会阻塞,service return会最后输出
	ch := make(chan string)
	// buffer chan，不会阻塞,service return会先输出
	//ch := make(chan string, 1)
	go func() {
		ret := service()
		ch <- ret
		fmt.Println("service return")
	}()
	return ch
}

func TestAsync(t *testing.T) {
	asyncService := AsyncService()
	task := otherTask()
	t.Log(task)
	t.Log(<-asyncService)
	time.Sleep(time.Second * 5)
}
