//go:build ignore 

package main

import (
	"fmt"
	"os"
)

func main() {
	deferExample1()
	deferExample2()
	deferExample3()
	statOfReadmeFile()
	panickingWithRecovery()
	panickingWithAnotherRecovery()
}

// `defer` keyword defers execution of code till everything else is executed but,
// before the return value of a func is returned
// defers takes actual expressions and function calls rather than function definitions
func deferExample1() {
	defer fmt.Println("deferExample1")
	fmt.Println("end of deferExample1")
}

// end of deferExample1
// deferExample1

// Multiple defer statements when used are executed in LIFO order.
func deferExample2() {
	defer fmt.Println("first defer of deferExample2")
	defer fmt.Println("second defer of deferExample2")
	fmt.Println("end of deferExample2")
}

// end of deferExample2
// second defer of deferExample2
// first defer of deferExample2

// when you defer a piece of code the arguments you pass to it as evaluated at the time of deferrment
// not at the later execution time.
func deferExample3() {
	i := 9000
	defer fmt.Println("power level is ", i)
	i = 9001
}

// power level is 9000

// defers are mainly used to close reources like files, http/flush buffers/close connection pools
// as soon as you open them so you dont forget to do it later on.
func statOfReadmeFile() {
	file, err := os.Open("looping.go") // open a file resource
	if err != nil {                    // check for errors
		fmt.Println(err)
		return // Lets return if its an error for now...
	}
	defer file.Close() // no errors mean it is successfully opened so lets defer closing it

	// the above 5 lines of code is a common pattern in golang

	fmt.Println(file.Name())
	fmt.Println(file.Stat())
}

// `panic` function from builtin package is used for raising an error
// it raises an error and stops execution of the function that called panic
// And if it has any deffered executions pending then it just does that before panicking ("raising") and stopping execution

func panicking() {
	defer func() { // you can also defer immediately invoked anonymous functions
		fmt.Println("closing any open resources before panicking...")
	}()

	fmt.Println("beginning panicking function")

	panic("Uhoh! We got company!!")

	fmt.Println("ending panicking function") // will not be executed
} // This func is not called in main because it will stop main execution unless... we recover

// beginning panicking function
// closing any open resources before panicking...
// panic: Uhoh! We got company!!

// recover is another builtin function that enables us to rescue from panics
// once you rescue you can suppress it or repanic it to bubble up the panic higher in the call stack

// recover can be called only in deferred functions (doesnt matter if anon or not). invoking recover
// 1. will automatically stop the pending panic (remember defers happen after function ends but before panic if it ended in a panic)
// it will simply rescue from all panic
func panickingWithRecovery() {
	defer func() {
		recover()
	}()

	panic("May Day May Day Ive been hit!!")
}

// nothing will happen here

// 2. checks how a function has ended if the functon ended naturally then it returns nil if not it returns the err msg the func panicked with.
// We can use this to rescue or re raise the panic.
func panickingWithAnotherRecovery() {
	defer func() {
		if err := recover(); err != nil {
			// err is not nil so that means this func ended by panicking
			fmt.Println(err)

			// do nothing to swallow err``
			panic(err) // or re panic by calling panic again
		}
	}()

	panic("May Day May Day I'm Going down!!")
}

// From a ruby analogy
// panic - raise
// recover - rescue
// defer - ensure
