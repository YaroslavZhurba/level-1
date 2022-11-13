package main

import (
	"fmt"
)

func createHugeString(n int) string {
	str := ""
	str += "体"
	for i := 0; i < n; i++ {
		str += "а"
	}
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	// string is a type alias for byte[], but there is unicode
	// and for one character spends several bytes 
	// but slice will take 100 bytes, not symbols
	vc := v[:100]
	fmt.Println(vc)
}

// converts string to rune, where one cell is one charachter
// slice will take 100 symbols, then convrt to string again to print it
func someFuncNew() {
	v := []rune(createHugeString(1 << 10))
	vc := string(v[:100])
	fmt.Println(vc)

}

func main() {
	someFunc()
	someFuncNew()
}
