---
layout: default
title: Channel
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/channel
---

Other examples:
- [go channels — the ultimate guide]({% link docs/languages/golang/channel.2023a03m31d.md %}) [^1]
- [go channels — deadlock]({% link docs/languages/golang/channel.deadlock.md %})


Channels can be categorized as send-and-receive, read-only, and send-only according to whether the internal `recvq` and `sendq` queues are permitted to store the corresponding waiters, the `goroutines`.

The key is to keep in mind the core of using channels, the __signal transmission between goroutines__. [^2]

![image](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*a7botFtrJhid0KWeKB4AYg.png)

One of the main features of Go programming language is its eponymous `go` statement. In Go, it’s super-easy to start a new goroutine. But how do you get its results? How do you know if it may have errored? How do you wait for it to complete? How do you cancel it, if you don’t need its results anymore? Those familiar with Go will say: well, obviously, use channels. But channel in Go is still a low level construct. _The more channels you have — the more problems. Channels can deadlock. Channels can panic._ [^3]

Different languages and frameworks have a design pattern to defer value. Some call it __Future__, other call it __Promise__, and a bunch of other names. A `go` statement is eager by its nature. [^3]

# Channel

In Go, channels are used to communicate between goroutines. Channels are like pipes through which you can pass data. However, when we try to read from a channel that has no data, it can cause the program __to block__. _To avoid blocking_, Go provides a built-in mechanism that allows us to detect if a channel is closed or if there is no data available. This mechanism is called the “comma ok” syntax.

The “comma ok” syntax is a Go idiom that is used to check if a channel is closed or if there is data available to read from the channel. The syntax is as follows:
- `value, ok := <-channel`

```golang
import "fmt"

func main() {
  ch := make(chan int)

  go func() {
    fmt.Println("2")
    ch <- 42
    fmt.Println("3")
    ch <- 10
    fmt.Println("6")
    close(ch)
    fmt.Println("7")
  }()

  fmt.Println("1")
  value, ok := <-ch

  fmt.Println("4")
  fmt.Println(value, ok)

  value, ok = <-ch
  fmt.Println("5")
  fmt.Println(value, ok)

  value, ok = <-ch
  fmt.Println("8")
  fmt.Println(value, ok)

  // output:
  // 1
  // 2
  // 3
  // 4
  // 42 true
  // 5
  // 10 true
  // 6
  // 7
  // 8
  // 0 false
}
```

## When to Use Channel [^2]

- trigger signal, including end and start
- transfer data in async. Asynchronous worker processes one by one for the non-urgent
- block for purpose. For critical steps, you need to block and wait
- the worker pool. The worker is either woken up by the task or long blocked until the task comes

----

[^1]: [Go Channels — The Ultimate Guide](https://medium.com/@lordmoma/go-channels-the-ultimate-guide-3a2552a2a458)
[^2]: [When and How to Use the Go Channel](https://betterprogramming.pub/when-and-how-to-use-the-go-channel-58f82605156c)
[^3]: [Build your own Future in Go](https://levelup.gitconnected.com/build-your-own-future-in-go-f66c568e9a7a)
