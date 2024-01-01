---
layout: default
title: Bitwise Operations
nav_exclude: true
parent: Python
grand_parent: Programming Languages
permalink: /docs/languages/python/other-bitwise-operations
---

__[back]({% link docs/languages/python/other.md %})__

Some samples of bitwise operations[^1]:
1. checking if a number is odd or even
1. checking if a number is the power of two
1. swapping variable values without using a temporary variable
1. counting bit set to ‘1’ in a binary number
1. toggling a bit
1. faster multiplication & division

## Checking If A Number Is Odd Or Even

```
  10 (2 in decimal)
& 01 (1 in decimal)
----
  00 -> Falsy value -> Even number
```

```
  11 (3 in decimal)
& 01 (1 in decimal)
----
  01 -> Truthy value -> Odd number
```

```python
def is_even(num):
  if (number & 1) == 0:
      print("The number is even.")
  else:
      print("The number is odd.")
```

## Checking If A Number Is The Power of Two

```
  1000 (8 in binary)
& 0111 (7 in binary)
------
  0000 -> 8 is a power of 2
```

```
  0111(7 in binary)
& 0110 (6 in binary)
------
  0110 -> 7 is not a power of 2
```

```python
def is_power_of_two(n):
    if (n > 0 and (n & (n - 1)) == 0):
      print(f"{n} is a power of 2.")
    else:
      print(f"{n} is not a power of 2.")
```

## Swapping Variable Values Without Using A Temporary Variable

```python
a = 5  # 0101 in binary
b = 3  # 0011 in binary

# XOR Swap Algorithm
a = a ^ b  # a now holds 6
b = a ^ b  # b now holds 5
a = a ^ b  # a now holds 3

print("a:", a)  # a: 3
print("b:", b)  # b: 5
```

## Faster Multiplication & Division

```python
num = 32 
power = 3

# Left shift for multiplication by 2^3 i.e. 8
mul_result = num << power  # Result is 256

# Right shift for division by 2^3 i.e. 8
div_result = num >> power  # Result is 4
```


-----

[^1]: [6 Mind-Blowing Uses Of Bitwise Operations That You Probably Never Knew Existed](https://medium.com/gitconnected/6-mind-blowing-uses-of-bitwise-operations-that-you-probably-never-knew-existed-9e0359cfeaaf)
