package main

import "fmt"

var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I am ASSIGN the VALUE "Shaken, not stirred"
var z = "Shaken, not stirred"

// This is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

var a = `"James said,
Shaken,
  not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	// cannot assign to the z
	//z = 43;

	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
