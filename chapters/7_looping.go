package main

import "fmt"

func main() {
	// only one looping in golang thats `for`

	// some examples of for statement (slight syntax modificaiton/sugar etc)

	for i := 0; i < 5; i++ { // i is only scoped inside the for block.
		fmt.Println(i) //  Also make sure `i` doesnt conflict with a var of the same name outside this scope
	}

	printDashes()

	k := 1 // you can start from any number
	for ; k < 5; k++ {
		fmt.Println(k)
	}

	printDashes()

	j := 0
	for ; ; j++ {
		if j == 3 {
			continue // skip this iteration
		}
		if j >= 5 {
			break // this keyword is used to exit a loop
		}
		fmt.Println(j) // prints 0 1 2 4
	}

	printDashes()

	counter := 1
	// this loops infinitely unless you break from inside
	for {
		if counter >= 8 {
			fmt.Println("This is too much beer!!")
			break
		}
		fmt.Printf("%d bottle(s) of beer.\n", counter)
		counter++
	}

	printDashes()

	// nested loops
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			fmt.Println(a, b)
			// if you break here, you only break out of the inner most loop.
		}
	}

	printDashes()

	// you can break out of any loop you want here if you label your loop and break to there
outerLoop:
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			fmt.Println(a, b)
			// if you break here, you only break out of the inner most loop.
			if a > 2 {
				break outerLoop // this breaks out all loops upto outerLoop
			}
		}
	}

	printDashes()

	// Looping through collections
	vegetables := [3]string{"cauliflower", "Broccoli", "Carrot"}
	for index, vegetable := range vegetables { // range keyword makes a collection into an iterable.
		fmt.Println(index, vegetable)
	} // if you dont need index then just use write only `_` variable

	printDashes()

	fruits := []string{"Apple", "Banana", "Orange"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	printDashes()
	countryCodes := map[string]int{
		"US": 1,
		"CA": 2,
	}
	for k, v := range countryCodes {
		fmt.Println(k, v)
	}
}

func printDashes() {
	fmt.Println("--------------------------------")
}
