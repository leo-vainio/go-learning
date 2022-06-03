// looping.go
package main

import "fmt"

func main() {
	// The most basic for loop. The curly braces are required.
	for i := 0; i < 4; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// Multiple loop variables.
	for i, j := 0, 0; i < 4; i, j = i+1, j+1 {
		fmt.Print(i, j)
	}
	fmt.Println()

	// We don't need to use all components in the for loop.
	i := 1
	for ; i < 4; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// Some syntactic sugar for creating a while loop in Go.
	j := 1
	for j < 4 {
		fmt.Print(j)
		j++
	}
	fmt.Println()

	// Infinite loop. We break out of a for loop using the break keyword. The break
	// breaks out of the innermost loop (if we have nested loops).
	for {
		break
	}

	// The continue statement forwards to the next iteration in the for loop.
	for i := 0; i < 4; i++ {
		fmt.Print("Printed")
		continue
		fmt.Print("Not printed")
	}
	fmt.Println()

	// A label can be used to break out of the outer loop from an inner loop.
Loop:
	for {
		for {
			break Loop
		}
	}
	fmt.Println("We broke out of the infinite loop")

	// Looping over collections. The syntax works for arrays, slices, maps, strings and channels.
	s1 := []string{"Leo", "Vainio", "Peter", "Anna"}
	for k, v := range s1 {
		fmt.Println(k, v)
	}

	// Looping over string. _ can be used to discard a value.
	s2 := "Leo Vainio"
	for _, v := range s2 {
		fmt.Print(string(v))
	}
	fmt.Println()

	// If we only need the keys we can use this syntax. This does not work for if we only want the values.
	s3 := []string{"Leo", "Vainio", "Peter", "Anna"}
	for k := range s3 {
		fmt.Println(s3[k])
	}
}
