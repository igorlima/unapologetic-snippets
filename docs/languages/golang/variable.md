---
layout: default
title: Variable
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/variable
---

# Other

## Zero Values

Unlike JavaScript, Go does not have something akin to the JavaScript undefined primitive value.

Variables declared without an explicit initial value are given their zero value.


## Difference between `new()` and `make()` in Golang [^1]

In Golang, there are two ways to allocate memory: and new . While both of these keywords can be used to allocate memory, they are used for different purposes.

The `new` keyword in Golang is used to create a new instance of a variable, and it __returns a pointer__ to the memory allocated. It is used for allocating memory for the value of a certain type. `new` takes a type as its argument and returns a pointer to a newly allocated zero value of that type.

The `make` keyword in Golang is used to create slices, maps, and channels, and it returns a value of the type that was created. Unlike `new`, `make` __returns a value__ of the type being created, __not a pointer__ to the type.

Keypoints
- `new` returns a pointer to the memory allocated, while `make` returns the value of the type being created.
- `new` only works with basic types such as `int`, `float`, `bool`, etc. `make` is used for creating slices, maps, and channels.
- `new` allocates zeroed memory, while `make` allocates memory and initializes it.

```golang
import "fmt"

func main() {
  var x *int = new(int)
  *x = 5

  y := new(int)
  *y = 6

  fmt.Println(*x)
  fmt.Println(*y)

  var s []int = make([]int, 3)
  s[0] = 1
  s[1] = 2
  s[2] = 3

  r := make([]int, 2)
  r[0] = 7
  r[1] = 8

  w := make([]int, 4)

  fmt.Println(s)
  fmt.Println(r)
  fmt.Println(w)

  // output
  // 5
  // 6
  // [1 2 3]
  // [7 8]
  // [0 0 0 0]
}
```

----

# Footnotes

[^1]: [Difference between new() and make() in Golang](https://saeed0x1.medium.com/difference-between-new-and-make-in-golang-f163b33236ee)
