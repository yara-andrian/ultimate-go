// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var sliceInt []int
	// Append numbers to the slice.
	for numbers := 1; numbers <= 10; numbers++ {
		// value := fmt.Sprintf("Rec: %d", numbers)
		sliceInt = append(sliceInt, numbers)
		fmt.Println(sliceInt)
		// println("%v\n", *sliceInt)
	}
	// Display each value in the slice.

	// Declare a slice of strings and populate the slice with names.
	names := []string{"Bull", "James", "Boll", "Bill"}

	// Display each index position and slice value.
	for i, name := range names {
		fmt.Printf("Index: %d Names: %s\n", i, name)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	sliceName := names[2:3]
	fmt.Printf("location %p value:%s", &sliceName, sliceName)

	for _, v := range names[1:3] {
		fmt.Println("step:", v)
	}
	// Display each index position and slice values for the new slice.

	sliceCopy := make([]string, len(names))
	copy(sliceCopy, names)
}
