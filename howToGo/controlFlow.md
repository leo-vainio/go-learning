# Control flow in Go

```go

// controlFlow.go
package main

import "fmt"

// Boolean values: true, false
// Boolean operators: <, >, ==, !=, <=, >=,
// Logical operators: !,  ||, &&

func main() {
	// The curly braces are necessary.
	if true {
		fmt.Println("TRUE")
	}

	// If with a short statement. Variables declared in if statement (see population and ok)
	// are only accessable in the scope of the if-statement.
	cityPopulations := map[string]int{
		"Stockholm":  1_000_000,
		"Malmö":      500_000,
		"Göteborg":   700_000,
		"Eskilstuna": 100_000,
	}
	if population, ok := cityPopulations["Stockholm"]; ok {
		fmt.Println(population)
	}

	// if, else if, else
	if 1 < 0 {
		fmt.Println("IF")
	} else if 0 > 1 {
		fmt.Println("ELSE IF")
	} else {
		fmt.Println("ELSE")
	}

	// Go implements short circuit evalutation. As soon as it finds the first value being true,
	// it doesn't care about the rest of the values.
	if true || false {
		fmt.Println("SHORT CIRCUIT")
	}

	// Switch statements. The break keyword is not needed.
	switch 2 {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("DEFAULT")
	}

	// Multiple values can be tested for in every case. There cannot be the same value
	// in multiple cases.
	switch 6 {
	case 1, 2, 5:
		fmt.Println(1)
	case 3, 4, 6:
		fmt.Println(2)
	default:
		fmt.Println("DEFAULT")
	}

	// Just like the if-statement we can use an initializer.
	switch i := 2 - 1; i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("DEFAULT")
	}

	// Switch statement without condition is the same thing as writing: switch true. With
	// this syntax the cases can overlap as seen below. The first case that is true will execute.
	i := 50
	switch {
	case i < 50:
		fmt.Println("Less than 50")
	case i < 100:
		fmt.Println("Less than 100")
	default:
		fmt.Println("Greater or equal to 100")
	}

	// We can use the fallthrough keyword if we want multiple cases to execute. If the first case is true
	// it will fallthrough to the second case no matter if that case is true or not.
	j := 25
	switch {
	case j < 50:
		fmt.Println("Less than 50")
		fallthrough
	case j < 100:
		fmt.Println("Less than 100")
	default:
		fmt.Println("Greater or equal to 100")
	}

	// Type switching.
	var k interface{} = 1.1
	switch k.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Float")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Other type")
	}

	// Break can be used to exit out of a switch statement early.
	switch 1 {
	case 1:
		fmt.Println("Printed")
		break
		fmt.Println("Not Printed")
	default:
		fmt.Println("DEFAULT")
	}
}

```