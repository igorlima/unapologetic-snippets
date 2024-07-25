package main

import "fmt"

func main() {
	/*
		The arguments of a deferred function call are all evaluated at the moment when
		the corresponding defer statement is executed (a.k.a. when the deferred call is
		pushed into the deferred call queue). The evaluation results are used when the
		deferred call is executed later during the existing phase of the surrounding
		call (the caller of the deferred call).
	*/

	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i+x)
		}
		x = 10
	}()

	fmt.Println()

	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i+x)
			}()
		}
		x = 20
	}()
	fmt.Println("vim-go")
}
