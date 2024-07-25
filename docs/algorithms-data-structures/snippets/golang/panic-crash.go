package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	 Example which shows a panicking goroutine exits without being recovered.  So
	 the whole program crashes.
	*/

	fmt.Println("hi!")

	go func() {
		time.Sleep(time.Second)
		panic(123)
	}()

	for {
		time.Sleep(time.Second)
	}

	fmt.Println("vim-go")
}
