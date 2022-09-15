package main

import (
	"fmt"
	"reflect"
)

func main() {
	// this is assigning a value to a variable
	foo := 42
	fmt.Println(foo) // 42

	// To assign a pointer of `foo` to another variable
	bar := &foo
	fmt.Println(bar) // 0x1400011c008

	// However if you print bar you get the address location of foo
	// To dereference the pointer you need to use the star operator
	fmt.Println(*bar) // 42

	// The type of bar variable is *int
	fmt.Printf("The type of bar is %T\n", bar) // *int
	// we can also use the reflect package to gather type info
	fmt.Println(reflect.TypeOf(bar)) // *int

	// we could have declared the bar variable as
	// var bar *int

	// Lets declare the a pointer variable and check the zero value of a pointer variable
	var baz *string
	fmt.Println(baz) // <nil>
	// Its a special type called nil. So always remember to check for nils if your function takes
	// a pointer variable as argument
	if baz == nil {
		fmt.Println("baz is an unassigned nil pointer variable")
	}

	// when you assign one variable to another variable - they are all pass by value
	// This includes all primitives and arrays.
	// When you assign slices and maps to one another theyre actually projections of an underlying datastructure
	// This means that they use pointers internally to compute their operations. Thats why they are pass by reference

	// Structs and pointers

	type Person struct {
		name string
		age  int
	}

	// lets declare a new variable of type pointer of a person
	var joe *Person

	fmt.Println(joe) // <nil>
	// As expected The zero value of a pointer is nil

	// lets assign value to this pointer
	joe = &Person{
		name: "Joe",
		age:  42,
	}

	fmt.Println(joe)  // &{Joe 42}
	fmt.Println(*joe) // {Joe 42}

	// We cannot do this below operation as dot operator takes precendece over * operator
	// fmt.Println(*joe.name)

	// So we have to do this
	fmt.Println((*joe).name)

	// However now there is a syntactic sugar in golang that if we run an accessor operator `.` on a struct pointer,
	// we can skip the `*` operator

	fmt.Println(joe.name)
	fmt.Println(joe.age)
}
