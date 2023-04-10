---
layout: default
title: Pointer
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/pointer-2023a03m05d
---

__[back]({% link docs/languages/golang/pointer.md %})__

# Go Pointers

A pointer holds a memory address for a value. `*` operator is used to denote a variable as a pointer of the type, also used in pointer dereferencing, and `&` operator is used to getting the underlying value.

A basic example is:

```golang
package main

import "fmt"

func main() {
  str := "Hello Cosmos"
 
  // Pointer to string
  var pStr *string
 
  // pStr points to addr of message
  pStr = &str
  fmt.Println("Message = ", *pStr)
  fmt.Println("Message Address = ", pStr)
 
  // Change message using pointer de-referencing
  *pStr = "Hello Universe"
  fmt.Println("Message = ", *pStr)
  fmt.Println("Message Address = ", pStr)
}
```

pass by value:

```golang
package main

import "fmt"

// pass by reference
func zero(x int) {
  x = 0
}

func main() {
  x := 5
  zero(x)
  fmt.Println(x)
}
```

pass by reference:

```golang
package main

import "fmt"

// pass by value
func zero(xPtr *int) {
  fmt.Println(xPtr)  // memory address of x
  fmt.Println(*xPtr) // value of x
  *xPtr = 0
}

func main() {
  x := 5
  zero(&x)        // pass the memory address of x
  fmt.Println(&x) // memory address of x stays the same
  fmt.Println(x)  // x is 0
}
```

## Pointers in slicing

Whenever you see a slice you see pointers in Go.

A slice is a reference to a section of an array. So it does not create a copy of the sliced part, it only holds a pointer.

## new & make

There are two different ways to create a data structure in Go, they’re `new` and `make`. The new returns a pointer to the created structure, while make returns the structure itself, not a pointer to it.

```
new(T) returns a pointer to a new T
make(T) returns a new T
```

- _further info:_
  - [variable - diff between new and make]({% link docs/languages/golang/variable.md %}#difference-between-new-and-make-in-golang-)

----

[^1]: [Go Pointers](https://medium.com/@nurettinabaci/go-pointers-a538c457a62e)
