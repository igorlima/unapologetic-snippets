package main

import "fmt"

func main() {
	func() {
		// SAMPLE A
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	func() {
		// SAMPLE B
		// condition-only for loop form.
		var i = 0
		for i < 10 {
			fmt.Println(i)
			i++
		}
		for i < 20 {
			fmt.Println(i)
			i++
		}
	}()

	func() {
		// SAMPLE C
		//  Condition portion is absent, compilers will view it as true.
		for i := 0; ; i++ { // <=> for i := 0; true; i++ {
			if i >= 10 {
				// "break" statement will be explained below.
				break
			}
			fmt.Println(i)
		}

		return
		// The following 4 endless loops are
		// equivalent to each other.
		for true {
		}
		for true {
		}
		for {
		}
		for {
		}
	}()

	func() {
		// SAMPLE D
		for i := 0; i < 3; i++ {
			fmt.Print(i)
			// The left i is a new declared variable,
			// and the right i is the loop variable.
			i := i
			// The new declared variable is modified, but
			// the old one (the loop variable) is not yet.
			i = 10
			_ = i
		}
	}()

	func() {
		// SAMPLE E
		// a `break` statement can be used to make execution jump out of a for loop
		// control flow block in advance
		i := 0
		for {
			if i >= 10 {
				break
			}
			fmt.Println(i)
			i++
		}
	}()

	func() {
		// SAMPLE F
		// a `continue` statement can be used to end the current loop iteration in
		// advance
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				continue
			}
			fmt.Print(i)
		}
		fmt.Println()
	}()

	func() {
		// SAMPLE G
		// for-range Control Flow Blocks
		for i := range 10 {
			if i%2 == 0 {
				continue
			}
			fmt.Print(i)
		}
	}()

	fmt.Println("vim-go")
}
