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

----

[^1]: [Go Channels — The Ultimate Guide](https://medium.com/@lordmoma/go-channels-the-ultimate-guide-3a2552a2a458)
