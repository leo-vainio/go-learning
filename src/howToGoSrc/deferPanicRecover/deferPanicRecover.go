// deferPanicRecover.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Defer pushes the function calls onto a stack, these get called at the end of the function
// in last in first out order (LIFO). I.e the deferred functions get called before the function returns.
func f1() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

// The defer keyword is often used to close files. This is otherwise pretty easy to forget. This means
// we can associate the opening and closing of a resource at the same place which is a lot easier.
func f2() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

// When you defer a function it takes the argument at the time of being called and uses that. In this case "Leo"
// will be printed and not "Kalle".
func f3() {
	name := "Leo"
	defer fmt.Println(name)
	name = "Kalle"
}

// Panic, basically Go's version of exceptions.
func f4() {
	a, b := 1, 0
	fmt.Println(a / b)
}

// We can cause panic by calling the built in panic function.
func f5() {
	panic("Bad things are happening")
}

// Go check localhost:8080 in the browser. Starting this application with two different terminals will cause
// a panic since the port will be blocked by the other terminal.
func f6() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

// Panics happen after deferred functions. Deferred gets printed but last does not. This means that opened resources
// can be closed even if the application panics.
func f7() {
	fmt.Println("First")
	defer fmt.Println("Deferred")
	panic("Bad things happened here")
	fmt.Println("Last")
}

// The panic exits this function but is recovered by the defer. The main function keeps
// executing.
func f8() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// panic(err)
		}
	}()
	panic("Bad things happened here")
	fmt.Println("Paniced")
}

func main() {
	f1()
	// f2()
	f3()
	// f4()
	// f5()
	// f6()
	// f7()

	fmt.Println("start")
	f8()
	fmt.Println("end")
}
