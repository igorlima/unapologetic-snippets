---
layout: default
title: Type Assertion
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/type-assertion
---

Other examples:
- [mastering type assertion]({% link docs/languages/golang/type-assertion.2023a03m10d.md %}) [^1]
- [interface check]({% link docs/languages/golang/type-assertion.interface-check.md %})

In Go programming language, assertion is a mechanism to check the dynamic type of a value during runtime.

Type Assertion is used to ensure that the value of a variable conforms to a certain type or interface before executing operations on it.
Assertion is particularly useful when dealing with interface types, where the actual type of the underlying value is unknown.

In Go, basic assertion is performed using the `.` _(dot)_ operator and the type assertion expression `x.(T)`, where `x` is the value being asserted, and `T` is the expected type. If the assertion is successful, the result is the underlying value of type `T`. If the assertion fails, a runtime error occurs.

Comma Ok Idiom. In Go, the comma ok idiom is a way to perform a safe assetion on an interface value without causing a runtime error.

```golang
package main

import "fmt"

func main() {
  var i1 interface{} = "hello"
  s := i1.(string)
  fmt.Println(s) // output: hello

  var i2 interface{} = 42
  s, ok := i2.(string)
  if ok {
    fmt.Println(s) // not executed
  } else {
    fmt.Println("i2 is not a string") // output: i2 is not a string
  }

  var i3 interface{} = 42
  // regular assertion to check if i3 is of type int
  if v, ok := i3.(int); ok {
    fmt.Println("i3 is of type int:", v) // output: i3 is of type int 42
  } else {
    fmt.Println("i3 is not of type int") // not executed
  }

  var p interface{}
  p = Person{name: "Alice", age: 25}
  // type assertion to acces the name field
  fmt.Println(p.(Person).name) // output: Alice
  // type assertion to acces the age field
  fmt.Println(p.(Person).age) // output: 25
}

type Person struct {
  name string
  age  int
}
```

----

[^1]: [Mastering Type Assertion in Go: A Comprehensive Guide](https://medium.com/@jamal.kaksouri/mastering-type-assertion-in-go-a-comprehensive-guide-216864b4ea4d)
