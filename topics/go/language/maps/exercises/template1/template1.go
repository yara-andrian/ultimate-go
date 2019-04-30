// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

// Add imports.
// type hobby struct {
// 	activity string

// }

// type player struct {
// 	name string
// 	hobbies slice[]
// }

func main() {

	// Declare and make a map of integer type values.
	var readMap map[string]int
	something := make(map[string]int)
	// Initialize some data into the map.
	fmt.Println(len(readMap))
	fmt.Println(len(something))
	something["james"] = 12
	something["maria"] = 13
	fmt.Println(len(something))
	// Display each key/value pair.
	for key := range something {
		fmt.Println("key:", key, "value:", something[key])
	}
}
