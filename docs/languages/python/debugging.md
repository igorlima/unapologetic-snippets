---
layout: default
title: Debugging
nav_order: 4
parent: Python
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/python/debugging
---

# Debugging

> “The art of debugging is figuring out what you really told your program to do rather than what you thought you told it to do.” _Andrew Singer_

## Using Rich Inspect to interrogate Python objects

<details markdown="block">
  <summary>
    sample
  </summary>

```python
# pip show rich
# pip install rich==13.5.2

from rich import inspect
text_file = open("foo.txt", "w")
inspect(text_file)
inspect(text_file, methods=True)
inspect(text_file, methods=True, help=True)
```

![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/16cfa090-7093-4b6e-add7-7014a1a7b73f)

![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/91ece1da-bbfd-4381-b9fc-2c01a18f877a)

</details>
