---
layout: default
title: Concurrency from Utkarsh Jha
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/concurrency-2023a04m02d
---

__[back]({% link docs/languages/golang/concurrency.md %})__

Examples from a medium post [^1]

# Channel

The code below creates a simple worker pool using goroutines and channels for communication. The `worker` function processes jobs received from the `jobs` channel and sends `results` to the results channel. The main function starts three workers and sends nine jobs for processing. The program demonstrates synchronization between goroutines using channels, allowing for controlled data flow and concurrent execution. [^2]

```golang
package main

import (
  "fmt"
  "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
  for j := range jobs {
    fmt.Println("worker", id, "started job", j)
    time.Sleep(time.Second)
    fmt.Println("worker", id, "finished job", j)
    results <- j * 2
  }
}

func main() {
  jobs := make(chan int, 100)
  results := make(chan int, 100)
  for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
  }

  for j := 1; j <= 9; j++ {
    jobs <- j
  }
  close(jobs)

  for a := 1; a <= 9; a++ {
    fmt.Println("results", <-results)
  }

  fmt.Println("vim-go")
}

/*
OUTPUT

worker 3 started job 1
worker 1 started job 2
worker 2 started job 3
worker 2 finished job 3
worker 2 started job 4
worker 1 finished job 2
worker 1 started job 5
worker 3 finished job 1
worker 3 started job 6
results 6
results 4
results 2
worker 2 finished job 4
worker 2 started job 7
worker 1 finished job 5
worker 1 started job 8
worker 3 finished job 6
worker 3 started job 9
results 8
results 10
results 12
worker 3 finished job 9
results 18
worker 1 finished job 8
results 16
worker 2 finished job 7
results 14
vim-go
*/
```

# Statement

```golang
package main

import (
  "fmt"
  "time"
)

func worker1(c chan string) {
  time.Sleep(time.Second * 2)
  c <- "worker 1"
}

func worker2(c chan string) {
  time.Sleep(time.Second * 1)
  c <- "worker 2"
}

func main() {
  c1 := make(chan string)
  c2 := make(chan string)

  go worker1(c1)
  go worker2(c2)

  select {
  case res := <-c1:
    fmt.Println(res)
  case res := <-c2:
    fmt.Println(res)
  }
  fmt.Println("vim-go")
}

/*
OUTPUT

worker 2
vim-go
*/
```

# Mutex

```golang
package main

import (
  "fmt"
  "sync"
  "time"
)

type SafeCounter struct {
  counter int
  mutex   sync.Mutex
}

func (s *SafeCounter) Inc() {
  s.mutex.Lock()
  s.counter++
  s.mutex.Unlock()
}

func (s *SafeCounter) Value() int {
  s.mutex.Lock()
  defer s.mutex.Unlock()
  return s.counter
}

func main() {
  counter := SafeCounter{}

  for i := 0; i < 1000; i++ {
    go counter.Inc()
  }

  time.Sleep(time.Second)
  fmt.Println(counter.Value())
  fmt.Println("vim-go")
}

/*
OUTPUT

1000
vim-go
*/
```

----

[^1]: [Mastering Advanced Go Concurrency: Powerful Concepts and Examples](https://medium.com/@utkarshjha_81903/mastering-advanced-go-concurrency-powerful-concepts-and-examples-2011d8b01855)
[^2]: [Golang Goroutines: Powering High-Performance Applications](https://medium.com/@cerebrovinny/golang-goroutines-powering-high-performance-applications-767742d961ce)
