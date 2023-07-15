---
layout: default
title: Concurrency Context
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/concurrency-context-2023a07m02d
---

__[back]({% link docs/languages/golang/concurrency.md %}#context)__

The problems `context` plans to solve are: [^1]

- Let’s say that you started a function and you need __to pass some common parameters to the downstream functions__. You cannot pass these common parameters each as an argument to all the downstream functions.
- You started a goroutine which in turn start more goroutines and so on. Suppose __the task__ that you were doing __is no longer needed__. Then how to inform all child goroutines to gracefully exit so that resources can be freed up.
- A task should be finished __within a specified timeout__ of say 2 seconds. If not it should gracefully exit or return.
- A task should be __finished within a deadline__ eg it should end before 5 pm . If not finished then it should gracefully exit and return

The most basic data structure of context is the Context interface as below:

```golang
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}

// Context is just an interface, which is very hard to imagine how to use it.
```

<details markdown="block">
  <summary>
    WithCancel
  </summary>

```golang
// Context With Cancel
package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  ctxA, cancelFuncA := context.WithCancel(context.Background())
  ctxB, cancelFuncB := context.WithCancel(context.WithValue(ctxA, "foo", "bar"))

  go stopTheTaskIfNotInterested(ctxA)
  go stopTheTaskIfNotInterested(ctxB)

  time.Sleep(3 * time.Second)
  fmt.Println("DONE")

  // cancelling the context
  cancelFuncA()
  cancelFuncB()
  time.Sleep(2 * time.Second)
}

func stopTheTaskIfNotInterested(ctx context.Context) {
  select {
  case <-ctx.Done():
    fmt.Println("exiting the task... because parent not interested..!!", ctx.Value("foo"))
    return
  }
}

/*
OUTPUT

DONE
exiting the task... because parent not interested..!! <nil>
exiting the task... because parent not interested..!! bar
*/
```

<br/>
<!-- WithCancel 1 -->
</details>

<details markdown="block">
  <summary>
    WithCancel <sup><i>plus <code>sync.WaitGroup</code>, <code>context.WithValue</code> </i></sup>
  </summary>

```golang
// Context With Cancel
package main

import (
  "context"
  "fmt"
  "sync"
)

func main() {
  var wg sync.WaitGroup
  ctx, cancelFunc := context.WithCancel(context.WithValue(context.Background(), "foo", "bar"))

  wg.Add(1)
  fmt.Println("1")
  go stopTheTaskIfNotInterested(ctx, &wg)
  fmt.Println("2")
  cancelFunc()

  fmt.Println("3")
  wg.Wait()
  fmt.Println("5")
  fmt.Println("DONE")

  select {
  case <-ctx.Done():
    fmt.Println("6")
    // the context is over, stop processing results
    fmt.Println("all done", ctx.Value("foo"))
  }

  fmt.Println("7")
}

func stopTheTaskIfNotInterested(ctx context.Context, wg *sync.WaitGroup) {
  select {
  case <-ctx.Done():
    fmt.Println("4")
    fmt.Println("exiting the task... because parent not interested..!!", ctx.Value("foo"))
    wg.Done()
    return
  }
}

/*
OUTPUT

1
2
3
4
exiting the task... because parent not interested..!! bar
5
DONE
6
all done bar
7
*/
```

<br/>
<!-- WithCancel 2 -->
<!-- WithCancel plus sync.WaitGroup, context.WithValue -->
</details>

<details markdown="block">
  <summary>
    WithTimeout
  </summary>

```golang
// Context With Timeout
package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  aux := context.Background()
  ctx, cancelFunc := context.WithTimeout(aux, time.Second*1)

  go stopTheTaskIfNotCompletedWithInGivenTime(ctx)

  time.Sleep(3 * time.Second)
  fmt.Println("...")
  time.Sleep(2 * time.Second)

  // cancels the context if timeout was not reached yet
  defer func() {
    cancelFunc()
    fmt.Println("Done")
  }()
  fmt.Println("DONE")
  time.Sleep(2 * time.Second)
}

func stopTheTaskIfNotCompletedWithInGivenTime(ctx context.Context) {
  select {
  case <-ctx.Done():
    fmt.Println("exiting the task. because i am not able to complete with in time..!!!")
    return
  }
}

/*
OUTPUT

exiting the task. because i am not able to complete with in time..!!!
...
DONE
Done
*/
```

<br/>
<!-- WithTimeout -->
</details>


<details markdown="block">
  <summary>
    WithDeadline
  </summary>

```golang
package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))

  go stopTheTaskIfNotCompletedWithInGivenTime(ctx)

  time.Sleep(3 * time.Second)
  fmt.Println("...")
  time.Sleep(2 * time.Second)

  // cancels the context if deadline was not reached yet
  defer func() {
    cancelFunc()
    fmt.Println("Done")
  }()
  fmt.Println("DONE")
  time.Sleep(2 * time.Second)
}

func stopTheTaskIfNotCompletedWithInGivenTime(ctx context.Context) {
  select {
  case <-ctx.Done():
    fmt.Println("exiting the task. because i am not able to complete with in time... dealine exceeded!!!")
    return
  }
}

/*
OUTPUT

exiting the task. because i am not able to complete with in time... dealine exceeded!!!
...
DONE
Done
*/
```

<br/>
<!-- WithDeadline -->
</details>
<br/>

- `context.Background()`: This function returns a background context, which serves as the root parent context. It is often used as the starting point for creating other contexts.
- `context.WithValue(parentContext, key, value)`: This function creates a child context derived from a parent context (`parentContext`). It associates a key-value pair with the child context, allowing for the passing of request-scoped values. The child context inherits the values from the parent context and can add or overwrite values specific to itself.

__Context propogate from parent to child.__
When a parent goroutine creates a child goroutine in Go, the context can be propagated from the parent to the child. Context propagation allows the child goroutine to inherit and carry forward the same context values, deadlines, and cancellations.
In Go, context propagation is achieved through the use of the `context.Context` type, which is passed as an argument to functions or goroutines that need access to the context.


----

[^1]: [Understand Golang Context by source code](https://medium.com/gitconnected/understand-golang-context-by-source-code-4863a851519)
