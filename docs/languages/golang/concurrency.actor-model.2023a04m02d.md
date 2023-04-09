---
layout: default
title: Actor model concurrency
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/concurrency-actor-model-2023a04m02d
---

__[back]({% link docs/languages/golang/concurrency.md %})__

Examples from a medium post [^1]

```golang
package main

import (
  "fmt"
)

func main() {
  actor := NewActor()
  go actor.Receive()

  actor.Send("Hello, world!")
  actor.Send("This is a message.")
  // time.Sleep(time.Second)
  fmt.Println("vim-go")
}

type Actor struct {
  inbox chan string
}

func NewActor() *Actor {
  return &Actor{
    inbox: make(chan string),
  }
}

func (a *Actor) Receive() {
  for msg := range a.inbox {
    fmt.Println(msg)
  }
}

func (a *Actor) Send(msg string) {
  a.inbox <- msg
}
```

----

[^1]: [Actor model for concurrent systems, an introduction in GO](https://medium.com/@joao_vaz/actor-model-for-concurrent-systems-an-introduction-in-go-75fd25f2f04e)
