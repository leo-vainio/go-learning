# Interfaces in Go

```go

// interfaces.go
package main

import (
	"bytes"
	"fmt"
)

// Interfaces define behaviour, an interface type is a set of method signatures. A
// value of interface type can hold any value that implements those methods.
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// We can embed interfaces in other interfaces. In this case a type must implement the
// methods from Writer and Closer in order to implement WriterCloser.
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// Constructor function. Used since we have to initialize the internal buffer.
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type Person interface {
	show()
}

type Student struct {
	name string
}

func (s *Student) show() {
	if s == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(s.name)
	s.name = "HEJ"
}

type person struct {
	Name string
	Age  int
}

// The Stringer interface, defined in the fmt package, is one of the most common interfaces.
// A Stringer is a type that can describe itself as a string. The fmt package and others look
// for this interface to print values.
func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	// Interfaces can be thought of as a tuple of a value and a concrete type. There are
	// value receivers and pointer receivers. Pointer receiver are most common since we
	// usually have to modify the object we are calling the method with. Methods with value
	// receivers can be called with both pointers and value types whilst methods with
	// pointer receivers can only be called with pointer types.
	var p1 Person = &Student{"Leo"}
	fmt.Printf("(%v, %T)\n", p1, p1)
	p1.show()
	fmt.Printf("Value changed: (%v, %T)\n", p1, p1)

	// If the concrete value in the interface is nil, the method will be called with a nil
	// receiver. The method can handle this as seen above.
	var p2 Person
	var s *Student
	p2 = s
	fmt.Printf("(%v, %T)\n", p2, p2)
	p2.show()
	p2 = &Student{"Petter"}
	fmt.Printf("(%v, %T)\n", p2, p2)
	p2.show()

	// The zero-value for an uninitialized interface type is nil. The nil interface
	// value holds neither a value nor a concrete type. Callin a method on a nil interface
	// will cause a run-time error since it doesn't contain a type to indicate which method
	// to call.
	var t interface{}
	fmt.Printf("The nil interface value: (%v, %T)\n", t, t)

	// The empty interface. The interface type that specifies zero methods. An empty interface
	// can hold values of any type since every type implements at least zero methods. Empty
	// interfaces are used by code that handles values of unknown type. any is an alias for
	// the empty interface.
	var i interface{}
	var empty any
	fmt.Printf("(%v, %T)\n", empty, empty)
	fmt.Printf("(%v, %T)\n", i, i)
	i = NewBufferedWriterCloser()
	fmt.Printf("(%v, %T)\n", i, i)
	i = 42
	fmt.Printf("(%v, %T)\n", i, i)

	// Type assertions provides access to an interface value's underlying concrete value.
	// If the interface does not hold the specified type, the statement will trigger a panic.
	// We can however test whether an interface value holds a specific type. A type assertion
	// can return two values, the underlying value and a boolean value that reports whether
	// the assertion succeeded. If it doesn't hold that type, the boolean will be false and
	// the value will receive its zero value. No panic will occur.
	var j interface{} = "Hello"
	fmt.Printf("(%v, %T)\n", j, j)
	a := j.(string)
	fmt.Println(a)
	a, ok := j.(string)
	fmt.Println(s, ok)
	f, ok := j.(float64)
	fmt.Println(f, ok)

	// Type switch with empty interface using the keyword type.
	var k interface{} = 0
	switch k.(type) {
	case int:
		fmt.Println("Integer")
	default:
		fmt.Println("I have no idea what type this is")
	}

	// WriterCloser can use methods from both embedded interfaces.
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello there you ugly piece of shit!"))
	wc.Close()

	// Testing the Stringer interface.
	y := person{"Arthur Dent", 42}
	z := person{"Zaphod Beeblebrox", 9001}
	fmt.Println(y, z)
}

```