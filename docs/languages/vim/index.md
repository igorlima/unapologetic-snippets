---
layout: default
title: Vim
nav_order: 2
parent: Programming Languages
has_children: true
permalink: /docs/languages/vim
---


__How to use VI to search for lines with a certain length?__

- [find any lines exceeding a certain length]({% link docs/languages/shell/text-processing.md %})

how to use VI to find lines longer than a certain number of characters:
```vim
/\%>80v.\+
```

to find lines shorter than 25 characters:
```vim
/^.\{,25}$/
```
