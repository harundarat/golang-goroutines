package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done waiting for condition with value:", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // Signal() will wake up one goroutine waiting on the condition
		}
	}()

	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	cond.Broadcast() // Broadcast() will wake up all goroutines waiting on the condition
	// }()

	group.Wait()
}
