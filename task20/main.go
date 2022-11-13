package main

import (
	"fmt"
	"strings"
)

// reverseWords
// make slice of words ising Split
// reverse words
// Join to the string and return the value
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog golang sun"
	ss := "snow dog java golang sun"
	fmt.Println(s)
	rws := reverseWords(s)
	fmt.Println(rws)

	fmt.Println()
	
	fmt.Println(ss)
	rwss := reverseWords(ss)
	fmt.Println(rwss)
}
