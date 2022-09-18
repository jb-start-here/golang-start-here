package main

import (
	"fmt"
	"strings"
)

// An interface is a way to defining a template for behavior.
// Conversely a struct is a way to define a template for data.
// With interfaces and structs combined you can easily model classes in golang

// `interface` keyword defines interfaces
// Lets define an interface that that requires implementation of
//
//	two methods - bark and adopt
type dog interface {
	bark(n int) string
	adopt()
}

// We can only implement an interface with methods not funcs
// if you recall we can have golang methods on any type.
// So any type can techinically implement that interface
// Lets define a struct that "implement" a dog interface
type heeler struct {
	name         string
	kennelNumber int
	adopted      bool
}

// creating a method bark with a heeler value receiver
func (h heeler) bark(n int) string {
	res := make([]string, n, n)
	for i := range res {
		if h.adopted {
			res[i] = "woof!"
		} else {
			res[i] = "whine!"
		}
	}
	return strings.Join(res, ", ")
}

// creating a method adopt with a heeler pointer receiver
// We need a pointer here because we need to mutate the receiver here.
func (h *heeler) adopt() {
	h.adopted = true
	h.kennelNumber = 0
}

// NOTE we can also call methods that take values of the receiver on a pointer of the receiver
// because of the golang syntactic sugar ((*ptr).attribute) -> ptr.attribute
// This works only for structs

// Now if you notice theres nothign special about the above struct.
// No one would know we want this struct to implement the dog interface.
// For everyone else its just another struct with two methods defined for it.

// Lets create another struct that seems to have same methods as the interfaces definition
type chihuahua struct {
	name         string
	kennelNumber int
	adopted      bool
}

func (c chihuahua) bark(n int) string {
	res := make([]string, n, n)
	for i := range res {
		if c.adopted {
			res[i] = "yip!"
		} else {
			res[i] = "growl!"
		}
	}
	return strings.Join(res, ", ")
}
func (c *chihuahua) adopt() {
	c.adopted = true
	c.kennelNumber = 0
}

func main() {
	// To actually implement the interface we have to declare a variable
	// of the type `dog`. And then golang compiler states that we can only assign
	// a value to it THAT has all the methods on it as required by the dog interface.

	// This is an implicit way to enforce interfaces as opposed to looking for a "implements" keyword
	// Its kinda like duck typing - if it has all the methods of an interface then it has implemented the interface

	var jack dog
	// We can assign it to a heeler pointer
	// Since atleast one of the methods on the dog interface receives a pointer
	// We want to assign a pointer to jack.

	jack = &heeler{
		"Jack",
		23,
		false,
	}

	// If a vars type is an interface then we cant access attributes of it
	// we can only call the methods of whatever struct type of it you assign it.

	// you can only do jack.bark() and jack.adopt() but you cant access its attributes
	// like jack.name. We have to do something else to access the attributes called type assertion.
	// To be looked at further down below

	// Because you cannot think of jack's type as a heeler struct but rather a dog interface

	fmt.Println(jack.bark(4))
	fmt.Println(jack)
	jack.adopt()
	fmt.Println(jack)
	fmt.Println(jack.bark(4))

	// An var thats declared as an interface can be thought of as being represented
	// internally by a tuple (type, value). type is the underlying
	// concrete type of the interface and value holds the value of the concrete type.
	fmt.Printf("Interface type %T value %v\n", jack, jack) // Interface type *heeler value &{Jack 0 true}
	// This changes depending on what we assign to the var of the type interface

	// Lets define an array of dogs
	dogsInThePound := [4]dog{
		&heeler{"Axle", 23, false},
		&chihuahua{"Garfield", 4, true},
		&heeler{"Megafon", 5, false},
		&chihuahua{"LordOfFlies", 6, true},
	}

	// Lets Adopt all dogs in the pound
	for _, doggo := range dogsInThePound {
		doggo.adopt()
		fmt.Println(doggo.bark(2))
	}

	// We can also so somethign like this
	describeAnyDog := func(d dog) {
		fmt.Println(d)
	}
	// This anonymous function takes any struct that implicitly implements a dog

	describeAnyDog(jack)
	for _, doggo := range dogsInThePound {
		describeAnyDog(doggo)
	}

	// An interface that has zero methods is called an empty interface.
	// It is represented as interface{}.
	// Since the empty interface has zero methods,
	// all types implement the empty interface.
	describeAnything := func(s interface{}) {
		fmt.Println(s)
	}

	describeAnything(jack)
	describeAnything(34)
	describeAnything("This is Dingles!")

	// Type assertion is used to extract the underlying value of the interface
	// aka the actual concrete entity that passed itself as an entity the interface tried to prescribe
	// We use th dot followed by parenthesis function. In the parenthesis function we pass in
	// the type of the var that was assigned to the var of the type interface.

	// in this example jack is a var of the type dog interface but we assigned it a
	// concrete value of the type heeler pointer
	// if you recall 	jack = &heeler{ "Jack", 23, false, }
	actualHeelerStructBehindJack := jack.(*heeler)

	fmt.Println(actualHeelerStructBehindJack.name)
	fmt.Println(actualHeelerStructBehindJack.adopted)
	fmt.Println(actualHeelerStructBehindJack.kennelNumber)

	// Sometimes we dont know what the actual type of the value assigned to the interface var
	// so this type assertion syntax provides us an ok variable as well

	//If the concrete type of interface var is what we passed as argument
	// then return value will have the actual underlying value of interface var
	// and ok will be true.

	// If the concrete type of interface var is not what we passed as argument
	// then ok will be false and return value will have the zero value of type we passed as argument
	//  and the program will not panic.
	maybeAChihuahua, ok := dogsInThePound[0].(*chihuahua)
	if ok {
		fmt.Println("It is a chihuahua")
		fmt.Println(maybeAChihuahua)
	} else { // This branch is executed
		fmt.Println("It is not a chihuahua")
		fmt.Println(maybeAChihuahua) // <nil> zero value of a pointer is nil
	}

	maybeAChihuahua, ok = dogsInThePound[1].(*chihuahua)
	if ok { // This branch is executed
		fmt.Println("It is a chihuahua")
		fmt.Println(maybeAChihuahua)
	} else {
		fmt.Println("It is not a chihuahua")
		fmt.Println(maybeAChihuahua) // <nil> zero value of a pointer is nil
	}

	// We can also use a switch case statement to switch on the type
	// This is a common pattern if we ever want to get the actual concrete value of an interface
	// when many types are potential implementations of the interface

	for _, doggo := range dogsInThePound {
		switch doggo.(type) {
		case *chihuahua:
			fmt.Printf("It is a chihuahua named %s\n", doggo.(*chihuahua).name)
		case *heeler:
			fmt.Printf("It is a heeler named %s\n", doggo.(*heeler).name)
		}
	}

	// Its also possible to do opposite get interface out of concrete values
	// Well, not really but we specifically we can check it a concrete value
	// satisfies a certain interface
	for _, doggo := range dogsInThePound {
		switch interfaceType := doggo.(type) {
		case dog:
			// You can directly call the bark method by referencing the interface
			fmt.Println(interfaceType.bark(1))
		default:
			fmt.Printf("Not a dog")
		}
	}
	// This `interface.(type)` is a special statement ans only works in switch statements

	// You can also ember interfaces in another interface
	type writer interface {
		write(input string)
	}

	type reader interface {
		read() string
	}

	// Go pattern is to have the single method inheritance named with suffix of "er" before
	// the method name.

	type IO interface {
		reader
		writer
	}

	// Zero value of an interface is nil
	var file IO
	fmt.Println(file)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	file.read() // this panics as invalid memory address or nil pointer dereference
}
