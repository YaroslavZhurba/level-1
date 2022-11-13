package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Create Waiting Group
// to wait all gourutines
var wg sync.WaitGroup

func sqrtAndPrintWG(a int) {
	defer wg.Done()
	res := a * a
	fmt.Println(res)
}

// The best implementation
func main3() {
	start := time.Now()
	array := []int{2, 4, 6, 8, 10}
	// add nubmer of gorutines
	wg.Add(len(array))
	for _, v := range array {
		go sqrtAndPrintWG(v)
	}
	// Wait all workers
	wg.Wait()
	elapsed := time.Since(start)
	defer log.Printf("Solution number 3 took %s", elapsed)
}
