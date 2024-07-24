package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	func() {
		// SAMPLE A
		rand.Seed(time.Now().UnixNano()) // needed before Go 1.20

		if n := rand.Int(); n%2 == 0 {
			fmt.Println(n, "is an even number.")
		} else {
			fmt.Println(n, "is an odd number.")
		}

		n := rand.Int() % 2 // this n is not the above n.
		if n%2 == 0 {
			fmt.Println("An even number.")
		}

		if n%2 != 0 {
			fmt.Println("An odd number.")
		}
	}()

	func() {
		// SAMPLE B
		if h := time.Now().Hour(); h < 12 {
			fmt.Println("Now is AM time.")
		} else if h > 19 {
			fmt.Println("Now is evening time.")
		} else {
			fmt.Println("Now is afternoon time.")
			h := h // the right one is declared above
			// The just new declared "h" variable
			// shadows the above same-name one.
			_ = h
		}
		// h is not visible here.
	}()

	fmt.Println("vim-go")
}
