---
layout: default
title: Concurrency Goroutines
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/concurrency-goroutines-2023a04m02d
---

__[back]({% link docs/languages/golang/concurrency.md %})__

One of the key features that sets Go apart from other languages is its support for concurrency through goroutines — a way of creating lightweight threads that can run simultaneously. [^1]

Known for its simplicity and ease-of-use, Go has quickly become a go-to for developers looking to create concurrent programs.

Golang concurrency is no stranger to the world of programming, but what makes it such a hot topic? With the rise of multithreading and parallel processing, mastering concurrency in Golang has become more relevant than ever.

The lightweight threads provide numerous advantages over traditional threads, such as lower memory overhead and better scalability.

Goroutines are lightweight threads that run concurrently within the same address space.

Synchronization with channels in Golang refers to controlling the flow of data between goroutines by establishing channels or data pipelines.

# Channel

Using channels for synchronization and data exchange

Channels in Golang allow goroutines to synchronize and exchange data, providing a simple and safe way to coordinate concurrent work.

In this example, a buffered channel `done` signals when the `doWork` function has completed execution. The `main` function waits for both workers to finish before exiting.

```golang
package main

import (
  "fmt"
  "time"
)

func main() {

  done := make(chan bool, 2)
  go doWork(1, done)
  go doWork(2, done)
  <-done
  <-done
  fmt.Println("vim-go")
}

func doWork(id int, done chan bool) {
  fmt.Printf("Worker %d started\n", id)
  time.Sleep(time.Second)
  fmt.Printf("Worker %d finished\n", id)
  done <- true
}
```

# Context

Using context.Context for cancelation and deadlines

The `context.Context` package can be used to handle cancelation and deadlines in concurrent programs.

The `context.Context` manages cancelation for the `doWorkWithContext` function.

```golang
package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  go doWorkWithContext(ctx, 1)
  go doWorkWithContext(ctx, 2)

  time.Sleep(2000 * time.Millisecond)
  cancel()
  time.Sleep(2 * time.Second)
  fmt.Println("vim-go")
}

func doWorkWithContext(ctx context.Context, id int) {
  select {
  case <-time.After(2 * time.Duration(id) * time.Second):
    fmt.Printf("Worker %d completed\n", id)
  case <-ctx.Done():
    fmt.Printf("Worker %d canceled\n", id)
  }
}
```

# Mutex

Using Mutexes to protect shared data

Mutexes can be used to protect shared data from concurrent access, preventing race conditions.

This example demonstrates using a mutex to protect the `value` field in the `Counter` struct. The `Increment` method locks the mutex before modifying the shared data and unlocks it afterward.

```golang
package main

import (
  "fmt"
  "sync"
)

func main() {
  counter := &Counter{}
  // without the WaitGroup, it ends without counting till 1000
  var wg sync.WaitGroup

  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
      counter.Increment()
      wg.Done()
    }()
  }

  wg.Wait()
  fmt.Printf("Counter value: %d\n", counter.value)
  fmt.Println("vim-go")
}

type Counter struct {
  value int
  mu    sync.Mutex
}

func (c *Counter) Increment() {
  c.mu.Lock()
  defer c.mu.Unlock()
  c.value++
}
```

# Select

Using select for multiple channel operations

The `select` statement allows a goroutine to perform multiple channel operations, choosing __the first one that is ready__.

In this example, the `select` statement waits for either a timeout or a tick. The main function exits when the timeout is reached, and ticks are printed every 200 milliseconds until then.

```golang
package main

import (
  "fmt"
  "time"
)

func main() {
  timeout := time.After(1 * time.Second)
  tick := time.Tick(200 * time.Millisecond)

  for {
    select {
    case <-timeout:
      fmt.Println("Timeout reached")
      return
    case <-tick:
      fmt.Println("Tick")
    }
  }
  fmt.Println("vim-go")
}
```

----

[^1]: [Golang Goroutines: Powering High-Performance Applications](https://medium.com/@cerebrovinny/golang-goroutines-powering-high-performance-applications-767742d961ce)
