//go:build ignore 

package main

import "fmt"

func main() {
	// basic if statement
	if true {
		fmt.Println("It's true!!!!")
	} else {
		fmt.Println("It's false!!!!")
	}

	// basic if else if statement with logical operators
	if 1 == 1 && 2 == 1 {
		fmt.Println("First!")
	} else if 1 != 1 || 2 <= 1 {
		fmt.Println("Second!")
	} else {
		fmt.Println("Third!")
	}

	// We can create variables and scope it inside the if block itself
	// the expresion after the `;` is evaluated as the boolean to control the if statement
	if x := sixty() == 60; x {
		fmt.Println("It's sixty!!!!")
	}

	switch 60 { // First kind of switch statement
	case 55:
		fmt.Println("its 55!!!!")
	case 56, 57, 58, 59:
		fmt.Println("Its either 56 or 57 or 58 or 59!!!!")
	case 60:
		fmt.Println("its 60!!!!")
	default:
		fmt.Println("its default!!!!")
	}

	switch y := sixty(); y { // alternate First kind of switch statement
	case 55:
		fmt.Println("its 55!!!!")
	case 56, 57, 58, 59:
		fmt.Println("Its either 56 or 57 or 58 or 59!!!!")
	case 60:
		fmt.Println("its 60!!!!")
	default:
		fmt.Println("its default!!!!")
	}

	// second type of switch statement
	switch {
	case sixty() != 60:
		fmt.Println("its not 60")
	case sixty() == 60:
		fmt.Println("its 60")
		fallthrough // This fallthrough keyword does what it says if a case statement hits,
		// it doesnt just stop there it follows through to the next case statement in line below
	case sixty()%2 == 0:
		fmt.Println("its even")
		fallthrough
		// weirdly,when we fallthrough the next case statment will be executed regardless of true or false
	case sixty()%2 != 0:
		fmt.Println("its odd") // This will also be printed. weird.
	}
}

func sixty() int {
	return 60
}
