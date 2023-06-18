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

A function to create a new file with the specified name in the current directory.

<details markdown="block">
  <summary>
    a full code snippet
  </summary>

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

</details>

# Open

A function to open the file with the specified name in the current directory.

<details markdown="block">
  <summary>
    a full code snippet
  </summary>

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

</details>

# Write

<details markdown="block">
  <summary>
    a full code snippet
  </summary>

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

</details>

# Read

It opens a file named "example.txt", read up to 100 bytes from the file using the Read function, and then print the number of bytes read and the read data.

<details markdown="block">
  <summary>
    a full code snippet
  </summary>

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

</details>

# Persisting a Go objects to disk

A simple code that lets save Go objects to disk, and read them back [^1].

<details markdown="block">
  <summary>
    a full code snippet
  </summary>

```golang
package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "log"
  "os"
  "time"
)

// Save saves a representation of v to the file at path.
func Save(path string, v interface{}) error {
  // Marshal is a function that marshals the object into an
  // io.Reader.
  // By default, it uses the JSON marshaller.
  var Marshal = func(v interface{}) (io.Reader, error) {
    b, err := json.MarshalIndent(v, "", "  ")
    // b, err := json.MarshalIndent(v, "", "\t")
    if err != nil {
      return nil, err
    }
    return bytes.NewReader(b), nil
  }

  f, err := os.Create(path)
  if err != nil {
    return err
  }
  defer f.Close()
  r, err := Marshal(v)
  if err != nil {
    return err
  }
  _, err = io.Copy(f, r)
  return err
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func Load(path string, v interface{}) error {
  // Unmarshal is a function that unmarshals the data from the
  // reader into the specified value.
  // By default, it uses the JSON unmarshaller.
  var Unmarshal = func(r io.Reader, v interface{}) error {
    return json.NewDecoder(r).Decode(v)
  }

  f, err := os.Open(path)
  if err != nil {
    return err
  }
  defer f.Close()
  return Unmarshal(f, v)
}

func main() {
  type obj struct {
    Name   string
    Number int
    When   time.Time
  }

  o := &obj{
    Name:   "Mat",
    Number: 47,
    When:   time.Now(),
  }
  if err := Save("./file.tmp", o); err != nil {
    log.Fatalln(err)
  }
  // load it back
  var o2 obj
  if err := Load("./file.tmp", &o2); err != nil {
    log.Fatalln(err)
  }
  fmt.Println(o2)
  // o and o2 are now the same
  // and check out file.tmp - you'll see the JSON file
  fmt.Println("vim-go")
}
```

<br/>
</details>
<br/>

[^1]: [Persisting Go objects to disk](https://medium.com/@matryer/golang-advent-calendar-day-eleven-persisting-go-objects-to-disk-7caf1ee3d11d)
