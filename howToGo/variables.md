# Variables in Go

variables.go (src/variables/variables.go) shows examples of how to assign and initialize variables in Go. Some extra information are also given about variable scope, naming conventions and type conversion.

### Naming conventions in Go
- Variable length should be as short as reasonable. It should reflect the lifetime of the variable in question.
- Acronyms such as URL or HTTP should be all uppercase and not Url or Http in variable names.
- Use PascalCase or camelCase.

```go

// variables.go
package main

import (
	"fmt"
	"strconv"
)

// The type can be stated explicitly or inferred from the value on the right hand side.
var color1 string = "Purple"
var color2 = "Red"

// Var declaration block, a shorthand for declaring variables in bulk. Usually variables that are related somehow.
var (
	firstName string = "Leo"
	lastName  string = "Vainio"
	height    int    = 185
)

// The var keyword can be used to assign and initialize a list of variables.
var c, python, java string = "c", "python", "java"

// Variables declared without an explicit initial value are given their zero value and can be assigned another value later on.
var isBlue, isYellow bool

// Variables beginning with a lower case letter are package scoped (only visible within this package).
var notExported int = 100

// Variables beginning with a capital case letter are exported. When importing a package, you can refer only to its exported names.
// Any "unexported" names are not accessible from outside the package. There is no private scope, i.e. we can't scope a variable to this file specifically.
var Exported int = 400

var shadowed = "package scope"

func main() {
	fmt.Println(color1, color2, firstName, lastName, height, c, python, java, isBlue, isYellow, notExported, Exported, shadowed)

	// Short hand declaration. Inferres type by looking at the value on the right side. Cannot use := syntax in package scope because every statement
	// outside of a function begins with a keyword (var, func, import etc).
	weight := 84
	fmt.Println(weight)

	// Block scope (inside a function, if-statement, for loop etc).
	var age int = 22
	fmt.Println(age)

	// A variable cannot be declared more than one time within the same scope but it can be declared multiple times if they are declared in different scopes.
	// The variable with the innermost scope takes precedence. This is called shadowing.
	var shadowed = "function scope"
	if true {
		var shadowed = "if scope"
		fmt.Println(shadowed)
	}
	fmt.Println(shadowed)

	// We can convert types by using T(v), which converts value v to type T.
	// Go does not do implicit type conversion: var y float64 = x, would not work.
	var x int = 50
	var y float64 = float64(x)
	fmt.Println(x, y)

	// The strconv package can be used for converting with strings.
	var ageString string = strconv.Itoa(age)
	fmt.Println(ageString)

	// Variables of different types can be declared on the same line as such:
	var name1, id1 = "Henrik", 5055
	var name2, id2 = string("Moa"), int(2353)
	name3, id3 := "Karin", 5534
	fmt.Println(name1, id1, name2, id2, name3, id3)
}

```
