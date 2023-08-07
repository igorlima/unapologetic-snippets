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

</details>
