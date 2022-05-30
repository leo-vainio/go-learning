# Variables in Go

variables.go (src/variables/variables.go) shows examples of how to assign and initialize variables in Go. Some extra information are also given about variable scope, naming conventions and type conversion.

```go

package main

// Single-line comment.

/*
 * General comment.
 */

/*
 Naming conventions in go:
 variable length should reflect the lifetime legnth of the variable.
 acronyms should be uppercase: googleURL, not googleUrl

 PascalCase or camelCase
*/

import (
	"fmt"
	"strconv"
)

// Cannot use := syntax in package scope
var firstName string = "Leo"

var (
	lastName     string = "Vainio"
	benchPressPR int    = 115
)

// only visible to package (package level scope) so every file with this package can see this.
var a int = 42

// There is no private scope. so we cant scope a variable to this file specifically

// exported variable (globally visible)
var A int = 42

func main() {
	// block scope
	var age int
	age = 22

	// All variables must be used in a go application

	var height int = 185

	// the variable with the innermost scope takes precedence (shadowing)
	var lastName string = "NotVainio"

	var (
		x int = 1
		y int = 2
	)

	// converty type (destinationType(var)).
	// go does not do implicit type conversion: var z float32 = x, would not work
	z := float32(x)

	// strconv for converting with strings
	var k string = strconv.Itoa(age)

	weight := 84

	fmt.Println(age, height, weight, firstName, lastName, benchPressPR, x, y, z, k)
}

```
