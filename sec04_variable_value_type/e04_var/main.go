package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initialization
var y = 43

// Declare there is a VARIABLE with the IDENTIFIER "z"
// and the that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
var z int

func main() {
	// short declaration operator
	// Declare a variable and assign a value (of a certain TYPE)
	x := 42
	fmt.Println(x)

	fmt.Println(y)
	fmt.Println(z)

	foo()
}

func foo() {
	fmt.Println("again:", y)
}
