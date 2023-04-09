---
layout: default
title: Type Assertion
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/type-assertion-2023a03m10d
---

__[back]({% link docs/languages/golang/type-assertion.md %})__

# Type Assertion

Assertion is used to ensure that the value of a variable conforms to a certain type or interface before executing operations on it. [^1]

## Basic Assertion

In Go, basic assertion is performed using the `.` (dot) operator and the type assertion expression `x.(T)`, where `x` is the value being asserted, and `T` is the expected type. If the assertion is successful, the result is the underlying value of type `T`. If the assertion fails, a runtime error occurs.

Here is an example of basic assertion:

```golang
var i interface{} = "hello"
s := i.(string)
fmt.Println(s) // output: hello
```

## Comma Ok Idiom

Here is an example of the comma ok idiom:

```golang
var i interface{} = 42
s, ok := i.(string)
if ok {
  fmt.Println(s) // not executed
} else {
  fmt.Println("i is not a string") // output: i is not a string
}
```

## Type Assertion in Struct

```golang
type Person struct {
  name string
  age  int
}

func main() {
  var p interface{}
  p = Person{name: "Alice", age: 25}

  // Type assertion to access the name field
  fmt.Println(p.(Person).name)

  // Type assertion to access the age field
  fmt.Println(p.(Person).age)
}
```

## Type Assertion with Interfaces

```golang
type Shape interface {
  Area() float64
}

type Rectangle struct {
  width  float64
  height float64
}

type Circle struct {
  radius float64
}

func (r Rectangle) Area() float64 {
  return r.width * r.height
}

func (c Circle) Area() float64 {
  return math.Pi * c.radius * c.radius
}

func main() {
  var s Shape
  s = Rectangle{width: 2, height: 3}

  // Type assertion to call the Area method on the Rectangle struct
  fmt.Println(s.(Rectangle).Area())

  s = Circle{radius: 1}

  // Type assertion to call the Area method on the Circle struct
  fmt.Println(s.(Circle).Area())
}
```

## Regular Assertion

```golang
var i interface{} = 42

// Regular assertion to check if i is of type int
if v, ok := i.(int); ok {
    fmt.Println("i is of type int")
} else {
    fmt.Println("i is not of type int")
}
```

## Tips

It’s important to note that if we attempt to use a type assertion to extract a value of the wrong type, a runtime panic will occur.

Therefore, it's important to always use the `ok` variable to check if the type assertion succeeded before using the extracted value.

When working with pointers, use type assertions on the pointer type rather than the value type.

----

[^1]: [Mastering Type Assertion in Go: A Comprehensive Guide](https://medium.com/@jamal.kaksouri/mastering-type-assertion-in-go-a-comprehensive-guide-216864b4ea4d)
