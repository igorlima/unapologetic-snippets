---
layout: default
title: Pointer
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/pointer
---

Other examples:
- [go pointers]({% link docs/languages/golang/pointer.2023a03m05d.md %}) [^2]
- [did u really know about pointers in golang?]({% link docs/languages/golang/pointer.2023a04m08d.md %}) [^1]


# Pointer

a one-pager to understanding pointers in Go

```golang
package main

import "fmt"

func main() {
  // a pointer is a variable that holds the memory address of another variable
  var p1 *int
  x := 42
  y := 40
  // to create a pointer to a variable, use the & operator followed by the variable's name
  p1 = &x
  p2 := &y
  // to access the value of the variable pointed to by a pointer, you use the * operator followed by the pointer variable name
  *p1 = 100

  increment(&y)
  increment(p2)

  fmt.Println(x)
  fmt.Println(p2)
  fmt.Println(&p2)
  fmt.Println(*p2)
}

// passing pointers to functions
// one common use of pointers is to pass a variable to a function by reference, allowing the function to modify the variable directly.
func increment(p *int) {
  *p++
}
```


```golang
package main

import "fmt"

func main() {
  var creature string = "shark"
  var pointer *string = &creature

  fmt.Println("creature = ", creature)
  fmt.Println("pointer = ", pointer)

  fmt.Println("*pointer = ", *pointer)
  fmt.Println("&creature = ", &creature)

  // output
  // creature = shark
  // pointer = 0xc000014260
  // *pointer = shark
  // &creature = 0xc000014260

  var creatureB Creature = Creature{Species: "shark"}
  fmt.Printf("1) %+v\n", creatureB)
  changeCreature(&creatureB)
  fmt.Printf("3) %+v\n", creatureB)
  // output
  // 1) {Species:shark}
  // 2) &{Species:jellyfish}
  // 3) {Species:jellyfish}

  var creatureC *Creature = &Creature{Species: "shark"}
  fmt.Printf("1) %+v\n", creatureC)
  changeCreature(creatureC)
  fmt.Printf("3) %+v\n", creatureC)
  // output
  // 1) {Species:shark}
  // 2) &{Species:jellyfish}
  // 3) &{Species:jellyfish}

  creatureD := &Creature{Species: "shark"}
  fmt.Printf("1) %+v\n", creatureD)
  changeCreature(creatureD)
  fmt.Printf("3) %+v\n", creatureD)
  // output
  // 1) {Species:shark}
  // 2) &{Species:jellyfish}
  // 3) &{Species:jellyfish}
}

type Creature struct {
  Species string
}

func changeCreature(creature *Creature) {
  creature.Species = "jellyfish"
  fmt.Printf("2) %+v\n", creature)
}
```

----

[^1]: [Did u really know about Pointers in Golang?](https://medium.com/@achmadrizkinf/did-u-really-know-about-pointers-in-golang-3e8be6ff668c)
[^2]: [Go Pointers](https://medium.com/@nurettinabaci/go-pointers-a538c457a62e)
