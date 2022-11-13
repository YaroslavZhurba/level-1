package main

import (
	"fmt"
	"reflect"
)

// Uses "reflect" to get type in runtime
func getType(a interface{}) reflect.Type {
	return reflect.TypeOf(a)
}

type Foo struct {
	str string
	bar int
}

func main() {
	a := 5
	b := 1.15
	c := Foo{str: "sre", bar: 4}

	fmt.Println(getType(a), getType(b), getType(c))
}
