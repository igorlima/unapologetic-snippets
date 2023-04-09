---
layout: default
title: Channel
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/channel-2023a03m31d
---

__[back]({% link docs/languages/golang/channel.md %})__

# Go Channel

Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other. [^1]

Channels send and receive values with the channel operator:

```golang
// the data flows in the direction of the arrow.
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```

Channels must be created __before__ use. They use the `make` keyword:

```golang
ch := make(chan int)
```

By default, sends and receives __block__ (which means your code safely waits) until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

## Deadlock

A deadlock is when a group of goroutines are blocking, so none of them can continue. This is a common problem in concurrent programming.

```golang
// a send to a nil channel blocks forever
func () {
  var c chan string
  c <- "let's get started" // deadlock
}()
```

```golang
// a receive from a nil channel blocks forever
func () {
  var c chan string
  fmt.Println(<-c) // deadlock
}()
```


## Channels Continued

Notice that we can block and wait until something is sent on a channel using the following syntax.

```golang
// this will block until it pops a single item off the channel, then continue, discarding the item.
<-ch
```

```golang
package main

import (
  "fmt"
  "time"
)

func main() {
  ch := getDatabasesChannel()

  <-ch
  <-ch
  <-ch
  <-ch

  time.Sleep(1 * time.Second)
  fmt.Println("mailio server ready")
  fmt.Println("vim-go")
}

func getDatabasesChannel() chan struct{} {
  ch := make(chan struct{})
  go func() {
    ch <- struct{}{}
    fmt.Println("first db online")
    ch <- struct{}{}
    fmt.Println("second db online")
    ch <- struct{}{}
    fmt.Println("third db online")
    ch <- struct{}{}
    fmt.Println("fourth db online")
  }()
  return ch
}

/*
OUTPUT

first db online
second db online
third db online
fourth db online
mailio server ready
vim-go
*/
```

## Buffered Channels

Channels can optionally be buffered. You may provide a buffer length as the second argument to `make()` to create a buffered channel:

```golang
// sending on a buffered channel only blocks when the buffer is full.
ch := make(chan int, 100)
```

```golang
package main

import (
  "fmt"
)

func main() {
  emailsToSend := make(chan string, 2)
  emailsToSend <- "Hello John, tell Kathy I said hi"
  emailsToSend <- "Whazzup bruther"
  fmt.Println(<-emailsToSend)
  fmt.Println(<-emailsToSend)
}

/*
OUTPUT
Hello John, tell Kathy I said hi
Whazzup bruther
*/
```

## Closing channels in Go

```golang
close(ch)

// receivers can check the ok value when receiving from a channel to test if a channel was closed.
v, ok := <-ch
```

`ok` is `false` if the channel is empty and closed.

Sending on a closed channel will cause a panic.


## Range

Similar to slices and maps, channels can be ranged over.

```golang
// the below will receive values over the channel (blocking at each iteration if nothing new is there) and will exit only when the channel is closed.
for item := range ch{
}
```

```golang
package main

import (
  "fmt"
)

func main() {
  ch := make(chan int, 10)
  go fibonacci(cap(ch), ch)
  for num := range ch {
    fmt.Println(num)
  }
}
func fibonacci(n int, ch chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    ch <- x
    x, y = y, x+y
  }
  close(ch)
}

/*
OUTPUT
0
1
1
2
3
5
8
13
21
34
*/
```

## Select

Sometimes we have a goroutine listening to multiple channels and want to process data in the order it comes through, no matter which channel sends the data first.

```golang
select {
  case i, ok := <- chInts:
    fmt.Println(i)
  case s, ok := <- chStrings:
    fmt.Println(s)
}
```

```golang
// Select Default Case
select {
  case i := <-cH:
    // use i
  default:
    // receiving from ch would block
}
```

## time.Ticker

- `time.Tick()` is a standard library function that returns a channel that sends a value on a given interval.
- `time.After()` sends a value once after the duration has passed.
- `time.Sleep()` blocks the current goroutine for the specified amount of time.

```golang
package main

import (
  "fmt"
  "time"
)

func main() {
  snapshotInterval := time.Tick(800 * time.Millisecond)
  saveAfter := time.After(2800 * time.Millisecond)
  for {
    select {
    case <-snapshotInterval:
      fmt.Println("taking backup snapshot...")
    case <-saveAfter:
      fmt.Println("all backups saved!")
      return
    default:
      fmt.Println("nothing to do, waiting...")
      time.Sleep(500 * time.Millisecond)
    }
  }
}

/*
OUTPUT
nothing to do, waiting...
nothing to do, waiting...
taking backup snapshot...
nothing to do, waiting...
nothing to do, waiting...
taking backup snapshot...
nothing to do, waiting...
taking backup snapshot...
nothing to do, waiting...
all backups saved!
*/
```

----

[^1]: [Go Channels — The Ultimate Guide](https://medium.com/@lordmoma/go-channels-the-ultimate-guide-3a2552a2a458)
