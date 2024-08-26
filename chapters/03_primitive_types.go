//go:build ignore 

package main

import "fmt"

func main() {
	// https://go.dev/ref/spec#Types

	// Short documentation link

	// the `type` keyword is used to create custom types or create type aliases and
	// https://go.dev/ref/spec#Type_definitions
	// https://go.dev/ref/spec#Type_identity

	type AString = string //alias
	// AString == string
	type ANumber int // custom type (although it just points to a int it doesnt create an alias - creates a new type)
	// ANumber != int

	var a AString = "Hello"
	var b ANumber = 45

	fmt.Printf("%v is of type %T\n", a, a)
	fmt.Printf("%v is of type %T\n", b, b)

	// Hello is of type string
	// 45 is of type main.ANumber

	// Every type has a default value.
	// when a variable of certain type is declared but not initialized yet then the the variable actually already has a value
	// this value is called is the default value
	// for int - 0
	// for string - ""
	// for bool - false
}
