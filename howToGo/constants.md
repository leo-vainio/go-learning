# Constants in Go

Constants in go use the same naming conventions as normal variables do. They are evaluated at compile time so they need to be assigned values that can be evaluated
at compile time or it will cause an error. Constants that are not given an explicit type works just like literals do, i.e. they replace the "name" with the literal value. Hence a constant given a value 5 for example can work with all integer types (and float).


```go

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
}

```
