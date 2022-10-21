package main

import "fmt"
//import "rsc.io/quote"

func main() {
    // x := 1.0
	// y := 2.0

	x, y := 12.0, 5.0
	//z := 3 // unused variable will occure compilation error. 

	// lesson: if-else declaration
	if x > 4 {
		fmt.Println("x is greater than 4")
	}else {
		fmt.Println("x is less or equal to 4")
	}

	//lesson: or condition
	if y == 5 || y > 5 {
		fmt.Println("y is equal to 5 or greater than 5")
	}


	// lesson: assignment and checking
	if frac := x / y; frac > 0.5 {
		fmt.Println("x is more than half of y")
	}

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}