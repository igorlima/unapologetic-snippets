package main

import "fmt"

func main() {
	/*
		String Concatenation Operator
	*/
	println("Go" + "lang") // Golang
	var a = "Go"
	a += "lang"
	println(a) // Golang

	/*
		fmt.Printf Format Verbs
	*/
	a1, b := 123, "Go"
	fmt.Printf("a1 == %v == 0x%x, b == %s\n", a1, a1, b)
	fmt.Printf("type of a1: %T, type of b: %T\n", a1, b)
	fmt.Printf("1%% 50%% 99%%\n")

	fmt.Println("vim-go")
}
