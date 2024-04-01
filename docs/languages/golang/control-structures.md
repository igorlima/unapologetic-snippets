---
layout: default
title: Control Structures
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/control-structures
---

<br/>
<details markdown="block">
  <summary> Loop </summary>

```golang
package main

import (
  "fmt"
  "sync"
  "time"
)

// Golang for Loop
// Concurrency with Goroutines
func main() {
  fmt.Println("vim-go")

  // example of using goroutines with for loop
  for i := 0; i < 3; i++ {
    go func(number int) {
      fmt.Println("Goroutine A:", number)
    }(i)
  }
  time.Sleep(2 * time.Second)

  // example of using WaitGroup with goroutines and for loop
  var wg sync.WaitGroup
  for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(number int) {
      defer wg.Done()
      fmt.Println("GoroutineB:", number)
    }(i)
  }
  wg.Wait()

  // example of using channels with fo loop
  messages := make(chan string)
  go func() {
    for i := 0; i < 3; i++ {
      messages <- fmt.Sprintf("channel message %d", i)
    }
    close(messages)
  }()
  for msg := range messages {
    fmt.Println(msg)
  }

  // example of using select with for loop
  channel1 := make(chan string)
  channel2 := make(chan string)
  go func() {
    for {
      time.Sleep(time.Second)
      channel1 <- "channel 1"
    }
  }()
  go func() {
    for {
      time.Sleep(2 * time.Second)
      channel2 <- "channel 2"
    }
  }()
  for {
    select {
    case msg1 := <-channel1:
      fmt.Println(msg1)
    case msg2 := <-channel2:
      fmt.Println(msg2)
    }
  }

  /*
    OUTPUT
    Goroutine A: 2
    Goroutine A: 0
    Goroutine A: 1
    GoroutineB: 2
    GoroutineB: 1
    GoroutineB: 0
    channel message 0
    channel message 1
    channel message 2
    channel 1
    channel 2
    channel 1
    channel 1
    channel 1
    channel 2
    channel 1
    channel 2
    channel 1
    channel 1
  */
}
```

-----
<!-- Loop -->
</details>

