---
layout: default
title: Bash - Further More
nav_exclude: true
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/bash-further-more
---

# Bash Samples

__[back]({% link docs/languages/shell/bash.md %})__

<br/>
<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>


## `echo` vs `printf`

Both `echo` and `printf` are commands used for displaying text in the shell but they have some important differences.

While echo simply outputs the strings provided to it, `printf` allows for more __detailed formatting__ of strings, numbers, and other data types.

```bash
name="Ashish"
age=28

printf "%s is %d years old.\n" "$name" "$age"

# Output: Ashish is 28 years old.
```

## `which` vs `whereis`

The `which` command is used to find the location of the executable that would be run if you type the command in the shell.

The command returns looks in directories listed in the `$PATH` environment variable and returns the path to the executable.

```bash
which ls

# Output: /bin/ls
```

On the other hand, the `whereis` command provides more detailed information about a command _which includes_ __its binary, source files, and manual pages__.

It searches in standard system directories and does not rely on the `$PATH` variable.

```bash
whereis ls

# Output: ls: /bin/ls /usr/share/man/man1/ls.1.gz
```

The output tells us where the executable binary for the `ls` command is located (`/bin/ls`) and where its associated manual/ man page is (`/usr/share/man/man1/ls.1.gz`).

----

[^1]: [...](https://www.google.com)

