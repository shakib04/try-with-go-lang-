package main

import "fmt"
//import "rsc.io/quote"

func main() {
    // x := 1.0
	// y := 2.0

	x, y := 1.0, 2.0
	z := 3 // unused variable will occure compilation error. 

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}