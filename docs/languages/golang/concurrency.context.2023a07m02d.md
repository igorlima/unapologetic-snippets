---
layout: default
title: Concurrency Context
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/concurrency-context-2023a07m02d
---

__[back]({% link docs/languages/golang/concurrency.md %})__

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

----

[^1]: [Understand Golang Context by source code](https://medium.com/gitconnected/understand-golang-context-by-source-code-4863a851519)
