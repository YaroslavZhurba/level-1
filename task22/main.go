package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result = mul(result, big.NewInt(int64(i)))
	}
	return result
}

func add(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Add(a, b)
	return result
}

func mul(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(a, b)
	return result
}

func sub(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Sub(a, b)
	return result
}

func div(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Div(a, b)
	return result
}

func main() {
	// initialization
	a := factorial(100)
	b := factorial(99)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---------------")

	fmt.Println("Sum:", add(a, b), "\n -----")
	fmt.Println("Sub:", sub(a, b), "\n -----")
	fmt.Println("Mul:", mul(a, b), "\n -----")
	fmt.Println("Div:", div(a, b), "\n -----")

}
