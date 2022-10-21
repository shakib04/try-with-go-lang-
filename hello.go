package main

import "fmt"

//import "rsc.io/quote"

func main() {
	// x := 1.0
	// y := 2.0

	x, y := 12.0, 5.0
	//z := 3 // unused variable will occure compilation error.

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)

	// lesson: if-else declaration
	if x > 4 {
		fmt.Println("x is greater than 4")
	} else {
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

	// lesson: switch case declare
	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("unknown")
	}

	// lesson: switch case with different approach
	switch {
	case n > 1:
		fmt.Println("greater than one")
	case n == 1:
		fmt.Println("equal to one")
	default:
		fmt.Println("unknown")
	}

	// lesson: for loop
	for i := 0; i < 5 ; i++{
		fmt.Println(i)
		if  i == 2 {
			break
		}
		if i == 1 {
			continue
		}
	}


	// lesson: before inializitation and increment
	a := 0
	for a < 3{
		fmt.Println(a)
		a++
	}


	



}
