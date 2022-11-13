package main

import (
	"fmt"
	"math/rand"
	"sort"
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

// return minimum index of element than greater are qeual to the sample
// return -1 if no such element
// return true if there is an equal element, otherwise return false
func lowerBond(slice []int, sample int) (int, bool) {
	l := -1
	r := len(slice) - 1

	for l < r-1 {
		mid := (l + r) / 2
		if slice[mid] < sample {
			l = mid
		} else {
			r = mid
		}
	}
	if slice[r] == sample {
		return r, true
	}
	if slice[r] < sample {
		return -1, false
	}
	return r, false
}

func main() {
	array := genArray(100)
	sort.Ints(array)
	fmt.Println(array)
	sample := rand.Intn(100)
	idx, ok := lowerBond(array, sample)
	if ok {
		if idx > 0 {
			fmt.Println("Element sample =", sample,
				"is finded, index =", idx, " --", array[idx-1], array[idx])
		} else {
			fmt.Println("Element sample =", sample,
				"is finded, index =", idx, " --", array[idx])
		}
	} else {
		if idx >= 0 {
			fmt.Println("Element sample =", sample, 
				"is NOT finded, index =", idx, " -- ", array[idx])
		} else {
			fmt.Println("Element sample =", sample,
				"is NOT finded, index =", idx)
		}
	}

}
