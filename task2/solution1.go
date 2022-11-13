package main

import (
	"fmt"
	"log"
	"time"
)

func sqrtAndPrint(a int) {
	fmt.Println(a * a)
}

func main1() {
	start := time.Now()
	array := []int{2, 4, 6, 8, 10}
	for _, v := range array {
		go sqrtAndPrint(v)
	}
	// Using time.Sleep to make concarrancy counting
	time.Sleep(500 * time.Microsecond)
	elapsed := time.Since(start)
	defer log.Printf("Solution number 1 took %s", elapsed)
}
