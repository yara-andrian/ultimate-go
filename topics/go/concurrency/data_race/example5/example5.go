// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how maps are not safe for concurrent use by default.
// The runtime will detect concurrent writes and panic.
package main

import (
	"fmt"
	"sync"
)

// scores holds values incremented by multiple goroutines.
var scores = make(map[string]int)

// Mutex is used to define a critical section of code.
var Mutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		Mutex.Lock()
		for i := 0; i < 1000; i++ {
			scores["A"]++
		}
		Mutex.Unlock()
		wg.Done()
	}()

	go func() {
		Mutex.Lock()
		for i := 0; i < 1000; i++ {
			scores["B"]++
		}
		Mutex.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final scores:", scores)
}
