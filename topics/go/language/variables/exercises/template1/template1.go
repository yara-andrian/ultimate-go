// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var a int
	var b bool
	var c string
	var d float64
	var e string

	// Display the value of those variables.
	fmt.Printf("var a int \t %T [%v] \n", a, a)
	fmt.Printf("var b bool \t %T [%v] \n", b, b)
	fmt.Printf("var c string \t %T [%v] \n", c, c)
	fmt.Printf("var d float64 \t %T [%v] \n\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10
	bb := true
	cc := "hello again"
	dd := 10.002
	e = "Andrian"

	// Display the value of those variables.
	fmt.Printf(" aa := 10 \t %T [%v] \n", aa, aa)
	fmt.Printf(" bb := true \t %T [%v] \n", bb, bb)
	fmt.Printf(" cc := \"hello\" \t %T [%v] \n", cc, cc)
	fmt.Printf(" dd := 10.002 \t %T [%v] \n", dd, dd)
	fmt.Printf(" e := \"Andrian\" \t %T [%v] \n\n", e, e)

	// Perform a type conversion.

	aaa := float64(10)

	// Display the value of that variable.
	fmt.Printf("aaa := float64(10) \t %T [%v] \n", aaa, aaa)
}
