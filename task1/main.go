package main

import (
	"fmt"
)

// "Base" struct
type Human struct {
	name string
	age int
}

func (human Human) Age() (int) {
	return human.age
}

func (human Human) Name() (string) {
	return human.name
}

func (human *Human) HappyByrthday() {
	human.age++
}

// "Inherited" struct
type Action struct {
	Human
	greetings string
}

// Overrided method
func (action *Action) HappyByrthday() {
	action.age += 2
}

func (action Action) SayHello(name string) {
	fmt.Println(action.greetings + name + "!")
}

func main() {
	human := Human{
		name: "Vasily",
		age:  42,
	} 

	action := Action {
		Human : Human {
			name : "Peter",
			age : 15},
		greetings : "Hello there, "}
		
	human.HappyByrthday()
	fmt.Println(human.Age())
	action.HappyByrthday()
	fmt.Println(action.Age())
	action.SayHello(human.Name())
	action.Human.HappyByrthday()
	fmt.Println(action.Age())
}