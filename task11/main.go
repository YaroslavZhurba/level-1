package main

import (
	"fmt"
)

func intersect(s1 map[int]bool, s2 map[int]bool) (map[int]bool) {
	result := make(map[int]bool)
	lenS1 := len(s1)
	lenS2 := len(s2)
	// Take set with the lowest size
	if lenS1 > lenS2 {
		s1, s2 = s2, s1
	}

	// Check that both sets contain given element
	for key := range s1 {
		if s2[key] {
			result[key] = true
		}
	}

	return result
}

// Uses map[int]bool to make a set
func main() {
	s1 := map[int]bool{
		1 : true,
		5 : true,
		2 : true,
		3 : true,
		9 : true,
	}

	s2 := map[int]bool{
		3 : true,
		10 : true,
		2 : true,
		6 : true,
		9 : true,
	}

	s3 := intersect(s1, s2)
	// Answer is [2, 3, 9]
	fmt.Println(s3)
}