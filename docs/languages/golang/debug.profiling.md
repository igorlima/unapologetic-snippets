---
layout: default
title: Profiling
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/debug-profiling
---

# Profiling
__[back]({% link docs/languages/golang/debug.md %}#go-doc)__
<br/>

Profiling is a way to analyse the performance and behavior of a Go program, such as its CPU usage, memory allocation, blocking operations, and more.

## sample one [^1]

A function that creates a string, gets a string length, and creates a string from the length.

<details markdown="block">
  <summary>
    sample code
  </summary>

```golang
// main.go
package main

import (
  "fmt"
  "strconv"
)

func createString(size int) string {
  var str string
  for i := 0; i < size; i++ {
    str += strconv.Itoa(i) // convert integer to string and concatenate
  }
  return str
}

func main() {
  fmt.Println("vim-go")
}
```

```golang
// main_test.go
package main

import "testing"

func BenchmarkCreateString(b *testing.B) {
  for i := 0; i < b.N; i++ {
    createString(1000)
  }
}
```

```sh
go mod init ilima.xyz/profiling-in-golang/20231228
go mod tidy

# `go test -bench=.` Runs all available benchmarks in the current package or directory.
# `-bench=CreateString` This is an optional parameter. If specified, only benchmark functions whose names contain the pattern “CreateString” will be run.
# `-benchmem` Enables memory profiling and generates a memory profile file (mem.prof) after the benchmarks are run.
# `-cpuprofile cpu.prof` Enables CPU profiling and generates a CPU profile file (cpu.prof) after the benchmarks are run.
# `-benchtime=5s` Sets the benchmark duration to 5 seconds.
go test -bench=. -benchmem -cpuprofile cpu.prof -memprofile mem.prof -benchtime=5s

# run this command below to see how much memory our test has consumed
# “mem.prof” is the name of the file that has been generated by the command we already used above.
go tool pprof mem.prof
help
top
list profiling-in-golang/20231228.createString
exit

# check the cpu.prof
go tool pprof cpu.prof
tree
list profiling-in-golang/20231228.createString
exit

# let’s check them graphically to see how it’s going by adding the -http flag to the pprof command.
go tool pprof -http=:8080 cpu.prof
# run the command below to show a specific memory profile generated using the pprof tool
go tool pprof -http=:8080 mem.prof
```

| `cpu.prof` | `mem.prof` |
| :---: | :---: |
| ![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/d53f5f67-f95b-4cc5-848d-07d76ef98dbf) | ![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/c9ab5b80-b7a4-4838-b37e-017b909411fc) |

  <details markdown="block"><summary>fixing it...</summary>

```golang
func createStringOne(size int) string {
  var str string
  for i := 0; i < size; i++ {
    str += strconv.Itoa(i) // Convert integer to string and concatenate
  }
  return str
}

func createStringTwo(size int) string {
  var buffer bytes.Buffer
  for i := 0; i < size; i++ {
    buffer.WriteString(strconv.Itoa(i))
  }

  return buffer.String()
}
```

```golang
func BenchmarkCreateStringOne(b *testing.B) {
  for i := 0; i < b.N; i++ {
    createStringOne(1000)
  }
}

func BenchmarkCreateTestTwo(b *testing.B) {
  for i := 0; i < b.N; i++ {
    createStringTwo(1000)
  }
}
```
  </details>

<br/>
</details>


----

[^1]: [Optimize your codes by profiling in Golang.](https://medium.com/@mohammad.sabbag70/optimize-your-codes-by-profiling-in-golang-05573c23e2a3)
