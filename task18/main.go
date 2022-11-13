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



const (
	CHAN_BUFFER_SIZE = 32
)

// concurrent counter
// count number of received ints by all workers
var goCounter int
// counter
// count number of sent ints by one sender
var mainCounter int

var wgSender sync.WaitGroup
var wgWorker sync.WaitGroup

// worker
// shutdowns using ctx.Cone()
// receives data with @inputChan channel
// updates @goCounter with @counterChan channel
func work(ctx context.Context, id int, inputChan <-chan int, counterChan chan int) {
	defer wgWorker.Done()
	// local counter
	// keeps data of number of received ints
	// using to update @goCounter
	localCounter := 0
	fmt.Println("Worker id = " + 
		strconv.Itoa(id) + 
		" started")
	for {
		select {
			case _, ok := <- inputChan:
				if ok {
					localCounter++
				}
			case <- ctx.Done():
				if localCounter != 0 {
					counterChan <- localCounter
				}
				fmt.Println("Worker id = " + 
								strconv.Itoa(id) + 
								" is STOPPED")
				return
			default:
				if localCounter > 0 {
					select {
						case counterChan <- localCounter:
							localCounter = 0
						default:
							time.Sleep(50 * time.Millisecond)
					}
				} else {
					time.Sleep(50 * time.Millisecond)
				}
      			
        }
    }
	
}

// sender
// shutdowns using ctx.Cone()
// sends int to worwer with @outputChan channel
// updates @mainCounter (number of sent ints)
func send(ctx context.Context, outputChan chan<- int) {
	fmt.Println("Counter started")
	defer wgSender.Done()
	for ;; {
		select {
			case <-ctx.Done():
				fmt.Println("Sender is STOPPED")
				close(outputChan)
				return
			case outputChan <- rand.Intn(10):
				mainCounter++
		}
	}
}

// counter
// receives localCounter values with @counterChan channel
// updates @goCounter
func count(counterChan chan int) {
	fmt.Println("Counter started")
	for c := range counterChan {
		// fmt.Println("Counter, counterChan = ", c)
		goCounter += c
	}
	fmt.Println("Counter is STOPPED")
}



func init() {
	rand.Seed(time.Now().UnixNano())
	goCounter = 0
	mainCounter = 0
}

func main() {
	var n int
	sendValueChan := make(chan int, CHAN_BUFFER_SIZE)
	counterChan := make(chan int)
	fmt.Println("Write number of workers:")
	fmt.Scanln(&n)
	wgSender.Add(1)
	wgWorker.Add(n)
	
	ctxSender, cancelSender := context.WithCancel(context.Background())
	ctxWorker, cancelWorker := context.WithCancel(context.Background())
	
	go send(ctxSender, sendValueChan)
	go count(counterChan)
	for i := 0; i < n; i++ {
		go work(ctxWorker, i, sendValueChan, counterChan)
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	fmt.Println("--------------------\n Terminated\n--------------------")

	cancelSender()
	wgSender.Wait()

	// wait until all workers clear the chan buffer
	// otherwise some sent values will be lost
	// (sent but not received)
	time.Sleep(1 * time.Second)
	cancelWorker()
	wgWorker.Wait()

	time.Sleep(1 * time.Second)
	close(counterChan)

	// values should be equal
	fmt.Println("Main goCounter =", mainCounter)
	fmt.Println("Go goCounter =", goCounter)
}
