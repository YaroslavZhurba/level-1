package main

import (
	"errors"
	"fmt"
)

// Sets a bit by index @ind to value "1"
func setBitOne(number int64, ind int) (int64, error) {
	if ind < 0 || ind > 63 {
		return number, errors.New("index should be 0..63")
	}

	one := int64(1 << ind)
	if number&one == 0 {
		number = number | one
	}
	return number, nil
}

// Sets a bit by index @ind to value "0"
func setBitZero(number int64, ind int) (int64, error) {
	if ind < 0 || ind > 63 {
		return number, errors.New("index should be 0..63")
	}

	one := int64(1 << ind)
	if number&one != 0 {
		number = number ^ one
	}
	return number, nil
}

func main() {
	var a int64

	fmt.Println(a)

	a, _ = setBitOne(a, 0)
	fmt.Println(a)
	a, _ = setBitOne(a, 4)
	fmt.Println(a)
	a, _ = setBitOne(a, 2)
	fmt.Println(a)
	a, _ = setBitOne(a, 4)
	fmt.Println(a)

	a, _ = setBitZero(a, 7)
	fmt.Println(a)
	a, _ = setBitZero(a, 2)
	fmt.Println(a)
	a, _ = setBitZero(a, 4)
	fmt.Println(a)
	a, _ = setBitZero(a, 0)
	fmt.Println(a)
}
