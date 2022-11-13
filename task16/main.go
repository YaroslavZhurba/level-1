package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates array with the given size
// and random values
func genArray(size int) []int {
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = rand.Intn(100)
		// res[i] = 0
	}
	return res
}

// Check is the given array sorted
func isSorted(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}

	return true
}

// Default quick sort
// takes the middle element to make sort itaretion
func quickSort(slice []int) {
	size := len(slice)
	l := 0
	r := size - 1
	sample := slice[size/2]
	for {
		for ;slice[l] < sample; l++ {
		}
		for ;slice[r] > sample; r-- {
		}
		if l <= r {
			slice[l], slice[r] = slice[r], slice[l]
			l++
			r--
		}
		if l > r {
			break
		}
	}
	
	if (r > 0) {
		quickSort(slice[:r + 1])
	}

	if (l < size - 1) {
		quickSort(slice[l:])
	}
}

func main() {
	array := genArray(1000)
	fmt.Println(array)
	quickSort(array)
	fmt.Println(array)

	fmt.Println(isSorted(array))
}
