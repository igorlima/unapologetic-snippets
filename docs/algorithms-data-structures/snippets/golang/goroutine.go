package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // sleep for 0 to 2.5 seconds
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	time.Sleep(2 * time.Second)
	/*
		When the main goroutine exits, the whole program also exits,
		even if there are still some other goroutines which have not exited yet.
	*/
	fmt.Println("vim-go")
}
