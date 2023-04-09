---
layout: default
title: Data Types
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/data-types
---

#  Data Types

1. __Numeric types__: These include integer types such as `int`, `int8`, `int16`, `int32`, and `int64`, as well as unsigned integer types such as `uint`, `uint8` (or `byte`), `uint16`, `uint32`, and `uint64`. In addition to these, there are also floating-point types `float32` and `float64`.
1. __Boolean type__: The boolean type in Go is represented by the values `true` and `false`. It is used for logical comparisons and control flow in a program.
1. __String type__: A string is a sequence of characters enclosed in double quotes, such as `"hello world"`. Go provides many built-in functions for manipulating strings.
1. __Array and slice types__: Arrays and slices are used to store a collection of elements of the same type. An array has a fixed size, whereas a slice is a dynamically-sized, flexible view of an underlying array.
1. __Map type__: A map is a collection of key-value pairs, where each key maps to a corresponding value. Maps are a powerful tool for storing and accessing data in a program.
1. __Pointer type__: A pointer is a variable that holds the memory address of another variable. Pointers are used extensively in Go for efficient memory management and to enable passing values by reference.
1. __Struct type__: A struct is a composite data type that groups together zero or more values with different types under a single name.

```golang
package main

import "fmt"

func main() {
  // Integer type, representing the number of items in a shopping cart
  var numItems int = 10

  // Unsigned integer type, representing the number of likes on a social media post
  var numLikes uint = 1000

  // Floating point type, representing the price of a product
  var price float64 = 19.99

  // Boolean type, representing whether a user is logged in
  var isLoggedIn bool = true

  // String type, representing a user's name
  var userName string = "John Doe"

  // Array type, representing the colors of a rainbow
  var rainbowColors [7]string = [7]string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}

  // Slice type, representing the items in a user's shopping cart
  var cartItems []string = []string{"shirt", "pants", "shoes"}

  // Map type, representing the number of views for each page on a website
  var pageViews map[string]int = map[string]int{
      "home": 1000,
      "about": 500,
      "contact": 200,
  }

  // Pointer type, representing the memory address of a variable
  var itemPtr *string = &cartItems[0]

  // Struct type, representing a user's profile information
  type UserProfile struct {
      Name string
      Age  int
      Email string
      Phone string
  }

  var user UserProfile = UserProfile{
      Name: "Jane Doe",
      Age:  25,
      Email: "jane@example.com",
      Phone: "555-555-5555",
  }

  // Output
  fmt.Println(numItems)
  fmt.Println(numLikes)
  fmt.Println(price)
  fmt.Println(isLoggedIn)
  fmt.Println(userName)
  fmt.Println(rainbowColors)
  fmt.Println(cartItems)
  fmt.Println(pageViews)
  fmt.Println(itemPtr)
  fmt.Println(user)
}
```

# Zero value

The zero value of a data type is the default value that a variable of that type will have if no value is explicitly assigned to it.

- `int`, `int8`, `int16`, `int32`, and `int64`: Default value is `0`.
- `uint`, `uint8`, `uint16`, `uint32`, and `uint64`: Default value is `0`.
- `float32` and `float64`: Default value is `0`.
- `bool`: Default value is `false`.
- `string`: Default value is an empty string `""`.
- `byte`: Default value is `0`.
- `rune`: Default value is `0`. (which is equivalent to the null rune `'\0'`)
- `complex64` and `complex128`: Default value is `(0+0i)`. (a complex number with real and imaginary parts both equal to `0`)

For composite types like arrays, slices, and maps, the default value is `nil`.

Understanding the zero value of a data type can help you write more concise and readable code.
