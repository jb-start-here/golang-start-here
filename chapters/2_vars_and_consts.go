// Golang is a strongly typed language
// this page explains how variables are declared, assigned, exported and scoped within golang.
// We can also take a look at how to use constants

package main

import "fmt"

func main() {
	fmt.Println("Variables, Constants and Types")
	printDashes()

	//We use the var keyword followed by name and then type because it feels natural to say
	// create a `var` called `language` of the type `string`

	// There are two ways of declaring variables. 
	// First kind is to declare first first and then assign it
	var language string;
	language = "en"
	
	// Second kind is to declare and assign in one line.
	var country string = "us"

	fmt.Printf("%s-%s\n", language, country) // => en-us

	// There is a another way to do the second type and its a golang syntactic sugar
	// We cant redeclare an already declared variable so for example;
	// `var nation string = "ca"` can be rewritten using golang syntactic sugar as
	
	nation := "ca"
	// This syntactic sugar was created to mimic some dynamically typed languages like ruby for example;
	// the `:=` symbol automatically figures the type of the value being assigned and created a var names nation with the inferred type

	printDashes()
	fmt.Printf("%s-%s\n", language, nation) // => en-ca

	// to declare constants we use 
	const defaultGreeting string = "Hello World"
	// defaultGreeting = "Goodbye World" => throws an error

	// consts can also be declared without type. (it actually is typed internally by default as string..there is no such thing an untyped in golang)
	// The value of a constant should be known at compile time. We cannot assign values to a constant, the result of a func or anything else.

	// IMPORTANT: If a variable/func name is pascal cased then its exported variable of this package.
	// If a variable/func name is camel cased then its not exported

	// vars and constants defined in curly braces are scoped to within the braces only even if theyre pascal cased!

	// Vars and consts can be mass/bulk declared and assigned!!!

	printDashes()
	var (
		name string = "Charlie"
		quirk string = "Illiteracy"
		job string = "Janitor/Rat Basher"
		age int = 42
	)

	fmt.Println(name, quirk, job, age)

	printDashes()
	const (
		animal string = "dog" 
		sound = "bark!" // this is also valid; remember consts by default is string so you dont need to declare types 
		numberOfLegs int = 4 // if its anything other than string then we must declare it 
	)
	fmt.Println(animal, sound, numberOfLegs)

	// for readability Acronyms can be in uppercase. this is not a rule but best practice

	var theURL string = "https://github.com/start-here/golang-start-here"
	fmt.Println(theURL)

	HTTPMethod := "GET"
	fmt.Println(HTTPMethod)

	// for write only variables you can use _. Also, you dont have to declare an _ so you can just use `=` not `:=`
	_ = returnAString
}

func returnAString() string {
	return "Ignore me!!"
}


// non exported
func printDashes() {
	fmt.Println("--------")
}

// Exported
// we could do `main.TestFunc` in another module.
func TestFunc() {
	fmt.Println("FIGHT MILK!!!!")
}
