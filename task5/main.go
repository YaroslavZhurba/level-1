package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"strconv"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


func receive(ctx context.Context,  c <-chan int) {
	defer wg.Done()
	for {
		select {
			case value := <- c:
				fmt.Println("Print value = " +  
					strconv.Itoa(value))
			case <- ctx.Done():
				fmt.Println("Receiver is STOPPED")
				return
			default:
      			time.Sleep(50 * time.Millisecond)
        }
    }
	
}

func send(ctx context.Context, c chan int) {
	defer wg.Done()
	for ;; {
		select {
			case <-ctx.Done():
				fmt.Println("Sender is STOPPED")
				close(c)
				return
			case c <- rand.Intn(10):
				time.Sleep(10 * time.Millisecond)
		}
	}
}

var wg sync.WaitGroup

const (
	WORK_TIME_SECONDS = 3
)

func main() {
	c := make(chan int)
	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())

	go send(ctx, c)
	go receive(ctx, c)
	// The same like a privious task 4
	// but now the program waits WORK_TIME_SECONDS
	// and then terminates all gorutines
	// Another way to implement is use ccontext with timer
	time.Sleep(time.Second * WORK_TIME_SECONDS)
	fmt.Println("------------------------\n Finished\n------------------------")
	cancel()
	wg.Wait()
}
