---
layout: default
title: Type Assertion - Interface Check
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/type-assertion-interface-check
---

__[back]({% link docs/languages/golang/type-assertion.md %})__

```golang
package main

import "fmt"

// Animal is an interface that requires the Speak method
type Animal interface {
  Speak() string
}

// Dog is a type that implements the Animal interface.
type Dog struct{}

func (d Dog) Speak() string {
  return "Woof!"
}

/*
`var _ Animal = (*Dog)(nil)` ensures at compile time that `Dog`
implements the `Animal` interface. If `Dog` did not have a `Speak`
method, this line would cause a compile error.

Why Need It? These checks are particuraly useful in larger codebases where an
interface and its implementations may be defined in different packages. They
provide a way to ensure that a type correctly implements an interface at the
point where the type is defined, rather than waiting for a runtime error when
the type is used.

Using interface checks, and placing them next to implementations meant to
satisfy any given interface, will protect you in the rare occurrence that there
isn't a static conversion elsewhere in the application.
*/
func main() {
  fmt.Println("vim-go")

  // this is an interface check - it will compile only if Dog implements the Animal interface
  var _ Animal = (*Dog)(nil)

  var a Animal = Dog{}
  fmt.Println(a.Speak())
}
```

----

[^1]: [...](...)
