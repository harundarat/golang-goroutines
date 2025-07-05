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

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Harun Al Rasyid"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Harun Al Rasyid"
}

func onlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go onlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// channel <- "Harun"
	// channel <- "Al"
	// channel <- "Rasyid"
	// // channel <- "Hehe" // This will make the channel blocked since the buffer is full

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	go func() {
		channel <- "Harun"
		channel <- "Al"
		channel <- "Rasyid"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Finish")
}
