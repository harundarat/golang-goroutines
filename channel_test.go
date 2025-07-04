package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// create a channel
	channel := make(chan string)
	// close channel after execution finnish
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Harun Al Rasyid"
		fmt.Println("Sending data finish")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
