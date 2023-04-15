---
layout: default
title: Golang Testing Beyond the Basics
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/test-cover-bench-fuzz
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

# A Guide to Testing in Go

__[back]({% link docs/languages/golang/test.md %})__
<br/>

Testing is an essential part of the development process, and it’s a critical part of the software development life cycle. It ensures that the app functionality correctly and meets a given need. [^1]

## Skipping Tests

Using the `Skip()` method allows you to separate unit tests from integration tests. 

`go test` does not decide whether tests are “short” or not. It requires to use a combination of `testing.Short()`, which is set to true when `-short` is used, and `t.Skip()`, as illustrated below:

```sh
go test -v -test.short
```

```golang
func TestFooerSkiped(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping test in short mode.")
  }
  result := Fooer(3)
  if result != "Foo" {
    t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
  }
}
```

## Coverage Tests

```sh
# note it can also have the flag `-v` for more detailed logs
go test -cover
```

Use `-coverpkg` to include all packages in the coverage calculation:

```sh
go test ./... -coverpkg=./...
```

Using the flag `-coverprofile` will create a local coverage report file. 

```sh
go test -coverprofile=output_filename
```

Use `go tool cover` to format the report. For example, the `-html` flag will open the default browser to display a graphical report.

```sh
go tool cover -html=output_filename
```

## Benchmark Tests

To create a benchmark test:

- the test function needs to be in a `*_test.go` file.
- the name of the function must start with `Benchmark`.
- the function must accept `*testing.B` as the unique parameter.
- the test function must contain a `for` loop using `b.N` as its upper bound.

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

```golang
// add_test.go
package add

import (
  "fmt"
  "testing"
)

func BenchmarkAdd(b *testing.B) {
  testCases := []struct{ args []int }{
    {args: []int{1, 2, 3}},
    {args: []int{-1, -2}},
    {args: []int{0}},
    {args: []int{-1, 2}},
  }

  for _, tc := range testCases {
    b.Run(fmt.Sprintf("args %#v", tc.args), func(b *testing.B) {
      for i := 0; i < b.N; i++ {
        Add(tc.args...)
      }
    })
  }
}
```

```sh
go test -bench=.
```

```
goos: darwin
goarch: arm64
pkg: add
BenchmarkAdd/args_[]int{1,_2,_3}-10             504180847                2.183 ns/op
BenchmarkAdd/args_[]int{-1,_-2}-10              640850622                1.871 ns/op
BenchmarkAdd/args_[]int{0}-10                   1000000000               0.9360 ns/op
BenchmarkAdd/args_[]int{-1,_2}-10               641459389                1.870 ns/op
PASS
ok      add     6.528s
```

To include memory allocation statistics in the benchmark output, add the `-benchmem` flag while running the benchmarks:

```sh
go test -bench=. -benchmem
```

### Comparing benchmark results

```sh
go test -bench=Add -count 1 > add-implementation-a.txt
# the `tee` command sends the output of the command to the specified file and prints it to the standard output
go test -bench=Add -count 1 | tee add-implementation-b.txt

# view the benchmark’s results with `benchstat`
go install golang.org/x/perf/cmd/benchstat@latest

benchstat add-implementation-a.txt add-implementation-b.txt
```

```
goos: darwin
goarch: arm64
pkg: add
                           │ add-implementation-a.txt │       add-implementation-b.txt        │
                           │          sec/op          │    sec/op      vs base                │
Add/args_[]int{1,_2,_3}-10               2.183n ± ∞ ¹    2.183n ± ∞ ¹       ~ (p=1.000 n=1) ²
Add/args_[]int{-1,_-2}-10                1.873n ± ∞ ¹    1.870n ± ∞ ¹       ~ (p=1.000 n=1) ³
Add/args_[]int{0}-10                    0.9329n ± ∞ ¹   0.9399n ± ∞ ¹       ~ (p=1.000 n=1) ³
Add/args_[]int{-1,_2}-10                 1.870n ± ∞ ¹    1.871n ± ∞ ¹       ~ (p=1.000 n=1) ³
geomean                                  1.634n          1.637n        +0.16%
¹ need >= 6 samples for confidence interval at level 0.95
² all samples are equal
³ need >= 4 samples to detect a difference at alpha level 0.05
```

## Fuzz Tests

To create a fuzz test:

- to test function needs to be in a `_test.go` file.
- the name of the function must start with `Fuzz`.
- the test function must accept `testing.F` as the unique parameter.
- the test function must define initial values, called _seed corpus_, with the `f.Add()` method.
- the test function must define a _fuzz target_.

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

```
// add_test.go
package add

import (
  "testing"
)

func FuzzAdd(f *testing.F) {
  f.Add(10, 100, -20)
  f.Fuzz(func(t *testing.T, a, b, c int) {
    Add(a, b, c)
  })
}
```

```sh
go test -fuzz FuzzAdd
go test -fuzz .
# by default, fuzzing will run indefinitely, as long as there isn’t a failure
# the `-fuzztime` flag should be to specify the maximum time for which fuzzing should run.
go test -test.fuzztime=10s -fuzzminimizetime=10s -test.fuzz FuzzAdd
go test -test.fuzztime=10s -fuzzminimizetime=10s -test.fuzz "^FuzzAdd$"
```

```
fuzz: elapsed: 0s, gathering baseline coverage: 0/1 completed
fuzz: elapsed: 0s, gathering baseline coverage: 1/1 completed, now fuzzing with 10 workers
fuzz: elapsed: 3s, execs: 1458678 (486116/sec), new interesting: 0 (total: 1)
fuzz: elapsed: 6s, execs: 2916276 (485949/sec), new interesting: 0 (total: 1)
fuzz: elapsed: 9s, execs: 4393201 (492290/sec), new interesting: 0 (total: 1)
fuzz: elapsed: 10s, execs: 4884244 (442683/sec), new interesting: 0 (total: 1)
PASS
ok      add     10.449s
```

## The Tesitfy Package

[Testify](https://github.com/stretchr/testify) is a testing framework providing many tools for testing Go code. There is a considerable debate in the Go community about whether you should use Testify or just the standard library.

Testify provides assert functions and mocks, which are similar to traditional testing frameworks, like JUnit for Java or Jasmine for NodeJS.

----

[^1]: [Comprehensive Guide to Testing in Go](https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/)
