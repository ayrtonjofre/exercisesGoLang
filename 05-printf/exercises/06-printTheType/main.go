// ---------------------------------------------------------
// EXERCISE: Print the Type
//
//  Print the type and value of 3 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3 is int
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	fmt.Printf("Type of %d is %[1]T", 3)
}
