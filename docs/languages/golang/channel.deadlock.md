---
layout: default
title: Channel Deadlock
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/channel-deadlock
---

# Go Channel - Deadlock
__[back]({% link docs/languages/golang/channel.md %})__

- _See more:_
  - [go channel — deadlock]({% link docs/languages/golang/channel.2023a03m31d.md %}#deadlock)

_What happens when you try to send data to a full channel or receive data from an empty channel?_

![image](https://user-images.githubusercontent.com/1886786/232350393-b08733cd-eff2-4785-ab81-459139aaf804.png)


When you try to send data to a full channel or receive data from an __empty__ channel, __the operation will block until the channel has the available capacity__ or until there is data to be received. [^1]

If a channel is __full__ and you try to send data to it, __the sending operation will block until there is available capacity in the channel__. This means that the sending goroutine will wait until a receiver removes an element from the channel to make space for the new element. [^1]

For example, consider a buffered channel that can hold up to 2 integers:

```golang
package main

import "fmt"

func main() {
  ch := make(chan int, 2)
  ch <- 1
  ch <- 2
  ch <- 3 // this will block until space is available in the channel
  fmt.Println("vim-go")
}
```

```sh
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/me/workstation/golang/2023a04m16d-go-doc/main.go:9 +0x5c
exit status 2
```

Similarly, if a channel is __empty__ and you try to receive data from it, __the receiving operation will block until there is data to be received__. This means that the receiving goroutine will wait until a sender sends data on the channel.

For example, consider a channel that is not buffered and can hold only one integer:

```golang
package main

import "fmt"

func main() {
  ch := make(chan int)
  value := <-ch // this will block until a value is sent on the channel
  fmt.Println(value)
  fmt.Println("vim-go")
}
```

```sh
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/me/workstation/golang/2023a04m16d-go-doc/main.go:7 +0x39
exit status 2
```

In this example, the receiving operation will block because there is no data to be received on the channel. The receiving goroutine will wait until a sender sends data on the channel.

## Channels Types?

__Buffered channels__: This type of channel has a fixed capacity to hold a certain number of values.

```golang
package main
import "fmt"
func main() {
  ch := make(chan int, 1)
  ch <- 10
  fmt.Println("Value sent")
  fmt.Println("Receiving value")
  fmt.Println(<-ch)
  fmt.Println("Value received")
}
```
```
Value sent
Receiving value
10
Value received
```

__Unbuffered channels__: An unbuffered channel has no capacity to hold any values. When a value is sent on an unbuffered channel, the sender will block it until the value is received by a receiver.

```golang
package main
import "fmt"
func main() {
  ch := make(chan int)
  go func() {
    fmt.Println("Sending value")
    ch <- 10
    fmt.Println("Value sent")
  }()
  fmt.Println("Receiving value")
  fmt.Println(<-ch)
  fmt.Println("Value received")
}
```

```
Receiving value
Sending value
Value sent
10
Value received
```


----

[^1]: [Mastering Go Channels- Part 2](https://medium.com/@varmapooja09/mastering-go-channels-part-2-ebc6f0e8c8f5)
