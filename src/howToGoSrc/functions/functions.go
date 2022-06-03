// functions.go
package main

import "fmt"

// The naming convention with upper- and lower case for exporting and importing is also true for functions.
func main() {
	printMessage("Hello there", 50)
	name := "Anton"
	sayGreeting("Hi", name, 505)
	pointGreeting("Hi there", &name)
	fmt.Println(name)
	sum(1, 2, 3, 4, 5, 6)
	fmt.Println(sum1(1, 3))
	s := sum2(1, 4)
	fmt.Println(*s)
	fmt.Println(sum3(5, 5))
	d, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)

	// Anonymous function.
	func() {
		fmt.Println("Hej")
	}()

	// Passing argument to anonymous function.
	func(i int) {
		fmt.Println(i)
	}(5)

	// Functions can be varables.
	f := func() {
		fmt.Println("Hello Go!")
	}
	f()

	// Methods.
	g := greeter{
		"Leo",
		"Hello",
	}
	g.greet()

	a := animal{
		"Ostrich",
	}
	a.printAnimal()

}

// Multiple parameters.
func printMessage(message string, index int) {
	fmt.Println(message, index)
}

// If we have paramters of the same type we can list them as below.
func sayGreeting(message, name string, index int) {
	fmt.Println(message, name, index)
}

// Passing pointers to functions. Passing a pointer is a lot more efficient on large pieces of data.
func pointGreeting(message string, name *string) {
	fmt.Println(message, *name)
	*name = "Greger"
}

// Ellipsis, pass a variable amount of arguments to a function.
func sum(values ...int) {
	fmt.Println(values)
	res := 0
	for _, v := range values {
		res = res + v
	}
	fmt.Println(res)
}

// Return values.
func sum1(a, b int) int {
	return a + b
}

// Returning a pointer.
func sum2(a, b int) *int {
	result := a + b
	return &result
}

// Named return values. The variable sum is implicitly returned.
func sum3(a, b int) (sum int) {
	sum = a + b
	return
}

// Returning multiple values.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}

// Methods.
type greeter struct {
	Name     string
	Greeting string
}

// Any type can have methods. For example: type counter int. In this case we are getting
// a copy of the greeter struct.
func (g greeter) greet() {
	fmt.Println(g.Name, g.Greeting)
}

// Methods with pointer.
type animal struct {
	Name string
}

// We can now manipulate the same object that we called the method with.
func (a *animal) printAnimal() {
	fmt.Println((*a).Name)
	fmt.Println(a.Name)
}
