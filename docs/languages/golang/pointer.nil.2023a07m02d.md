---
layout: default
title: Pointer Nil
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/pointer-nil-2023a07m02d
---

__[back]({% link docs/languages/golang/pointer.md %})__

# Understanding `nil`

What is nil? When do we need to use nil? [^1]

__What is nil?__

In Go, nil is a predefined identifier that represents the zero value of many types and has no type of its own.

A __*zero*__ value is the default value assigned to a variable of a given type if no value is specified. Some common zero values in Go include 0 for integers, false for booleans, “” for strings, and __*nil*__ for __*pointers*__, __*slices*__, __*maps*__, __*channels*__, __*interfaces*__, and __*functions*__. Understanding zero values is important when initializing variables to ensure that they are not accidentally left in an uninitialized state.

__How to fix a pointer receiver with nil?__

<details markdown="block">
  <summary>
    panic example
  </summary>

```golang
package main

import "fmt"

type Person struct {
  Name string
}

func (p *Person) GetName() string {
  return p.Name
}

func main() {
  var p *Person            // p is pointer, zero value of p is nil
  fmt.Println(p.GetName()) // panic: runtime error: invalid memory address
  // panic: runtime error: invalid memory address or nil pointer dereference
}
```

<br/>
</details>

<details markdown="block">
  <summary>
    how to fix it?
  </summary>

```golang
package main

import "fmt"

type Person struct {
  Name string
}

func (p *Person) GetName() string {
  if p == nil {
    return "empty"
  }
  return p.Name
}

func main() {
  var p *Person            // p is pointer, zero value of p is nil
  fmt.Println(p.GetName()) // empty
}
```

<br/>
</details>


<details markdown="block">
  <summary>
    another pattern
  </summary>

```golang
package main

import "fmt"

type Node struct {
  next  *Node
  value int
}

func (n *Node) Next() *Node {
  if n == nil {
    return nil
  }

  return n.next
}

func (n *Node) Len() int {
  if n == nil {
    return 0
  }

  cnt := 0

  c := n
  for c != nil {
    c = c.Next()
    cnt++
  }

  return cnt
}

func main() {
  n0 := &Node{value: 2}
  var n1 *Node = &Node{value: 5}
  var n2 *Node
  n0.next = n1
  n1.next = n2

  fmt.Println("root node length:", n0.Len())

  fmt.Println(n2 == nil) // true
  fmt.Println(n2.Len())  // 0
}
```

<br/>
</details>

----

[^1]: [Understanding nil in Golang](https://medium.com/@ahmad1997rezamarashi/understanding-nil-df33d1b0fd31)
