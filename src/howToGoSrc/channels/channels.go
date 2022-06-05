// channels.go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// We need to use the make keyword, and the channel can only send and receive messages
	// of the specified type. The zero value of channels are nil. Channels are reference types.
	ch1 := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := <-ch1 // Receive data from channel
		fmt.Println(i)
	}()
	go func() {
		defer wg.Done()
		ch1 <- 555 // Send data to channel
	}()
	wg.Wait()

	// We can create multiple sender and receivers that use the same channel. We need to make
	// sure that all messages that get sent also get received to not cause a deadlock.
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			j := <-ch1
			fmt.Println(j)
		}()
		go func() {
			defer wg.Done()
			ch1 <- 555
		}()
	}
	wg.Wait()

	// A goroutine can be both a sender and a receiver. Although the normal thing to do is to dedicate
	// a goroutine to be only one of those.
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := <-ch1
		fmt.Println(i)
		ch1 <- 111
	}()
	go func() {
		defer wg.Done()
		ch1 <- 555
		fmt.Println(<-ch1)
	}()
	wg.Wait()

	// We can enforce the direction of a channel (sending or receiving) by passing the
	// channel as an argument to the function with the types seen below.
	wg.Add(2)
	go func(ch <-chan int) { // Receive only
		defer wg.Done()
		i := <-ch
		fmt.Println(i)
	}(ch1)
	go func(ch chan<- int) { // Send only
		defer wg.Done()
		ch <- 555
	}(ch1)
	wg.Wait()

	// We can make buffered channels. This means the channel can hold a specified amount
	// of values internally. We can therefore send values to this channel without receivinng
	// right away. Below we send three messages and receive only one. This would cause a
	// deadlock if the channel were not buffered. We do however not use two messages.
	ch2 := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		defer wg.Done()
		i := <-ch
		fmt.Println(i)
	}(ch2)
	go func(ch chan<- int) {
		defer wg.Done()
		ch <- 555
		ch <- 444
		ch <- 333
	}(ch2)
	wg.Wait()

	// We can range over the channel to receive messages. The sender can close the channel
	// to indicate that there are no more messages coming. The for range loop detects that
	// and will stop. You dont need to close a channel but its useful if the receiver need
	// to be told there are no more values coming.
	ch3 := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch3)
	go func(ch chan<- int) {
		defer wg.Done()
		ch <- 555
		ch <- 444
		ch <- 333
		ch <- 222
		ch <- 111
		close(ch)
	}(ch3)
	wg.Wait()

	// We can also check if the channel is closed in the way shown below, by retreiving
	// a second value from the channel. ok will be false if there are no more values to
	// be received and the channel is closed.
	ch4 := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i, ok)
			} else {
				break
			}
		}
		wg.Done()
	}(ch4)
	go func(ch chan<- int) {
		defer wg.Done()
		ch <- 555
		ch <- 444
		ch <- 333
		ch <- 222
		ch <- 111
		close(ch)
	}(ch4)
	wg.Wait()

	// The select statement.
	ch5 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch5)
		}
		quit <- 0
	}()
	fibonacci(ch5, quit)

	// The default case is run if no other cases are ready. This can be used to send and
	// receive without blocking.
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// The select statement lets a goroutine wait on multiple channels. A select blocks until
// one of the cases can run and then executes that case. It chooses randomly if multiple
// are ready.
func fibonacci(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
