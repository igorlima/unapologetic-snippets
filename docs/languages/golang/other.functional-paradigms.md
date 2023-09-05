---
layout: minimal
title: Functional Paradigms in Go
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/other/functional-paradigms
---

__[go to Golang - other resources session]({% link docs/languages/golang/other.md %})__

## Functional Paradigms in Go

<details markdown="block">
  <summary>
    <sup>Mapping & Filtering</sup>
  </summary>

A sample from a Medium post [^1].

```golang
package main

import "fmt"

type Predicate func(int) bool
type Transformer func(int) int

func Map(nums []int, fn Transformer) []int {
  result := make([]int, len(nums))
  for i, v := range nums {
    result[i] = fn(v)
  }
  return result
}

func Filter(nums []int, fn Predicate) []int {
  var result []int
  for _, v := range nums {
    if fn(v) {
      result = append(result, v)
    }
  }
  return result
}

func main() {
  nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  // Double each number
  doubled := Map(nums, func(n int) int {
    return n * 2
  })

  // Filter out even numbers
  odds := Filter(nums, func(n int) bool {
    return n%2 != 0
  })

  fmt.Println(doubled) // [2 4 6 8 10 12 14 16 18 20]
  fmt.Println(odds)    // [1 3 5 7 9]
}
```
----
<br/>
<!-- Mapping & Filtering -->
</details>



<details markdown="block">
  <summary>
    <sup>Closures with State</sup>
  </summary>

A sample from a Medium post [^1].

```golang
package main

import "fmt"

func Counter(start int) func() int {
  count := start
  return func() int {
    count++
    return count
  }
}

func main() {
  counterA := Counter(0)
  fmt.Println(counterA()) // Outputs: 1
  fmt.Println(counterA()) // Outputs: 2

  counterB := Counter(10)
  fmt.Println(counterB()) // Outputs: 11
}
```
----
<br/>
<!-- Closures with State -->
</details>



<details markdown="block">
  <summary>
    <sup>Immutability with Structs</sup>
  </summary>

A sample from a Medium post [^1].

```golang
package main

import "fmt"

type ImmutablePoint struct {
  X, Y int
}

func NewImmutablePoint(x, y int) *ImmutablePoint {
  return &ImmutablePoint{x, y}
}

func (p *ImmutablePoint) Move(deltaX, deltaY int) *ImmutablePoint {
  // Instead of mutating the existing point, return a new one
  return NewImmutablePoint(p.X+deltaX, p.Y+deltaY)
}

func main() {
  point := NewImmutablePoint(2, 3)
  newPoint := point.Move(1, 1)

  fmt.Printf("Original Point: %v\n", point) // {2 3}
  fmt.Printf("Moved Point: %v\n", newPoint) // {3 4}
}
```
----
<br/>
<!-- Immutability with Structs -->
</details>

----


[^1]: [Embracing Functional Paradigms in Go](https://blog.stackademic.com/embracing-functional-paradigms-in-go-d9f6a9963a96)
