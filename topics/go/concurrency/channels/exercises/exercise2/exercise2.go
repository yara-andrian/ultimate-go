// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffer channel so no send every blocks. Don't allocate more buffers than
// you need. Have the main goroutine display each random number is receives and
// then terminate the program.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	goroutines = 100
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	wg.Add(goroutines)
	// Create the buffer channel with a buffer for
	// each goroutine to be created.
	values := make(chan int, goroutines)

	// Iterate and launch each goroutine.
	for gr := 0; gr < goroutines; gr++ {
		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			defer wg.Done()
			n := rand.Intn(1000)
			if n%2 == 0 {
				return
			}
			values <- n
		}()
	}

	go func() {
		wg.Wait()
		close(values)
	}()

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	var nums []int
	// for wait > 0 {
	// 	nums = append(nums, <-values)
	// 	wait--
	// }

	for n := range values {
		nums = append(nums, n)
	}

	// Print the values in our slice.
	fmt.Println(nums)
}
