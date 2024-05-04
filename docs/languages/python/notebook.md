---
layout: default
title: Notebook
nav_order: 4
parent: Python
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/python/notebbok
---

# Notebook

- [What is a data science Notebook?]({% link docs/algorithms-data-structures/index.md %}#data-science-notebook)

```sh
# to install
pip3 install jupyter

# to run
# by default, run on: http://localhost:8888
jupyter notebook
```

__virtual env__

```sh
python3 -m venv my-jupyter-env       # to create a virtualenv
source my-juptyer-env/bin/activate   # to activate the virtualenv
deactivate                           # to leave the virtualenv
```

__Jupyter in the terminal__

```sh
# Jupyter in the terminal.
# https://github.com/davidbrochart/jpterm
# Key bindings
# https://github.com/davidbrochart/jpterm/blob/10d9e93105ab80053314944e805689a5c9a01449/docs/plugins/notebook_editor.md

# enter: enter the edit mode, allowing to type into the cell.
# esc: exit the edit mode and enter the command mode.

# to install
pip3 install jpterm

# to install a Python kernel
pip3 install ipykernel

# REFERENCE
# The Predecessor:
# Jupyter Notebooks in the terminal.
#  - https://github.com/davidbrochart/nbterm
#  - https://blog.jupyter.org/nbterm-jupyter-notebooks-in-the-terminal-6a2b55d08b70
#  - `pip3 install nbterm`
#   enter: enter the edit mode, allowing to type into the cell.
#   esc: exit the edit mode and enter the command mode.
#   ctrl-h: show help.
```

<details markdown="block">
  <summary>
    <i>sample notebook</i>
  </summary>

```json
{
  "cells": [
    {
      "source": "This is a raw cell",
      "cell_type": "raw",
      "metadata": {}
    },
    {
      "cell_type": "markdown",
      "metadata": {},
      "source": "This is a markdown cell"
    },
    {
      "execution_count": 1,
      "cell_type": "code",
      "source": "a = 3\nprint(a+1)\n",
      "outputs": [],
      "metadata": {}
    }
  ],
  "metadata": {
    "kernelspec": {
      "language": "python",
      "name": "python3",
      "display_name": "Python 3"
    },
    "language_info": {
      "version": "3.9.2",
      "mimetype": "text/x-python",
      "name": "python",
      "file_extension": ".py"
    }
  },
  "nbformat": 4,
  "nbformat_minor": 4
}
```

-----
<!-- sample notebook -->
</details>

------ ------

[^1]: [...](...)
