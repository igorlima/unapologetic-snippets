---
layout: default
title: Design Pattern
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/design-pattern
---

# Prototype

```golang
package main

import "fmt"

func main() {
  person1 := &Person{Name: "Alice", Age: 30}
  person2 := person1.Clone().(*Person)
  fmt.Printf("%+v %T %p\n", person1, person1, person1)
  fmt.Printf("%+v %T %p\n", person2, person2, person2)
  // output
  // &{Name:Alice Age:30} *main.Person 0xc000094018
  // &{Name:Alice Age:30} *main.Person 0xc000094030
}

type Cloneable interface {
  Clone() Cloneable
}

type Person struct {
  Name string
  Age  int
}

func (p Person) Clone() Cloneable {
  return &Person{
    Name: p.Name,
    Age:  p.Age,
  }
}
```
