---
layout: default
title: Pointer
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/pointer-2023a04m08d
---

# Pointer

## What is a pointer?

A pointer is a variable that stores the memory address of another variable. In other words, it is a reference to the memory location of a variable. By using pointers, you can access and modify the contents of a variable indirectly.

In Go, you can declare a pointer variable by prefixing an asterisk (*) before the variable name. For example:

```golang
var ptr *int
```

## How to use pointers?

To use a pointer, you need to assign it the memory address of a variable. You can obtain the memory address of a variable by using the ampersand (&) operator.

```golang
var a int = 42
var ptr *int = &a
```

To access the value of the variable through the pointer, you can use the dereference operator (*). This operator is used to retrieve the value that is stored at the memory location pointed to by the pointer variable.

```golang
fmt.Println(*ptr)
```

## Unsafe Pointers

This package allows you to work with memory in a way that is not type-safe or guaranteed to be safe. Therefore, you should only use the unsafe package if you know what you’re doing and understand the risks.

The unsafe.Pointer type is used to represent a pointer to an arbitrary memory location. You can convert a pointer to a specific type to an unsafe pointer using the unsafe.Pointer() function.

```golang
var p *int = new(int)
unsafePtr := unsafe.Pointer(p)
```

----

[^1]: [Did u really know about Pointers in Golang?](https://medium.com/@achmadrizkinf/did-u-really-know-about-pointers-in-golang-3e8be6ff668c)
