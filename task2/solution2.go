package main

import (
	"fmt"
	"log"
	"time"
)

func sqrtAndPrintChan(a int, c chan bool) {
	res := a * a
	fmt.Println(res)
	c <- true
}

func main2() {
	start := time.Now()
	array := []int{2, 4, 6, 8, 10}
	// Using channels to make concarrancy counting
	c := make(chan bool)
	for _, v := range array {
		go sqrtAndPrintChan(v, c)
	}
	// Wait for all channels
	for i := 0; i < len(array); i++ {
		<-c
	}
	close(c)
	elapsed := time.Since(start)
	defer log.Printf("Solution number 2 took %s", elapsed)
}