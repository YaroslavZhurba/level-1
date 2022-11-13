package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 27
	fmt.Println(a, b)
	
	// go trick
	a, b = b, a
	fmt.Println(a,b)

	// math trick
	b = a ^ b
	a = a ^ b
	b = a ^ b
	fmt.Println(a,b)
}