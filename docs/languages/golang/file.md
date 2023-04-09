---
layout: default
title: File
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/file
---

File handling in Go is done using the `os` package. The `os` package provides a set of functions for creating, opening, reading, writing, and closing files.

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

# Create

The Create function creates a new file with the specified name in the current directory.

```golang
package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.Create("example.txt")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()
  fmt.Println("vim-go")
}
```

# Open

The Open function opens the file with the specified name in the current directory.

```golang
package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("example.txt")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()
  fmt.Println("vim-go")
}
```

# Write

```golang
package main

import (
  "fmt"
  "os"
)

func main() {
  data := "Hello, World!"

  // os.Open has the associated file descriptor as mode O_RDONLY
  // https://pkg.go.dev/os#Open
  // file, err := os.Open("example.txt")

  // file, err := os.Create("example.txt")
  file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()

  _, err = file.WriteString(data)
  if err != nil {
    fmt.Println("hello")
    fmt.Println(err)
    return
  }

  fmt.Println("vim-go")
}
```

# Read

It opens a file named "example.txt", read up to 100 bytes from the file using the Read function, and then print the number of bytes read and the read data.

```golang
package main

import (
  "fmt"
  "os"
)

func main() {
  // os.Open has the associated file descriptor as mode O_RDONLY
  // https://pkg.go.dev/os#Open
  file, err := os.Open("example.txt")

  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()

  data := make([]byte, 100)
  count, err := file.Read(data)

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("read %d bytes: %q\n", count, data[:count])
  fmt.Println("vim-go")
}
```
