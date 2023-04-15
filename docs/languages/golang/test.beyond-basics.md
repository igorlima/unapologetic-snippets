---
layout: default
title: Golang Testing Beyond the Basics
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/test-beyond-basics
---

<br/>
<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

# Golang Testing Beyond the Basics

__[back]({% link docs/languages/golang/test.md %})__
<br/>

Discover advanced Go testing features like parallel testing, subtests, teardown functions, and test helpers to enhance your testing capabilities... [^1]

```sh
go mod init add
go mod tidy
```

```golang
// add.go
package add

func Add(a ...int) int {
  total := 0
  for i := range a {
    total += a[i]
  }
  return total
}
```

## Dealing with Failed Marking

How to mark a test as failed in Go? There are several ways to do this besides simply returning `t.Fail()`. A brief summary of the different options available:

- `t.Fail()`: This marks the test as failed, but allows the execution to continue.
- `t.FailNow()`: marks the test as failed and stops its execution immediately using `runtime.Goexit()`.
- `t.Errorf()`: This combines logging an error message `t.Logf()` with failing the test `t.Fail()`.
- `t.Fatalf()`: combines logging an error message using `t.Logf()` with immediately failing the test using `t.FailNow()`.

## Table Driven Test

Overly duplicated and cumbersome? Instead, use a table-driven test approach:

```golang
// add_test.go
package add

import "testing"

func TestAdd(t *testing.T) {
  testCases := []struct {
    args []int
    want int
  }{
    {args: []int{1, 2, 3}, want: 6},
    {args: []int{-1, -2}, want: -3},
    {args: []int{0}, want: 0},
    {args: []int{-1, 2}, want: 1},
  }

  for _, tc := range testCases {
    if res := Add(tc.args...); res != tc.want {
      t.Errorf("Add(%v) = %d; want %d", tc.args, res, tc.want)
    }
  }
}
```

It can’t tell much about the test apart from whether it failed or passed. Also, no way to know how many test cases there were, or which particular test case failed.

### Subtests: Running Multiple Testcases

To use subtests in the testing package, it's needed to get familiar with a new function called `t.Run()`:
- `t.Run(name string, f func(t *testing.T)) (isSuccess bool)`

`t.Run()` creates a subtest with the given name and runs the function `f` in a separate goroutine. Even though each subtest runs in its own goroutine, they run **sequentially**. 

```golang
// add_test.go
package add

import "testing"

func TestAdd(t *testing.T) {
  testCases := []struct {
    name string
    args []int
    want int
  }{
    {name: "case 1 2 3", args: []int{1, 2, 3}, want: 6},
    {name: "case -1 -2", args: []int{-1, -2}, want: -3},
    {name: "case 0", args: []int{0}, want: 0},
    {name: "case -1 2", args: []int{-1, 2}, want: 1},
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      if res := Add(tc.args...); res != tc.want {
        t.Errorf("Add(%v) = %d; want %d", tc.args, res, tc.want)
      }
    })
  }
}
```

```sh
# command to get a verbose output
go test -v
```

```
=== RUN   TestAdd
=== RUN   TestAdd/case_1_2_3
=== RUN   TestAdd/case_-1_-2
=== RUN   TestAdd/case_0
=== RUN   TestAdd/case_-1_2
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/case_1_2_3 (0.00s)
    --- PASS: TestAdd/case_-1_-2 (0.00s)
    --- PASS: TestAdd/case_0 (0.00s)
    --- PASS: TestAdd/case_-1_2 (0.00s)
PASS
ok      add     0.328s
```

### Running Subtests Concurrently

To run subtests in parallel, use `t.Parallel()` to turn on the parallel mode. This can be useful when test cases are independent of each other, since it can make our tests run faster:

```golang
// add_test.go
package add

import "testing"

func TestAdd(t *testing.T) {
  testCases := []struct {
    name string
    args []int
    want int
  }{
    {name: "case 1 2 3", args: []int{1, 2, 3}, want: 6},
    {name: "case -1 -2", args: []int{-1, -2}, want: -3},
    {name: "case 0", args: []int{0}, want: 0},
    {name: "case -1 2", args: []int{-1, 2}, want: 1},
  }

  for _, tc := range testCases {
    tc := tc // what does it do? This addresses a problem with closures in Go
    t.Run(tc.name, func(t *testing.T) {
      t.Parallel() // <---- mark this line
      if res := Add(tc.args...); res != tc.want {
        t.Errorf("Add(%v) = %d; want %d", tc.args, res, tc.want)
      }
    })
  }
}
```

Below is some code to help illustrate the closure concept:

```golang
func main() {
  for i := 0; i < 3; i++ {
    go func() {
      fmt.Println(i)
    }()
  }
  
  time.Sleep(5 * time.Second)
}
```

----

[^1]: [Take Golang Testing Beyond the Basics](https://levelup.gitconnected.com/take-golang-testing-beyond-the-basics-960ae3705a76)
