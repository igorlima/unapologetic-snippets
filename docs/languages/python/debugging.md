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

```python
# Full-screen console debugger for Python
# https://documen.tician.de/pudb/
# https://github.com/inducer/pudb
import pudb; pu.db
t = 10
y = 'Hello'
pu.db
```

<details markdown="block">
  <summary>
    <sup><i>further more...</i></sup>
  </summary>

__Breakpoints__ can be set just by pointing at a source line and hitting `b` and then edited visually in the breakpoints window. Or hit `t` to run to the line under the cursor.

Drop to a Python shell in the current environment by pressing `!`. Or open a command prompt alongside the source-code via `Ctrl-X`.

```sh
# https://documen.tician.de/pudb/starting.html
PYTHONBREAKPOINT="pudb.set_trace" python3 -c '
def hello_world():
  name = "you"
  print(f"Hey {name}")
breakpoint()
'
```

```sh
python3 -c '
import pudb
def hello_world():
  name = "you"
  print(f"Hey {name}")
pu.db
'
```
</details>

![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/a30a6667-0375-41ba-8972-73e6c5127934)


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

## Building a small REPL in Python

Python comes with a bunch of batteries included, starting from embedding a normal Python REPL in the project. [^1]

<details markdown="block">
  <summary>
    sample
  </summary>

```python
import code

print("Welcome to the Python interactive shell!")

t = 10
y = 'Hello'

# • dir()      will give you the list of in-scope variables
# • globals()  will give you a dictionary of global variables
# • locals()   will give you a dictionary of local variables
# print(locals())
# print(globals())

# Documentation:
# Interpreter base classes
# https://docs.python.org/3/library/code.html#code.InteractiveConsole
repl = code.InteractiveConsole(locals=locals())
repl.interact()
```

If you (`chmod +x repl.py` and) run this, you get what looks like a normal Python REPL, plus a little extra output:

```python
#!/usr/bin/env python3
# repl.py
import code
repl = code.InteractiveConsole()
repl.interact()
```

</details>

## Other alternatives for debugging

- libraries
  - [web-pdb](https://github.com/romanvm/python-web-pdb)
    - _Web interface for Python's built-in PDB debugger_
  - [birdseye](https://github.com/alexmojaki/birdseye)
    - _Graphical Python debugger which lets you easily view the values of all evaluated expressions_


------ ------

[^1]: [Building a small REPL in Python](https://bernsteinbear.com/blog/simple-python-repl/)
