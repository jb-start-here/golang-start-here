package main

import "fmt"

func main() {
	fmt.Println("Maps and Structs")

	// maps can be create using map keyword
	// map[key type]value type{initializer data}

	var countryCodes map[string]int //declare
	countryCodes = map[string]int{  // assign
		"USA": 1,
		"IND": 91,
		"AUS": 54,
	}
	fmt.Println(countryCodes) // use

	var countryLocales map[string]string = map[string]string{ // declare and assign
		"USA": "en-us",
		"CAN": "en-ca",
		"ENG": "en-gb",
	}
	fmt.Println(countryLocales) // use

	countryFactions := map[string]string{ // declare and assign w/ syntactic sugar
		"USA": "Allies",
		"GER": "Axis",
		"ENG": "Allies",
	}
	fmt.Println(countryFactions) // use

	// To access elements we can use square brackets

	fmt.Println(countryFactions["USA"]) // "Allies"

	// to delete elements from a map we use `delete` keyword

	fmt.Println(countryFactions) // map[ENG:Allies GER:Axis USA:Allies]
	delete(countryFactions, "ENG")
	fmt.Println(countryFactions) // map[GER:Axis USA:Allies]

	// When trying to access a value for a key that doesnt exist in map, it wont return a "nil" like value
	// rather, it returns the default value for the value type
	fmt.Println(countryFactions["ENG"]) // returns empty string

	// So to figure out if a certain key is a map hit or miss we look at the second value the accessor "[]" returns

	_, ok := countryFactions["ENG"] // keep in mind we used a := instead of =
	fmt.Println(ok)                 // returns false indicating a miss

	var value string                   // we have to declare value
	value, ok = countryFactions["GER"] // we dont have to declare ok because it was already declared above
	fmt.Println(value, ok)             // returns true indicating a hit

	// maps are all pass by reference so they will all be modified if you modify any one of the "copies"

	// Structs

	// structs in go are similar to structs in c/c++
	//  we can use type keyword from before to create a custom type of the type struct
	type Role struct {
		name         string
		job          string
		age          int
		affiliations []string
	}

	var charlie Role = Role{
		name:         "Charlie Kelly",
		job:          "Janitor/Rat Basher",
		age:          38,
		affiliations: []string{"Gruesome Twosome", "Freight Train", "Pigeon Boys"},
	}

	dennis := Role{
		name:         "Dennis Reynolds",
		job:          "Executive Vice president of WolfCola/Bartender",
		age:          42,
		affiliations: []string{"Golden Geese"},
	}

	// You can also instantiate another struct without the keys - just keep in mind the sequence of the values
	sweetDee := Role{
		"Deandra Reynolds",
		"Waitress also",
		42,
		[]string{"Golden Geese", "ZingingCutie23"},
	}

	fmt.Println(charlie, dennis, sweetDee)
	fmt.Println(charlie.affiliations) // This is how we access attributes of a struct [Gruesome Twosome Freight Train Pigeon Boys]

	// You can also include a struct inside a struct - this can be used to implement compositional patterns

	type Actor struct {
		name  string
		roles []Role
	}

	charlieDay := Actor{
		name: "Charlie Day",
		roles: []Role{
			charlie,
			Role{
				name:         "Dale Arbus",
				job:          "Dental Hygienist",
				age:          32,
				affiliations: []string{},
			},
		},
	}

	fmt.Println(charlieDay)
}
