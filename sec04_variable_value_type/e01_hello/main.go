package main

import "fmt"

func main() {
	foo()
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("Hey, this is foo")
}
