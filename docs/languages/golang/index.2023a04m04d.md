---
layout: default
title: Golang
nav_exclude: true
parent: Programming Languages
permalink: /docs/languages/golang-2023a04m04d
---

__[back]({% link docs/languages/golang/index.md %})__

# Golang

Statements in Go can be classified into various categories such as, [^1]
- __Declaration statements__ — use to declare variables, constants, and types.
- __Control flow statements__ — use to control the flow of execution.
- __Function statements__ — use to define and invoke functions.
- __Advanced statements__ — more complex statements that deal with advanced Go features.

## Control flow statements

__Panic and Recover__

The __panic__ statement is used to `stop the normal execution of a program with error`.

The __recover__ function is used to `handle the panic` and `continues the execution`. It returns the value that was passed to the `panic` function, which can be used to return an appropriate error response to the client.


__Goto statements__

This allows you to jump to a _labeled statement_ within the same function.

Here’s an example of how you can use the __goto__ statement:

```golang
func main() {
  i := 0
loop:
  if i < 10 {
    fmt.Printf("%d ", i)
    i++
    goto loop
  }
}
//output
0 1 2 3 4 5 6 7 8 9
```

It’s worth noting that the use of `goto` is limited to the same function and can only be used to jump to a labeled statement within that function.


__Break and Continue__

These statements are used to alter the normal flow of execution in loops.

The __break__ is used to exit a loop prematurely inside a `for`, `switch`, or `select` statement. It immediately terminates the innermost loop and control resumes at the next statement following the loop.

__continue__ is used to skip the current iteration of a loop and move on to the next iteration. When `continue` is encountered, the current iteration is immediately terminated and the loop continues with the next iteration.


## Functional statements

__Anonymous function__

An anonymous function is a function without a name. It can be defined inline and used immediately, or assigned to a variable and used later.


__Function closures__

A function closure is created when a function returns an inner function that references a variable defined in the outer function’s scope.


__Blank identifier__

We all are using `_` (underscore), when we want to ignore the returned values of functions. This is known as a blank identifier in Go.


__Variadic functions__

Variadic functions are functions that can take a dynamic number of arguments. This is useful when you don’t know how many arguments you will need to pass to a function. It will take ... as an argument.


## Other statements

__Type assertions__

Type assertions are used to convert an interface value to a concrete type. This is useful when you have an interface value and you want to use it as a specific type.

```golang
func performOperation(args ...interface{}) {
    for _, arg := range args {
        switch val := arg.(type) {
        case int:
            // Handle integer input
            fmt.Println("Integer input:", val)
        case float64:
            // Handle floating-point input
            fmt.Println("Float input:", val)
        case string:
            // Handle string input
            fmt.Println("String input:", val)
        default:
            // Handle unrecognized input
            fmt.Println("Unrecognized input:", val)
        }
    }
}
// use function as below
performOperation(77, "hey", 0.54)
```

__Defer statements__

This can be used to defer the execution of a function until the surrounding function has been completed.


__Label statements__

This can be used to define a label for a certain block of code. Labels are most commonly used with the `break,continue and goto` statements to control the flow of execution in nested loops or switch statements.

```golang
outer:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if i*j >= 10 {
            break outer
        }
        fmt.Println(i, j)
    }
}
```

If we had used a `break` statement without the label, it would have only broken out of the inner loop, and the outer loop would have continued to execute.

It’s worth noting that labels can only be used within the same function as the labeled statement. You cannot use a label to jump to a statement in a different function.


__Select statement__

This is used to wait on multiple channels for input. It blocks until one of the channels has data ready to be processed.


----

[^1]: [From Unknown to Expert: Golang Statements You Need to Know](https://blog.canopas.com/from-unknown-to-expert-golang-statements-you-need-to-know-a1ac97213c04)
