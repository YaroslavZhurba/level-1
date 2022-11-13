package main

import (
	"fmt"
)

// Implementainon of the conveyor
// input is a slice.
// Values from the slice are moving to the firstChannel
// using firstPipe method. Second pipe receives values,
// multiplys them and sends to the secondChannel (output channel)

func firstPipe(slice []int, c chan<- int) {
	for _, v := range slice {
		c <- v
	}

	close(c)
}

func secondPipe(inChan <-chan int, outChan chan<- int) {
	for v := range inChan {
		outChan <- v * 2
	}

	close(outChan)
}

func main() {
	slice := []int{2, 4, 6, 8, 10}

	firstChan := make(chan int)
	secondChan := make(chan int)

	go firstPipe(slice, firstChan)
	go secondPipe(firstChan, secondChan)

	for v := range secondChan {
		fmt.Println(v)
	}
}

