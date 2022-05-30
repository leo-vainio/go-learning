# Primitives in Go

// u cant mix types of the same family (float32 + float64) for example

```go

package main

import (
	"fmt"
)

// Booleans (zero value: false)
var boolZero bool
var isSwedish bool = true
var isNorwegian bool = false
var equals bool = 1 == 1

// Integers (zero value: 0)
var age = 22
var height int8 = 100      // int, int8, int16, int32, int64,
var balance uint64 = 55555 // uint, uint8, uint16, uint32, uint64, uintptr (rarely used), byte (alias for uint8)
var character byte = '#'

// Floating-point (zero value: 0)
var pi float32 = 3.14 // float32, float64

// Complex (zero value: 0 + 0i)
// the real and imaginary part consist of float32 or float64
var c complex128 = 1 + 4i // complex64, complex128

// String (zero value: "")
var name string = "Leo"

// Rune (type alias for int32, zero value: 0)
var r rune = 'b'

func main() {
	fmt.Println(boolZero, isSwedish, isNorwegian, equals)
	fmt.Println(age, height, balance, character, pi, c, name, r)

	// Integer literals
	var num1 int = 14 // Decimal
	num1 = 0xE        // Hex
	num1 = 0o14       // Octal
	num1 = 0b10011    // Binary
	num1 = 1_1_1_11_1
	num1 = 0xBaD_fAcE
	num1 = 0b_1111
	fmt.Println(num1)

	// Floating-point literals
	num2 := 0.5
	num2 = 2.
	num2 = 0009.3
	num2 = 10.e1
	num2 = 6.67428e-11
	num2 = 1e6
	num2 = .55
	num2 = 0.15e+0_2
	fmt.Println(num2)
}

```
