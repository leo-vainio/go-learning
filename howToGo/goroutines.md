# Goroutines in Go

```go

// goroutines.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Waitgroups can block with wg.Dait() until all goroutines are done. A goroutine
// can signal that its done with the wg.Done() method. wg.Add() just increments
// the goroutine counter. Instead of having a global variable like this we could
// also use a pointer and pass it to the goroutines.
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter int = 0

func main() {
	// Creates a new goroutine and keeps on executing the main goroutine.
	wg.Add(1)
	go f1()

	// We can use goroutines on anonymous functions.
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(1, "Anonymous function")
	}()

	// The anonymous functions can use variables declared outside of their scope. It is however
	// a better practice to pass them as arguments since they will then use the value of the varibale
	// when it was called. In the case below, msg changes value after the go function has been called
	// and therefore that function will use the changed value. This is called race condition. To check
	// for race conditions we can take help from the go compiler by running our program with:
	// $ go run -race goroutines.go
	msg := "Hello there cutie!"
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(2, msg)
	}()
	msg = "Hello there ugly hoe!"

	// In this case, the function will always use the value "Hello there ugly hoe!".
	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(3, msg)
	}(msg)
	msg = "Well hi there again!"

	// We can add more than one to the waitgroup. Notice that the goroutines execute in random
	// order. This is because they race against eachother.
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayHi()
		go sayBye()
	}

	// Using mutexes and locking values (read and write locks) we can ensure that we get the
	// behaviour we want. In this case we are just removing all benefits of concurrency. But
	// it shows how mutexes work. But we have made sure that these now execute in order.
	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.RLock()
		go sayMsg()
		m.Lock()
		go increment()
	}

	// Changing the amount of threads the application uses. This value can be set to more than
	// the amount of cores your computer have. This value can be tuned to increase performance.
	// The amount of goroutines spawned can also be a good thing to tune and test to see performance.
	// The default setting is the amount of cores of your computer.
	fmt.Println("Number of cpu cores: ", runtime.GOMAXPROCS(-1), runtime.NumCPU())
	runtime.GOMAXPROCS(4)
	fmt.Println("Number of cpu cores:", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100)
	fmt.Println("Number of cpu cores:", runtime.GOMAXPROCS(-1))

	wg.Wait()

}

func f1() {
	fmt.Println("Function 1!")
	wg.Done()
}

func sayHi() {
	fmt.Println("Hi")
	wg.Done()
}

func sayBye() {
	defer wg.Done()
	fmt.Println("Bye")
}

func sayMsg() {
	fmt.Println("Hello:", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

```