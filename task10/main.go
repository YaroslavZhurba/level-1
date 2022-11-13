package main

import (
	"fmt"
	"math"
	"strconv"
)

var array = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

// Map
// Combines values by floor value with "10" accuracy
type TenMap struct{
	data map[int][]float64
}
 
func NewTenMap() (*TenMap) {
	return &TenMap{
		data : make(map[int][]float64)}
}

// From values get its' number using math.Floor
func (m *TenMap) SetFromArray(slice []float64) {
	for _, v := range slice {
		key := int(math.Floor(v / 10)) * 10
		if _, ok := m.data[key]; !ok {
			m.data[key] = make([]float64, 0)
		}
		m.data[key] = append(m.data[key], v)
	}
}

func (m *TenMap) Get(key int) ([]float64, bool){
	val, ok := m.data[key]
	return val, ok
}

func (m *TenMap) Println() {
	for key, value := range m.data {
		fmt.Print("Key = " + strconv.Itoa(key) + 
			", value = ")
		fmt.Println(value)
	}
}

func main() {
	tenMap := NewTenMap()
	tenMap.SetFromArray(array)
	tenMap.Println()
}