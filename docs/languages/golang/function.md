---
layout: default
title: Function
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/function
---

A function is just another type — just like `int` and `string` and `bool`. [^1]

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

## a function as a parameter

```golang
package main

import "fmt"

func add(x, y int) int {
  return x + y
}

func mul(x, y int) int {
  return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  return arithmetic(arithmetic(a, b), c)
}

func main() {
  fmt.Println(aggregate(2, 3, 4, add)) // prints 9
  fmt.Println(aggregate(2, 3, 4, mul)) // prints 24
  fmt.Println("vim-go")
}
```

```golang
package main

import "fmt"

func add(x, y int) int {
  return x + y
}

func mul(x, y int) int {
  return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  return arithmetic(arithmetic(a, b), c)
}

func main() {
  fmt.Println(aggregate(2, 3, 4, add)) // prints 9
  fmt.Println(aggregate(2, 3, 4, mul)) // prints 24
  fmt.Println("vim-go")
}
```

## currying

Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function.

```golang
package main

import "fmt"

func main() {
  squareFunc := selfMath(multiply)
  doubleFunc := selfMath(add)

  fmt.Println(squareFunc(5)) // prints 25
  fmt.Println(doubleFunc(5)) // prints 10
}

func multiply(x, y int) int {
  return x * y
}
func add(x, y int) int {
  return x + y
}
func selfMath(mathFunc func(int, int) int) func(int) int {
  return func(x int) int {
    return mathFunc(x, x)
  }
}
```

## closures

A closure is a function that references variables from outside its body. The function may access and assign to the referenced variables.

```golang
package main

import "fmt"

func concatter() func(string) string {
  doc := ""
  return func(word string) string {
    doc += word + " "
    return doc
  }
}

func main() {
  makeHarryPotter := concatter()
  makeHarryPotter("Mr.")
  makeHarryPotter("and")
  makeHarryPotter("Mrs.")
  makeHarryPotter("Dursley")
  makeHarryPotter("of")
  makeHarryPotter("number")
  makeHarryPotter("four,")
  makeHarryPotter("Privet")
  fmt.Println(makeHarryPotter("Drive"))
  // prints Mr. and Mrs. Dursley of number four, Privet Drive
}
```

```golang
package main

import "fmt"

func adder() func(int) int {
  sum := 0
  return func(num int) int {
    sum += num
    return sum
  }
}

func main() {
  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(
      "You've sent",
      pos(i),
      "emails and it has cost you",
      neg(3*i),
      "cents.",
    )
  }
}

/**
OUTPUT

You've sent 0 emails and it has cost you 0 cents.
You've sent 1 emails and it has cost you 3 cents.
You've sent 3 emails and it has cost you 9 cents.
You've sent 6 emails and it has cost you 18 cents.
You've sent 10 emails and it has cost you 30 cents.
You've sent 15 emails and it has cost you 45 cents.
You've sent 21 emails and it has cost you 63 cents.
You've sent 28 emails and it has cost you 84 cents.
You've sent 36 emails and it has cost you 108 cents.
You've sent 45 emails and it has cost you 135 cents.
*/
```

## anonymous functions

Anonymous functions are useful when defining a function that will only be used once or to create a quick closure.

```golang
package main

import "fmt"

// doMath accepts a function that converts one int into another
// and a slice of ints. It returns a slice of ints that have been
// converted by the passed in function.
func doMath(f func(int) int, nums []int) []int {
  var results []int
  for _, n := range nums {
    results = append(results, f(n))
  }
  return results
}

func main() {
  nums := []int{1, 2, 3, 4, 5}

  // Here we define an anonymous function that doubles an int
  // and pass it to doMath
  allNumsDoubled := doMath(func(x int) int {
    return x + x
  }, nums)

  fmt.Println(allNumsDoubled) // prints: [2 4 6 8 10]
}
```

----

[^1]: [Learn Golang — Advanced Functions](https://medium.com/@lordmoma/learn-golang-advanced-functions-e69853996e39)
