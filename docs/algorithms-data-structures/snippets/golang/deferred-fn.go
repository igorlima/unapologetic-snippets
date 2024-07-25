package main

import "fmt"

func main() {
	defer fmt.Println("The third line.")
	defer fmt.Println("The second line.")
	fmt.Println("The first line.")

	/*
		The output:
		 The first line.
		 The second line.
		 The third line.
	*/

	fmt.Println("vim-go")
}
