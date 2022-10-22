package main

import "fmt"

//import "rsc.io/quote"

func main() {
	mapsBasic()
}

// count how many times each word appears in a text. This is a very basic step in many text processing alogirthms.
// To split text to words, use the "Fields" fuction from the "strings" package

func mapsChallange()  {
	
}


func mapsBasic()  {
	stocks := map[string]float64{
		"AMAZON": 2084.98,
		"GOOGLE": 2540.85,
		"MICROSOFT": 287.70, // MUST havé trailing comma in multi line
	}

	// number of items
	fmt.Println(len(stocks))

	// get a value
	fmt.Println(stocks["MICROSOFT"])

	// get zero value if not found
	fmt.Println(stocks["TESLA"])

	// use two value to see if found
	value, ok := stocks["TESLA"]
	fmt.Println(value)
	fmt.Println(ok)

	if !ok {
		fmt.Println("TESLA not found")
	}else {
		fmt.Println(value)
	}

	// set
	stocks["TESLA"] = 822.12
	fmt.Println(stocks)

	// delete
	delete(stocks, "AMAZON")
	fmt.Println(stocks)

	// for range (key,value)
	for key, v := range stocks {
		fmt.Print(key + " ")
		fmt.Println(v)
	}

}

func findTheMaxValueChallange()  {
	nums := []int{16, 8, 42, 4, 23, 15}

	// write a sorting algorithm
	// bubble sort
	anySwapNeed := true

	for anySwapNeed {
		anySwapNeed = false
		for i := 0; i < len(nums) - 1; i++ {
			current := nums[i]
			next := nums[i + 1]
			if current > next {
				nums[i] = next
				nums[i+1] = current
				anySwapNeed = true
			}
		}
	}

	fmt.Println(nums[len(nums) - 1])

}

func slice()  {
	// similar to list or array
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// 0 indexing
	fmt.Println(loons[1]) // daffy

	// slices
	fmt.Println(loons[1:])

	// for
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// single value range
	for i := range loons {
		fmt.Println(i)
	}

	// doubel value range
	for i, value := range loons {
		fmt.Printf("%s at %d\n", value, i)
	}

	// append
	loons = append(loons, "elmer")
	fmt.Println(loons) 
}


func evenEndedNumbersChallange()  {
	// problem is ambigous. Due to this it has not solved yet
}

func stringBasics() {

	book := "the color of magic"
	fmt.Println(book)

	//length
	fmt.Println(len(book))

	// uint8 is a byte
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	// strings in go are immutable ()
	//book[0] = 116 // this will cause error

	// slice (start, end), 0 based, this is known as a 1/2 empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// use + to concatenate strings
	fmt.Println("T" + book[1:])

	// multi line
	poem := `
	the road goes ever on 
	down from the door where it began
	`
	fmt.Println(poem)

}

// print the numbers between 1 and 20, one per line.
// - if the number is divisible by 3 (for example, 9), print "fizz"
// - if the number is divisible by 5 (for example, 9), print "buzz"
// - if the number is divisible by 3 and 5 (for example, 15), print "fizz buzz"

func fizzBuzzChallange() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func conditionAndLoopBasic() {
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
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			break
		}
		if i == 1 {
			continue
		}
	}

	// lesson: before inializitation and increment
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

}
