package main

import (
	"fmt"
	"sync"
	"time"
)

var Revision string
var counter = 0
var m sync.Mutex

func increment() {
	for i := 0; i < 10000; i++ {
		m.Lock()
		counter++
		m.Unlock()
	}
}

func decrement() {
	for i := 0; i < 10000; i++ {
		m.Lock()
		counter--
		m.Unlock()
	}
}

func main() {

	go increment()
	go decrement()

	time.Sleep(1 * time.Second)

	fmt.Println("Counter: ", counter)
}
