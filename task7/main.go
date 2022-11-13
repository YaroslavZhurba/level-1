package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)
// ConcurrentMap
// Uses RWMutex to prevent race condition
type ConcurrentMap struct {
	data  map[int]int
	mutex sync.RWMutex
}

func NemMap() *ConcurrentMap {
	return &ConcurrentMap{
		data:  make(map[int]int),
		mutex: sync.RWMutex{},
	}
}

func (m *ConcurrentMap) Get(key int) (int, bool) {
	m.mutex.Lock()
	val, ok := m.data[key]
	m.mutex.Unlock()
	return val, ok
}

func (m *ConcurrentMap) Set(key int, value int) {
	m.mutex.Lock()
	m.data[key] = value
	m.mutex.Unlock()
}

func (m *ConcurrentMap) Delete(key int) {
	m.mutex.Lock()
	delete(m.data, key)
	m.mutex.Unlock()
}

func write(m *ConcurrentMap) {
	for {
		m.Set(rand.Intn(10), rand.Intn(10))
	}
}

func getAndPrint(m *ConcurrentMap) {
	for {
		val, ok := m.Get(rand.Intn(15))
		if !ok {
			fmt.Println("No such element with the given key")
		} else {
			fmt.Println("Element = " + strconv.Itoa(val))
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cMap := NemMap()
	go write(cMap)
	for i := 0; i < 4; i++ {
		go getAndPrint(cMap)
	}
	time.Sleep(100 * time.Millisecond)
}
