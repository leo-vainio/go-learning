// constants.go
package main

import (
	"fmt"
)

// Constants can be shadowed just like normal variables.
const shadowed int = 16

// A const block together with iota can be used to create an enum in Go.
const (
	a = iota
	b = iota
	c = iota
)

// Iota only need to be specified on the first row, same expression will be applied to the others also.
const (
	d = iota + 5
	e
	f
)

// Use '_' to disgard a value.
const (
	_ = iota
	g
	h
	_
	i
)

// Arithmetic can be used together with iota.
const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
)

func main() {
	const pi float64 = 3.14
	fmt.Println(pi)

	const shadowed int = 32
	fmt.Println(shadowed)

	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Println(g, h, i)
	fmt.Println(KB, MB, GB)

	const num1 = iota
	const num2 = iota
	fmt.Println(num1, num2)

}
