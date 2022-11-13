package main

import (
	"fmt"
	"strings"
)

// isUnique
// Uses strings.ToLower to make all characters slowercase
// it means the program is caseinsensetive
// Uses map like a hash table to check 
// if the letter does not appear twice
func isUnique(str string) bool {
	// Use
	hashTable := make(map[rune]bool, len(str))
	str = strings.ToLower(str)
	for _, v := range str {
		_, ok := hashTable[v]
		if ok {
			return false
		}
		hashTable[v] = true
	}

	return true
}

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefA"))
	fmt.Println(isUnique("aabcd"))
}
