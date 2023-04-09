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


# Goroutine

A goroutine is a lightweight thread that runs concurrently with other goroutines in the same address space.

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

----

[^1]: [Mastering Advanced Go Concurrency: Powerful Concepts and Examples](https://medium.com/@utkarshjha_81903/mastering-advanced-go-concurrency-powerful-concepts-and-examples-2011d8b01855)
[^2]: [Actor model for concurrent systems, an introduction in GO](https://medium.com/@joao_vaz/actor-model-for-concurrent-systems-an-introduction-in-go-75fd25f2f04e)
[^3]: [Golang Goroutines: Powering High-Performance Applications](https://medium.com/@cerebrovinny/golang-goroutines-powering-high-performance-applications-767742d961ce)
