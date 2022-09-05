package main

import "fmt"

func main() {
	// how to declare, instantiate, assign arrays in golang
	// we need to know the array size at compile time itself. sucks.

	// [3]int{6, 7, 8} - this is how you delcare an array literal [sizeOfArray]type{elements}

	var a [3]int        // delcare a variable of the type array
	a = [3]int{1, 2, 3} // assign
	fmt.Println(a)      // use

	var x [3]int = [3]int{1, 2, 3}   // declare and assign
	fmt.Println(x)                   // use
	y := [2]string{"hello", "world"} // declare and assign
	fmt.Println(y)                   // use

	// use ... to auto infer the size of an array. compiler just counts the items passed to an array and auto fills the number

	t := [...]int{4, 5, 6, 7, 8} // declare and assign
	fmt.Println(t)               // use

	// Access elements of an array
	chars := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	fmt.Println(chars[0]) // a
	fmt.Println(chars[1]) // b
	fmt.Println(chars[2]) // c

	// You can also get a slice of an array (a partial of an array)
	// slices are much more useful and widely used entities than actual arrays in golang
	// all slices are backed by an actual array data structure in golang

	// easiest way to get a slice of an array is
	fmt.Println(chars[0:4]) // inclusive:exclusive [a, b, c, d] (0, 1, 2, 3 indexes)
	fmt.Println(chars[3:])  // 3 to end
	fmt.Println(chars[:5])  // 0 to 4

	// you can also create slices manually by skipping the size while array var declaration

	var slice []int // This represents a slice of an unknown size but backed a fixed size array
	// we can append items to slice at runtime using append function from the builtin package
	// builtin package is automatically imported into every package

	// data is immutable in golang so we have to reassign.
	slice = append(slice, 3, 4, 5)
	fmt.Println(slice) // [3 4 5]

	// If we fill up the backing array completely then golang automatically copies it to a a new array twice the size
	// you can use the len function to check the length of an array and cap function to check the capacity of the backing array

	fmt.Printf("Capacity is %v and Length is %v\n", cap(slice), len(slice))
	var anotherSlice []int
	fmt.Printf("Capacity is %v and Length is %v\n", cap(anotherSlice), len(anotherSlice))

	// All arrays are pass by values unles you assign pointers

	// original string array
	strArray := [3]string{"Apple", "Mango", "Guava"}

	// data is passed by value
	Arraybyval := strArray

	Arraybyval[0] = "PAPAYA"
	fmt.Println(strArray)   // [Apple Mango Guava]
	fmt.Println(Arraybyval) // [PAPAYA Mango Guava]

	// data is passed by reference
	Arraybyref := &strArray
	Arraybyref[0] = "BANANA"
	fmt.Println(strArray)   // [BANANA Mango Guava]
	fmt.Println(Arraybyref) // &[BANANA Mango Guava]
}
