package main

import (
	"fmt"
)

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	// use map to remove dublicates
	set := make(map[string]int)

	
	for _, v := range array {
		set[v] += 1 
	}

	result := make([]string, len(set))
	
	i := 0
	for v := range set {
		result[i] = v
		i++
	}

	fmt.Println(result)
}