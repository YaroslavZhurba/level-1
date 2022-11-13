package main

import (
	"fmt"
	"time"
)

// Sleep implemetation using time.After
// Channel is blocked till the end of timer
func Sleep(x int) {
	<-time.After(time.Millisecond * time.Duration(x))
}


// Sleep implemetation using time.Now()
// wait until startMoment + duration > time.Now()
func Slp(x int) {
	startMoment := time.Now()
	for ;time.Now().Before(startMoment.Add(time.Millisecond * time.Duration(x))); {

	}
}

func main() {
	Sleep(2000)
	Slp(2000)
	fmt.Println("Hello")
}
