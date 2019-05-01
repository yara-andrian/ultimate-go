// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffered channel so no send ever blocks. Don't allocate more capacity
// than you need. Have the main goroutine store each random number it receives
// in a slice. Print the slice values then terminate the program.
package main

// Add imports.
import (
	"fmt"
	"math/rand"
)

// Declare constant for number of goroutines.
var numbGo int

func main() {
	room := 100

	// Create the buffered channel with room for
	buffChannel := make(chan int, room)
	// each goroutine to be created.

	// Iterate and launch each goroutine.
	for i := 0; i < room; i++ {
		go func() {
			buffChannel <- rand.Intn(1000)
		}()
	}
	// Create an anonymous function for each goroutine that
	// generates a random number and sends it on the channel.
	var nums []int
	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	for room > 0 {
		p := <-buffChannel
		nums = append(nums, p)
		room--
	}

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	fmt.Println(nums)
	// Print the values in our slice.
}
