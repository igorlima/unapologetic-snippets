package main

import (
	"fmt"
	"time"
)

func main() {
	func() {
		c := make(chan int) // an unbuffered channel
		go func(ch chan<- int, x int) {
			time.Sleep(time.Second)
			// <-ch    // fails to compile
			// Send the value and block until the result is received.
			ch <- x * x // 9 is sent
		}(c, 3)
		done := make(chan struct{})
		go func(ch <-chan int) {
			// Block until 9 is received.
			n := <-ch
			fmt.Println(n) // 9
			// ch <- 123   // fails to compile
			time.Sleep(time.Second)

			done <- struct{}{}
		}(c)
		// Block here until a value is received by
		// the channel "done".
		<-done
		fmt.Println("bye")
	}()

	func() {
		c := make(chan int, 2) // a buffered channel
		c <- 3
		c <- 5
		close(c)
		fmt.Println(len(c), cap(c)) // 2 2
		x, ok := <-c
		fmt.Println(x, ok)          // 3 true
		fmt.Println(len(c), cap(c)) // 1 2
		x, ok = <-c
		fmt.Println(x, ok)          // 5 true
		fmt.Println(len(c), cap(c)) // 0 2
		x, ok = <-c
		fmt.Println(x, ok) // 0 false
		x, ok = <-c
		fmt.Println(x, ok)          // 0 false
		fmt.Println(len(c), cap(c)) // 0 2
		return
		close(c) // panic!
		// The send will also panic if the above
		// close call is removed.
		c <- 7
	}()

	func() {
		var ball = make(chan string)
		kickBall := func(playerName string) {
			for {
				fmt.Println(<-ball, "kicked the ball.")
				time.Sleep(time.Second)
				ball <- playerName
			}
		}
		go kickBall("John")
		go kickBall("Alice")
		go kickBall("Bob")
		go kickBall("Emily")
		ball <- "referee" // kick off
		var c chan bool   // nil
		return
		<-c // blocking here for ever
	}()

	// Channel Element Values Are Transferred by Copy
	func() {
		fibonacci := func() chan uint64 {
			c := make(chan uint64)
			go func() {
				var x, y uint64 = 0, 1
				for ; y < (1 << 63); c <- y { // here
					x, y = y, x+y
				}
				close(c)
			}()
			return c
		}
		c := fibonacci()
		for x, ok := <-c; ok; x, ok = <-c { // here
			time.Sleep(time.Second)
			fmt.Println(x)
		}
	}()

	func() {
		var c chan struct{} // nil
		select {
		case <-c: // blocking operation
		case c <- struct{}{}: // blocking operation
		default:
			fmt.Println("Go here.")
		}
	}()

	func() {
		c := make(chan string, 2)
		trySend := func(v string) {
			select {
			case c <- v:
			default: // go here if c is full.
			}
		}
		tryReceive := func() string {
			select {
			case v := <-c:
				return v
			default:
				return "-" // go here if c is empty
			}
		}
		trySend("Hello!") // succeed to send
		trySend("Hi!")    // succeed to send
		// Fail to send, but will not block.
		trySend("Bye!")
		// The following two lines will
		// both succeed to receive.
		fmt.Println(tryReceive()) // Hello!
		fmt.Println(tryReceive()) // Hi!
		// The following line fails to receive.
		fmt.Println(tryReceive()) // -
	}()

	func() {
		c := make(chan struct{})
		close(c)
		select {
		case c <- struct{}{}:
			// Panic if the first case is selected.
		case <-c:
		}
	}()
	fmt.Println("vim-go")
}
