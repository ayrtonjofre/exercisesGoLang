package main

import "fmt"

func main() {
	// var name string = "Linguica"
	// var name = "Linguica"
	name := "Linguica"

	// var isScientist bool = true
	// var isScientist = true
	isScientist := true

	// var age int = 62
	// var age = 62
	age := 62

	// var degree float64 = 5.
	// var degree = 5.
	degree := 5.

	fmt.Println(name, isScientist, age, degree)

	// type inference also works for variables
	//
	// Go gets the type of the variable and assigns it
	//   to the newly declared variable
	//
	// The type of the name2 variable is `string` now
	name2 := name
	fmt.Println(name2)
}
