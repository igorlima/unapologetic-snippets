---
layout: default
title: Control Structures and Loops
nav_order: 4
parent: Python
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/python/control-structures
---

# Control Structures and Loops

This is an draft page.

## Iteration

```python
# iterator
print('iterator')
iterator = iter((1, 2, 3))
while True:
  try:
    print(next(iterator))
  except StopIteration:
    break
print()

# tuple
print('tuple')
for n in (1, 2, 3):
  print(n)
print()

# list
print('list')
for n in [1, 2, 3]:
  print(n)
print()

# dict
print('dict')
for n in {1: 10, 2: 20, 3: 30}:
  print(n)
# for Python 3.x:
for key, value in {1: 10, 2: 20, 3: 30}.items():
  print(key, value)
print()

# string
print('string')
for n in "123":
  print(n)
print()
```
