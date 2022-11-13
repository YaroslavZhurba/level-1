package main

import (
	"errors"
	"fmt"
)

// deletes and keeps the order of elemets
// Slow
func deleteWithOrder(slice []int, idx int) ([]int, error) {
	length := len(slice)
	if idx < 0 || idx > length-1 {
		return slice, errors.New("index out of range")
	}
	copy(slice[idx:], slice[(idx+1):])
	slice = slice[:length-1]
	return slice, nil
}

// deletes and does not keep the oerder of elementss
// Fast 
func deleteWithoutOrder(slice []int, idx int) ([]int, error) {
	length := len(slice)
	if idx < 0 || idx > length-1 {
		return slice, errors.New("index out of range")
	}
	slice[idx] = slice[length-1]
	slice = slice[:length-1]
	return slice, nil
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original slice = ", slice)
	slice, _ = deleteWithOrder(slice, 4)
	fmt.Println("Delete with order slice = ", slice)
	slice, _ = deleteWithoutOrder(slice, 3)
	fmt.Println("Delete without order (fast) slice = ", slice)
}
