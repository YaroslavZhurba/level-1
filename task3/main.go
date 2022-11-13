package main

import (
	"fmt"
)

func sumArray(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v*v
	}
	c <- sum
}

func main() {
	array := []int{2,4,6,8,10}
	middle := len(array) / 2
	c1 := make(chan int)
	c2 := make(chan int)
	// split calculations between two gouruitines
	go sumArray(array[:middle], c1)
	go sumArray(array[middle:], c2)

	x,y := <- c1, <-c2
	close(c1)
	close(c2)

	fmt.Println(x + y)
	
}