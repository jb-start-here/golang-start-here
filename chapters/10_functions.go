//go:build ignore 

package main

import (
	"Strings"
	"fmt"
	"strconv"
)

func main() {
	// main mfunction in main package is the entry point for all apps.
	// cli tools, apps and other compiled to be executables all have main
	// modules may not have main function.

	// Read through the functions below the main function.
	// The main function is only used to demonstrate the functions

	looper(4)
	msg := constructGreeterMsg("Jack", "Morning")
	fmt.Println(msg)
	fmt.Println(add3Numbers(1, 1, 1))
	fmt.Println(joinWithDashes("Hello", "World", "Dinosaur"))
	fmt.Println(joinWithDashes("Hello", "World", "Dinosaur", "Carpet"))

	nums := [3]int{4, 5, 6}
	makeAll3Zero(&nums) // pass a pointer to nums
	fmt.Println(nums)

	fmt.Println(yoloImplicit(4))

	anonymounFunctionContainer()

	str, ok := doCalc(true)
	if ok {
		fmt.Println("Sucessfully calculated")
		fmt.Println(str)
	}

	vibe, err := positiveVibesOnly(-45)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(vibe)
	}

	joe := Person{
		"Joe",
		"Schmoe",
	}
	fmt.Println(joe.greeter())
	joe.legallyChangeNameTo("Joseph", "Schmoseph")
	fmt.Println(joe.greeter())

	var mySpecialInt SpecialCustomInt
	mySpecialInt = SpecialCustomInt(43) // 43 is an int though so we should type cast it

	mySpecialInt.valueReceiverMethod()
	mySpecialInt.pointerReceiverMethod()
}

// example of fuction delcaration with arguments
func looper(times int) {
	for i := 1; i <= times; i++ {
		fmt.Println("Counter: ", i)
	}
}

// example of function with return value and argument
func constructGreeterMsg(name string, time string) string {
	return fmt.Sprintf("Hello, %s! Good %s!!", name, time)
}

// if you all multiple arguments of the same type then you can just declare type once
// at the end of the comma spearated list
func add3Numbers(a, b, c int) int {
	return a + b + c
}

// If you dont know the nunber of arguments you have then you can just use
// rest operator for (These are variadic functions)

// `...` is the operator. You must prefix it before the type
// The arguments will be populated into the variable as a slice of the variadic type specified
func joinWithDashes(strs ...string) string {
	fmt.Println("The following strings will be joined with dashes:")
	for i, str := range strs {
		fmt.Printf("\t%d. %s\n", i+1, str)
	}

	// Join method from the strings package is useful to join strings
	return strings.Join(strs, "-")
}

// One thing to note is that when we pass around data to functions as arguments
// those arguments will be passed by value (aka a copy will be passed to the function)
// unlesss ofcourse the argument happens to be a slice or a map because as we know
// they are always copied by reference (due to them just being pointers to the underlying data structure)
// To get around this pass the pointers as arguments to functions
func makeAll3Zero(numArray *[3]int) {
	(*numArray)[0] = 0
	(*numArray)[1] = 0
	(*numArray)[2] = 0
}

// You can also predeclare the return value of a function in the first line itself
// The function body scope will have that varuable already declared and ready to use
// You dont need even need to return the value explicitly, just return but the function
// auto returns the predeclared return variable

// To use it add a variable name to the return type of the function and wrap it in parenthesis

// implicit returns
func yoloImplicit(num int) (result string) {
	// Notice we didnt have to use "result := " because its already defined
	// we can use the `strconv` package to type case int to string
	result = "YOLO x" + strconv.Itoa(num)

	// this function returns result automatically even though you didnt specify it in the return statement
	return
}

// you can return multiple values from a function. In fact its a common pattern to return
// the vaule and an additional ok or err value to say of the execution was successfully rather than simple panicking.
func doCalc(mockControl bool) (string, bool) {
	if mockControl {
		return "Calculated", true
	} else {
		return "Calculated", false
	}
}

// instead of returning bool to signify success you can also use an error type
// you can also pack a msg in here so its more useful for debugging purposes or be more actionable
func positiveVibesOnly(num int) (int, error) {
	// fmt.Errorf is used to create an error object
	if num < 0 {
		return 0, fmt.Errorf("Negative numbers not allowed")
	} else if num == 0 {
		return 0, fmt.Errorf("Zero is neither positive nor negative")
	} else {
		return num, nil // nil because we dont want this case to be a number
	}
}

// anonymous functions
func anonymounFunctionContainer() {
	// funtions are first class citizens in golang. So we can create anonmous functions,
	// we can create IIFEs and even assign them to a variable. Anonymous functions are not "hoisted"
	// so make sure to define them before calling them.

	// The syntax is the exact same as named functions - just skip the name.

	func() {
		fmt.Println("Hello, Anonymous World!")
	}() // Immediate function invocation

	// We can assign function to a variables. That means functions are a type too
	var a func()
	a = func() {
		fmt.Println("Goodbye, Anonymous World!")
	}
	a()

	// More complex example;
	var b func(cookies map[string]string) bool
	b = func(cookies map[string]string) bool {
		if _, ok := cookies["Auth"]; ok {
			return true
		}

		return false
	}

	fmt.Println(b(map[string]string{
		"Auth": "BEARER cccccbbujtekdgktlcdguddrjrgljichgvtlfitdttek",
	}))

	// anonymous functions also act as a great scope gates
	x := 45
	func() {
		x := 56
		fmt.Println(x) // 56
	}()
	fmt.Println(x) // 45
}

// Methods

// Methods are functions that are evaluated against the context of a certain type
// If you define a function as a method of a certain type, then you call this function
// from any variable of the type as if the function was a method of the type.
// It is most commonly used with struct types.

// Structs are the closes thing to classes in golang and methods are a way to spoof behavior of class
// like in methods of ruby classes for example

// The syntax for this is to just use another set of parenthesis between `func` keyword and func name
// This context is called the receiver.
// Method = Function + Receiver

// Two types of receivers
// Value receiver makes a copy of the type and pass it to the function.
// The function stack now holds an equal object but at a different location on memory.
// That means any changes done on the passed object will remain local to the method.
// The original object will remain unchanged.

// Pointer receiver passes the address of a type to the function.
// The function stack has a reference to the original object.
// So any modifications on the passed object will modify the original object.
type Person struct {
	fname string
	lname string
}

func (p Person) greeter() string {
	return "Hello! " + p.fname + " " + p.lname
}

// to actually modify the receiver object then we need to create a method with pointer receiver
func (p *Person) legallyChangeNameTo(newFName, newLName string) {
	p.fname = newFName
	p.lname = newLName

	// it should actually be (*p).fname = newFName but golang as syntactic sugar exception for pointers of
	// structs as its inferred that you want attrs of the actual struct and not its pointer
}

// if `joe` was a var of type person struct then you can call `joe.greeter()`

// We've only seen methods on structs .i.e the receiver has been a struct or a pointer to a struct
// BUT, technically, the receiver of a method can be any type in golang. We just need to make sure
// we create a type wrapper or a type clone. (this is diffferent from type aliasing)

// This is because primitive of inbuilt singular unit types cannot be receivers
// Struct however is not a singular unit type its a collection of a shape of types.

// If we were to say create a custom type that just wraps an int then we could make it
// a receiver because golang compiler doesnt consider it a primitive type anymore...

// func (receiver *int) method() will not work
// func (receiver int) method() will not work

// However
type SpecialCustomInt int // Here were just cloning int type into a new entity and giving it a custom name
func (receiver SpecialCustomInt) valueReceiverMethod() {
	fmt.Println(receiver, "Good Number!")
}
func (receiver *SpecialCustomInt) pointerReceiverMethod() {
	fmt.Println(*receiver, "Great Number!")
}
