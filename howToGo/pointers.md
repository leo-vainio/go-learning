# Pointers in Go

```go

// pointers.go
package main

import (
	"fmt"
)

type Animal struct {
	Name  string
	Speed int
}

func main() {
	// Address-of operator: &, dereferncing operator: *
	var name string = "Leo"
	var p *string = &name
	fmt.Println(name, &name, p, *p)
	name = "peter"
	fmt.Println(name, &name, p, *p)
	*p = "Karl"
	fmt.Println(name, &name, p, *p)

	// Pointer aritmetic is not allowed in go, i.e. we can't do &a[0] - 4.
	// There is an unsafe package that does allow this.
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)

	// Pointer to a struct.
	var bird *Animal
	bird = &Animal{
		Name:  "Leo",
		Speed: 54,
	}
	fmt.Println(bird)

	// Using the new keyword. A pointer's zero-value is nil. It's a good practice
	// to check for nil values in functions that receive pointers.
	var dog *Animal
	fmt.Println(dog)
	dog = new(Animal)
	(*dog).Name = "German shepherd" // Need to parenthesise
	dog.Speed = 50                  // Syntactic sugar
	fmt.Println(dog)
	fmt.Println(dog.Name)
}

```