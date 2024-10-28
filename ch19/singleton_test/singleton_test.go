package singleton_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func init() {
	singleInstance = new(Singleton)
}

func getSingleInstance() *Singleton {
	once.Do(func() {
		fmt.Println("create...")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		/*		go func() {
				instance := getSingleInstance()
				fmt.Printf("%x\n", unsafe.Pointer(instance))
				wg.Done()
			}()*/
		go func() {
			fmt.Printf("%x\n", unsafe.Pointer(singleInstance))
			wg.Done()
		}()
	}
	wg.Wait()
}
