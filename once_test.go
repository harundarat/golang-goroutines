package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce) // This function will only be executed once, no matter how many goroutines call it.
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
