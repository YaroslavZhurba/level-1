package main

import (
	"context"
	"fmt"
	"time"
)

// Stops using closed channel
func first(с chan int) {
	for value := range с {
		fmt.Println(value)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("First STOPPED")
}

// Stops using receiving data from channel
func second(quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Println("Second STOPPED")
			return
		default: 
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// Stops using context
func third(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Third STOPPED")
			return
		default: 
		time.Sleep(10 * time.Millisecond)
		}
	}
}

var flag bool

// Stops using global flag
func fours() {
	for {
		if flag {
			fmt.Println("Fours STOPPED")
			return
		}
	}
}

func main() {
	quit := make(chan bool)
	channelFirst := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go first(channelFirst)
	channelFirst <- 1
	time.Sleep(time.Second)
	close(channelFirst)

	go second(quit)
	time.Sleep(time.Second)
	quit <- true
	
	go third(ctx)
	time.Sleep(time.Second)
	cancel()

	go fours()
	time.Sleep(time.Second)
	flag = true

	time.Sleep(time.Second)

}