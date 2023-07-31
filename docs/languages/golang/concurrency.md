---
layout: default
title: Concurrency
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/concurrency
---

Other examples:
- [more samples for channel, statement, and mutex]({% link docs/languages/golang/concurrency.2023a04m02d.md %}) [^1]
- [golang goroutines: powering high-performance applications]({% link docs/languages/golang/concurrency.goroutines.2023a04m02d.md %}) [^3]
- [actor model for concurrent systems]({% link docs/languages/golang/concurrency.actor-model.2023a04m02d.md %}) [^2]
- [understanding `context`]({% link docs/languages/golang/concurrency.context.2023a07m02d.md %})

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

# Goroutine

A goroutine is a lightweight thread that runs concurrently with other goroutines in the same address space.

<details markdown="block">
  <summary>
    sample one
  </summary>

```golang
package main

import (
  "fmt"
  "time"
)

func printMessage(message string, ch chan string, sec int) {
  time.Sleep(time.Duration(sec) * time.Second)
  ch <- message
}

func main() {
  ch := make(chan string)
  fmt.Println("1")
  go printMessage("Hello, Go!", ch, 2)
  go printMessage("Welcome to Concurrency!", ch, 1)

  fmt.Println("2")
  msg1 := <-ch
  fmt.Println(msg1)

  msg2 := <-ch
  fmt.Println(msg2)
  fmt.Println("3")
}
```

<br/>
</details>

<details markdown="block">
  <summary>
    sample two
  </summary>

```golang
package main

import (
  "fmt"
  "time"
)

func main() {
  go printMessage("Hello", 5)
  go printMessage("world", 5)
  time.Sleep(1 * time.Second)
  fmt.Println("goroutine")
}

func printMessage(msg string, count int) {
  for i := 0; i < count; i++ {
    fmt.Println(msg)
    time.Sleep(100 * time.Millisecond)
  }
}

/*
OUTPUT

world
Hello
Hello
world
Hello
world
Hello
world
world
Hello
goroutine
*/
```

<br/>
</details>

# Channels

A channel is a typed conduit thorugh which you can send and receive values with the `<-` operator. Channels ensure safe and efficient communication between concurrent processes.

```golang
package main

import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan string)
  go func() {
    ch <- "Hello from channel!"
    time.Sleep(2 * time.Second)
    ch <- "Hi from channel!"
    close(ch)
  }()

  msg, ok := "", true
  for ok {
    msg, ok = <-ch
    fmt.Println(msg)
  }
}

/*
OUTPUT

Hello from channel!
Hi from channel!
*/
```

## Buffered Channels

Buffered channels are channels that can hold a certain number of values before they are read.
Buffered channels are created using the `make` function with a second argument specfying the buffer size.

```golang
package main

import (
  "fmt"
)

func main() {
  ch := make(chan string, 1)
  go func() {
    ch <- "Hello from channel!"
    fmt.Println(1)
    ch <- "Hi from channel!"
    fmt.Println(2)
    ch <- "Hi-five from channel!"
    fmt.Println(3)
    ch <- "Hi there from channel!"
    fmt.Println(4)
    close(ch)
  }()

  msg, ok := "", true
  for ok {
    msg, ok = <-ch
    fmt.Println("---")
    fmt.Println(msg)
  }
}

/*
OUTPUT

1
2
---
Hello from channel!
---
Hi from channel!
---
Hi-five from channel!
3
4
---
Hi there from channel!
---
*/
```

# Select Statement

The select statement in Go allows you to wait on multiple channel operations simultaneously.
The select statement blocks until one of its cases can proceed, at which point it executes that case.

```golang
package main

import (
  "fmt"
  "time"
)

func main() {
  ch1 := make(chan string)
  ch2 := make(chan string)

  go func() {
    time.Sleep(time.Second)
    ch1 <- "Hello from channel 1!"
    time.Sleep(time.Second)
    ch1 <- "Hello twice from channel 1!"
  }()

  go func() {
    time.Sleep(time.Second)
    ch2 <- "Hello from channel 2!"
  }()

  select {
  case msg1 := <-ch1:
    fmt.Println(msg1)
  case msg2 := <-ch2:
    fmt.Println(msg2)
  }
}

/*
OUTPUT

Hello from channel 1!
*/
```

# Mutex and WaitGroup

A mutex is a mutual exclusion lock that allows only one goroutine to access a resource at a time. Any other goroutines that attempt to acces the resource while it's locked will be blocked until the lock is released.

A WaitGroup waits for a collection of goroutines to finish before continuing.

A Wait Group is essentially a counter that __keeps track of the number of goroutines that are active or pending completion__. It allows the main goroutine to wait until all other goroutines have finished executing.

```golang
package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  var wg sync.WaitGroup
  c := Counter{count: 0}
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(n int) {
      defer wg.Done()
      time.Sleep(time.Duration(n) * time.Second)
      c.Increment()
      fmt.Println("Goroutine", n, "finished")
    }(i)
  }
  wg.Wait()
  fmt.Println(c.Count())
  fmt.Println("All goroutines finished")
}

type Counter struct {
  count int
  mutex sync.Mutex
}

func (c *Counter) Increment() {
  c.mutex.Lock()
  defer c.mutex.Unlock()
  c.count++
}

func (c *Counter) Count() int {
  c.mutex.Lock()
  defer c.mutex.Unlock()
  return c.count
}

/*
OUTPUT

Hello from channel 1!
*/
```

# Context

In Go, the `context` package provides a powerful mechanism for managing the lifecycle of operations and __propagating cancellation signals across goroutines__. It allows to set deadlines, timeouts, and cancelation signals, helping to control the execution of the code and handle resource cleanup gracefully. There are three commonly used functions in the context package for managing cancellation: `WithCancel`, `WithTimeout`, and `WithDeadline`. [^4]
- [understanding `context`]({% link docs/languages/golang/concurrency.context.2023a07m02d.md %})

__What is the purpose of the `context` package in Go?__
The `context` package provides a way to pass cancellation signals, deadlines, and other request-scoped values across API boundaries.
It is used to manage the lifecycle of operations and facilitate cancellation and timeout handling.

__What is `context` and why it is needed?__
In Go, a context is a mechanism for managing concurrent operations, such as goroutines, by passing information and signals between different parts of a program. It allows for handling timeouts, cancellations, and sharing values in a controlled manner.


----

[^1]: [Mastering Advanced Go Concurrency: Powerful Concepts and Examples](https://medium.com/@utkarshjha_81903/mastering-advanced-go-concurrency-powerful-concepts-and-examples-2011d8b01855)
[^2]: [Actor model for concurrent systems, an introduction in GO](https://medium.com/@joao_vaz/actor-model-for-concurrent-systems-an-introduction-in-go-75fd25f2f04e)
[^3]: [Golang Goroutines: Powering High-Performance Applications](https://medium.com/@cerebrovinny/golang-goroutines-powering-high-performance-applications-767742d961ce)
[^4]: [Context in GoLang _by Sai Teja_](https://medium.com/@saiteja180/context-in-golang-e8cde4fba1e2)
