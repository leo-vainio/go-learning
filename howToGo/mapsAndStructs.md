# Maps and Structs in Go

```go

// mapsAndStructs.go
package main

import (
	"fmt"
	"reflect"
)

// Naming syntax with uppercase and lower case also works on structs. upper case are exported and lower case are within package.
// This includes the fields of the structs.
type Doctor struct {
	number     int
	name       string
	companions []string
}

// Go does not support inheritance. Uses composition instead by embedding a struct in another struct.
// Without naming that embedded field as seen below in the Bird struct.
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// Tags.
type Student struct {
	Name string `required max:"100"`
	Age  int
}

func main() {
	// ----- MAPS ----- //

	// Create a map: map[key]value{}
	// Keys need to be able to be tested for equality and the zero-value for a map is nil.
	// The ordering of a map is not guaranteed.
	cityPopulations := map[string]int{
		"Stockholm":  1_000_000,
		"Malmö":      500_000,
		"Göteborg":   700_000,
		"Eskilstuna": 100_000,
	}
	fmt.Println(cityPopulations)

	// Create map using the make function: make(map[key]value) OR make(map[key]value, size)
	nameHeight := make(map[string]int)
	nameHeight = map[string]int{
		"Leo":  185,
		"Karl": 190,
	}
	fmt.Println(nameHeight)

	// Adding a new key and deleting a key.
	nameHeight["Filip"] = 500
	delete(nameHeight, "Karl")
	fmt.Println(nameHeight)

	// Manipulating values in map.
	fmt.Println(cityPopulations["Stockholm"])
	cityPopulations["Stockholm"] = 5
	fmt.Println(cityPopulations["Stockholm"])

	// Comma ok syntax. Referencing a key that is not in the map returns the zero-value for the maps element type
	// and does not create an error. We can however check for this using ok, ok will be false if the key is not in
	// the map and true if it is in the map. The naming convention for this is to use ok.
	population, ok := cityPopulations["Gööööööteborg"]
	if ok {
		fmt.Println(population)
	}

	// Check how many elements are in a map.
	fmt.Println("Number of elements:", len(nameHeight))

	// Maps are reference types. Deleting a key from cp in this case also deletes the value from cityPopulations.
	cp := cityPopulations
	delete(cp, "Stockholm")
	fmt.Println(cityPopulations["Stockholm"])

	// ----- STRUCTS ----- //

	// The name of the fields can be ommitted and just inferred by the position. This is usually not recommended
	// since it can create maintanence issues if new fields get added. When using the naming syntax, the order
	// of the fields does not matter when creating a struct.
	doctor1 := Doctor{
		number: 3,
		name:   "Leo",
		companions: []string{
			"Leif",
			"Greger",
			"Fritchof",
		},
	}
	fmt.Println(doctor1)

	// Accessing a field of a struct.
	doctor1.name = "Peter"
	fmt.Println(doctor1.name)

	// Anonymous structs. Usually used if the struct is shortlived.
	doctor2 := struct{ name string }{name: "Leif"}
	doctor3 := struct{ name string }{"Leif"}
	fmt.Println(doctor2, doctor3)

	// Structs are passed by value. Be careful!
	doctor4 := doctor3
	doctor4.name = "NotLeif"
	fmt.Println(doctor3, doctor4)

	// The fields of Animal can be accessed directly via the syntax below. this is syntactic sugar and not inheritance.
	bird := Bird{}
	bird.Name = "Ostrich"
	bird.Origin = "Sweden"
	bird.SpeedKPH = 100
	bird.CanFly = false
	fmt.Println(bird)

	// Tags. Go doesnt do anything with the tag but we can read in this tag and then decide what to do with it ourselves.
	s := reflect.TypeOf(Student{})
	field, _ := s.FieldByName("Name")
	fmt.Println(field.Tag)
}

```