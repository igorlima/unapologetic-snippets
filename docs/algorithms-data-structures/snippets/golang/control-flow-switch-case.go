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
		switch n := rand.Intn(100); n % 9 {
		case 0:
			fmt.Println(n, "is a multiple of 9.")

			// Different from many other languages,
			// in Go, the execution will automatically
			// jumps out of the switch-case block at
			// the end of each branch block.
			// No "break" statement is needed here.
		case 1, 2, 3:
			fmt.Println(n, "mod 9 is 1, 2 or 3.")
			// Here, this "break" statement is nonsense.
			break
		case 4, 5, 6:
			fmt.Println(n, "mod 9 is 4, 5 or 6.")
		// case 6, 7, 8:
		// The above case line might fail to compile,
		// for 6 is duplicate with the 6 in the last
		// case. The behavior is compiler dependent.
		default:
			fmt.Println(n, "mod 9 is 7 or 8.")
		}
	}()

	func() {
		// SAMPLE B
		// Go provides a fallthrough keyword to do this task. For example, in the
		// following example, every branch code block will get executed, by their
		// orders, from top to down.
		rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
		switch n := rand.Intn(100) % 5; n {
		case 0, 1, 2, 3, 4:
			fmt.Println("n =", n)
			// The "fallthrough" statement makes the
			// execution slip into the next branch.
			fallthrough
		case 5, 6, 7, 8:
			// A new declared variable also called "n",
			// it is only visible in the current
			// branch code block.
			n := 99
			fmt.Println("n =", n) // 99
			fallthrough
		default:
			// This "n" is the switch expression "n".
			fmt.Println("n =", n)
		}
		/*
			 note:
				- a fallthrough statement must be the final statement in a branch.
				- a fallthrough statement can't show up in the final branch in a
				switch-case control flow block.
		*/
	}()

	func() {
		// SAMPLE C
		// the following code blocks are all legal, all of them can be viewed as no-ops.

		switch n := 5; n {
		}

		switch 5 {
		}

		switch _ = 5; {
		}

		switch {
		}

		switch { // <=> switch true {
		case true:
			fmt.Println("hello")
		default:
			fmt.Println("bye")
		}
	}()

	func() {
		// SAMPLE D
		// Another obvious difference from many other languages is the order of the
		// `default` branch in a switch-case control flow block can be arbitrary.

		switch n := rand.Intn(3); n {
		case 0:
			fmt.Println("n == 0")
		case 1:
			fmt.Println("n == 1")
		default:
			fmt.Println("n == 2")
		}

		switch n := rand.Intn(3); n {
		default:
			fmt.Println("n == 2")
		case 0:
			fmt.Println("n == 0")
		case 1:
			fmt.Println("n == 1")
		}

		switch n := rand.Intn(3); n {
		case 0:
			fmt.Println("n == 0")
		default:
			fmt.Println("n == 2")
		case 1:
			fmt.Println("n == 1")
		}
	}()
	fmt.Println("vim-go")
}
