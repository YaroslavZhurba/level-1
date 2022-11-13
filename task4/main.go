package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"strconv"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	CHAN_BUFFER_SIZE = 32
)

// Worker
// Receives data from the Sender using main channel
// Many workers
// Using Context.Done() and select expression to finish all gourutines
func work(ctx context.Context, id int, c <-chan int) {
	defer wg.Done()
	for {
		select {
			case value := <- c:
				fmt.Println("Worker id = " + 
					strconv.Itoa(id) + 
					", print value = " +  
					strconv.Itoa(value))
			case <- ctx.Done():
				fmt.Println("Worker id = " + 
					strconv.Itoa(id) + 
					" is STOPPED")
				return
			default:
      			time.Sleep(50 * time.Millisecond)
        }
    }
	
}

// sender
// Sends data to the main channel
// Only one sender
func send(ctx context.Context, c chan int) {
	defer wg.Done()
	for ;; {
		select {
			case <-ctx.Done():
				fmt.Println("Sender is STOPPED")
				close(c)
				return
			case c <- rand.Intn(10):
				// time.Sleep(10 * time.Millisecond)
		}
	}
}

var wg sync.WaitGroup

func main() {
	var n int
	// Main channel to make interaction between
	// sender and workers
	c := make(chan int, CHAN_BUFFER_SIZE)
	fmt.Scanln(&n)
	wg.Add(n+1)
	// create context to finish all gorutines
	ctx, cancel := context.WithCancel(context.Background())

	go send(ctx, c)
	for i := 0; i < n; i++ {
		go work(ctx, i, c)
	}
	// channel to receive termination signals
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	fmt.Println("--------------------\n Terminated\n--------------------")
	cancel()
	//
	wg.Wait()
}
