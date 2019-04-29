// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// Add imports.

// Add user type and provide comment.

func main() {

	// Declare variable of type user and init using a struct literal.
	type json struct {
		name        string
		age         int
		phoneNumber string
	}

	var e3 struct {
		name string
		age  int
	}

	var e1 json

	// Display the field values.
	fmt.Printf("this is the json struct %+v\n", e1)

	// Declare a variable using an anonymous struct.
	e2 := json{
		name:        "andrian",
		age:         10,
		phoneNumber: "12345678",
	}

	// Display the field values.
	fmt.Println("Name", e2.name)
	fmt.Println("Age", e2.age)
	fmt.Println("PhoneNumber", e2.phoneNumber)

	e3.name = "Andrian"
	fmt.Printf("%+v \n", e3)
}
