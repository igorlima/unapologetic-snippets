package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		The arguments of a goroutine function call are all evaluated at the moment
		when the corresponding goroutine function statement is executed.
	*/

	var a = 12
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Println(x, a) // 12 20
	}(a)

	a = 20

	time.Sleep(2 * time.Second)
	fmt.Println("vim-go")
}
