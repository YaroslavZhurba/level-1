package main

import (
	"fmt"
	"math"
)

// struct Point
type Point struct {
	x float64
	y float64
}

// Construct of the Point
func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Default getter
func (p *Point) X() float64 {
	return p.x
}

// Default getter
func (p *Point) Y() float64 {
	return p.y
}

// Counts distance using "math" module
func distance(p1 *Point, p2 *Point) float64 {
	diffX := p1.X() - p2.X()
	diffY := p1.Y() - p2.Y()

	return math.Sqrt(diffX*diffX + diffY*diffY)
}

func main() {
	p := NewPoint(30, 40)
	fmt.Println(distance(NewPoint(0, 0), NewPoint(3, 4)))
	fmt.Println(distance(NewPoint(0, 0), p))
}
