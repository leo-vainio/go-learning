// arrays.go
package main

import "fmt"

func main() {
	// Arrays have fixed size and the size needs to be known at compile time. Arrays are zero-indexed.
	names := [3]string{"Leo", "Johannes", "Vainio"}
	grades := [...]string{"A", "B", "C", "D", "E", "F"} // Let the compiler count the number of elements.
	var scores [3]int
	fmt.Println(names, grades, scores)

	// Accessing elements of an array.
	scores[0] = 99
	scores[1] = 100
	scores[2] = 15
	fmt.Println("scores: ", scores, len(scores), cap(scores))

	// 2D arrays
	identityMatrix := [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix)

	// 2D slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)

	// Arrays are values in Go and not pointers like in some other languages.
	// This means an array has to be copied if an explicit pointer is not being used.
	a := [3]byte{100, 150, 200}
	b := a
	b[2] = 0
	fmt.Println(a, b)

	// Slices are references to an underlying array. Changing a value in a slice therefore
	// changes the value in the underlying array.
	roles := []string{"Admin", "User", "SuperAdmin", "Vip"}
	fmt.Println(roles, len(roles), cap(roles))

	// Slices can be created from arrays or other slices. Lower bound inclusive, upper bound exclusive.
	// All of these are referencing the same underlying array.
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d := arr[:5]
	e := arr[5:]
	f := arr[2:3]
	g := arr[:]
	h := g[:5]
	fmt.Println(arr, d, e, f, g, h, len(d), cap(d), len(e), cap(e))

	// The make function can be used to create a slice that can be used as a dynamically sized array.
	i := make([]int, 3)
	j := make([]int, 3, 100) // Specifies capacity of underlying array.
	fmt.Println(i, j, cap(i), cap(j))

	// Slices can change size. The capacity counts from the first element of the slice.
	k := []int{}
	fmt.Println(k, len(k), cap(k))
	k = append(k, 1) // expensive operation (copy)
	fmt.Println(k, len(k), cap(k))
	k = append(k, 2, 3, 4, 5)
	k = append(k, []int{1, 2, 3}...)
	fmt.Println(k, len(k), cap(k))

	// Remove operations on slices.
	q := []int{1, 2, 3, 4, 5}
	p := q[1:] // remove first element
	fmt.Println(q, p)
	p = q[:len(q)-1] // remove last element
	fmt.Println(q, p)
	p = append(q[:2], q[3:]...) // remove element from middle (this messes with the underlying array so be careful).
	fmt.Println(q, p)

	// Slices cant go back and retrieve information below their 0 index.
	l := []int{66, 77, 88, 99}
	l = l[1:]
	fmt.Println(l, len(l), cap(l))
	l = l[:] // This does not go and retrieve the first element that was "lost" two lines above.
	fmt.Println(l, len(l), cap(l))

	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
