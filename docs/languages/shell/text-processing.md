---
layout: default
title: advanced text processing
nav_order: 1
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/text-processing
---

# Advanced Text Processing

- [sed]({% link docs/languages/shell/sed.md %})

```sh
sed 's/[old_string]/[new_string]/g' [file_name] # replaces all occurrences of a string in a file.
awk '/[pattern]/ {action}' [file_name] # scans a file line by line and performs actions on lines that match a pattern.
tail -n [number] [file_name] # displays the last few lines of a file.
fold -w [width] [file_name] # wraps each input line to fit in specified width.
tac [file_name] # concatenate and print files in reverse.
sed -n [pattern]p [file_name] # prints only those lines that match the pattern.
awk 'BEGIN {action}' [file_name] # executes an action before reading the lines of a file.
```

Find any lines exceeding a certain length:
```sh
# find any lines exceeding a certain length
grep '.\{80\}' file
perl -nle 'print if length$_>79' file
awk 'length>79' file
sed -n '/.\{80\}/p' file
```
