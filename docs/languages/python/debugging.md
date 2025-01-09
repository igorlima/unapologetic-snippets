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

- using PuDB as REPL:
  - ```sh
    # https://documen.tician.de/pudb/starting.html
    PYTHONBREAKPOINT="pudb.set_trace" python3 -c '
    def hello_world():
      name = "you"
      print(f"Hey {name}")
    breakpoint()
    '
    ```
  - ```sh
    python3 -c '
    import pudb
    def hello_world():
      name = "you"
      print(f"Hey {name}")
    pu.db
    '
    ```
</details>

## IPython


<details markdown="block">
  <summary>
    <sup><i>Python vs IPython</i></sup>
  </summary>

Key differences between Python and IPython:
- Interactive Shell:
  - __Python__ is a general-purpose programming language that can be used to write scripts and applications.
  - __IPython__ is an enhanced interactive Python shell that provides a more powerful and user-friendly environment for working with Python.
- Features:
  - __Python__ is the base language with standard syntax and functionality.
  - __IPython__ adds many extra features on top of Python, such as:
    - Tab completion
    - Syntax highlighting
    - Magic commands (prefixed with `%`)
    - Rich media and visualization support
    - Enhanced debugging and profiling tools
- Code Execution:
  - __Python__ is typically used to run Python scripts or programs.
  - __IPython__ allows you to execute Python code interactively, one line or cell at a time.
- Notebooks:
  - __Python__ does not have a built-in notebook interface.
  - __IPython__ introduced the Jupyter Notebook, which provides a web-based interactive coding environment.
- Compatibility:
  - __Python__ code will generally run in IPython without modification.
  - Some __IPython__-specific features, like magic commands, will not work in a standard Python shell.

*In summary, __IPython__ is an enhanced interactive environment built on top of the
__Python__ programming language, providing a more powerful and user-friendly
interface for working with Python, especially for tasks like data analysis and
exploration.*

</details>

<details markdown="block">
  <summary> <sup><i>ipython magic commands</i></sup> </summary>

IPython magic commands are a set of special commands that can be used to
enhance the functionality of the IPython shell. These commands are prefixed
with the `%` character and can be used to perform various tasks such as modifying
the behavior of a code cell, simplifying common tasks like timing code
execution, profiling, etc. There are two types of magic commands: line magics
and cell magics.

- __Line Magics__
  - Line magics are used to provide a special functionality to a single line of code. They begin with the `%` character followed by the line command. Here are some examples of line magics:
    - `%autocall`:
      - This magic function makes a function automatically callable without having to use parentheses. It takes three possible mode parameters: 0 (off), 1 (smart), or 2 (always on).
    - `%automagic`:
      - This magic function makes magic functions callable without having to type the initial `%`. Without arguments, it toggles on/off. To deactivate, set to 0.
    - `%cd`:
      - This line magic changes the current directory. It automatically maintains an internal list of directories you visit during your IPython session, in the variable _dh. You can also use cd -<tab> to see directory history conveniently.
    - `%dhist`:
      - This magic command prints all directories you have visited in the current session. Every time `%cd` command is used, this list is updated in the _dh variable.
    - `%timeit`:
      - This function displays the time required by IPython to execute a Python expression. It can be used both as a line and cell magic.
    - `%who`:
      - This line magic prints all interactive variables, with some minimal formatting. If any arguments are given, only variables whose type matches one of these are printed.
- __Cell Magics__
  - Cell magics are used to provide a special functionality to multiple lines of code. They begin with the `%%` character followed by the cell command. Here are some examples of cell magics:
    - `%%time`:
      - This cell magic times the execution of a cell. It can be used to measure the time taken by a cell to execute.
    - `%%debug`:
      - This cell magic activates the interactive debugger.
    - `%%bookmark`:
      - This cell magic manages IPython's bookmark system. It can be used to set, list, or remove bookmarks.
- __Other Magic Commands__
  - `%run`:
    - This magic command runs the named file inside IPython as a program.
  - `%save`:
    - This magic command saves a set of lines or a macro to a given filename.
  - `%set_env`:
    - This magic command sets environment variables.
  - `%edit`:
    - This magic command calls upon the default text editor of the current operating system for editing a Python script. The script is executed as the editor is closed.
  - `%env`:
    - This magic command gets, sets, or lists environment variables.
- __Creating Custom Magic Commands__
  - IPython allows users to create and register their own custom magic commands. These custom magic commands can be used to provide special functionalities that are not available in the standard Python shell.

<!-- ipython magic commands -->
</details>

- `?` and `??`
  - Appending `?` to the end of a magic command and executing it will display its docstring.
  - Appending `??` will display its source code.

```
  ?         -> Introduction and overview of IPython's features (this screen).
  help      -> Access Python's own help system.
```

```ipython
%lsmagic - List currently available magic functions.
%save?
%lsmagic?

```

```python
# finding info about modules
#  or from bash console:
#  `$> pydoc math`
import math
print(dir(math))

help('math')
```


<details markdown="block">
  <summary> <sup><i>ipython extensions</i></sup> </summary>

- [ipychat](https://github.com/vinayak-mehta/ipychat): An AI extension for IPython that makes it work like Cursor
  - ipychat is an AI extension for IPython that helps you understand, debug, and write Python faster.
  - ...
    ```sh
    pip install ipychat
    ipychat config
    ```
    ```ipython
    %load_ext ipychat
    %ask what can I do with the cities dataframe
    %models
    ```

<!-- ipython extensions -->
</details>

<details markdown="block">
  <summary> <sup><i>troubleshooting</i></sup> </summary>

- `No module named 'sqlite3'`
  - on CentOS
    - ```sh
      yum install -y gcc openssl-devel bzip2-devel libffi-devel make sqlite-devel

      # it may need to reinstall Python again
      pyenv install -f -v 3.10.4
      ```

<!-- ipython troubleshooting -->
</details>


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

![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/a30a6667-0375-41ba-8972-73e6c5127934)

------ ------

[^1]: [Building a small REPL in Python](https://bernsteinbear.com/blog/simple-python-repl/)
