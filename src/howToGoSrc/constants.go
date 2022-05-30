// constants.go
package main

import (
	"fmt"
)

// naming convention: same as normal variables.

const shadowed int = 16

// const block
const (
	a = iota
	b = iota
	c = iota
)

// iota is scoped to the constant block. Iota basically works as enum now
const (
	d = iota
	e
	f
)

// if we wanna disgard the 0. can be useful since 0 is the zero value of numbers int
const (
	_ = iota
	g
	h
	i
)

func main() {
	// has to be set at compile time
	const pi float64 = 3.14
	fmt.Println(pi)

	const shadowed int = 32
	fmt.Println(shadowed)

	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
