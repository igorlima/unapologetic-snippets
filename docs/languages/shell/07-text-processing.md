---
layout: default
title: advanced text processing
nav_order: 7
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/text-processing
---

# Advanced Text Processing

- [sed]({% link docs/languages/shell/08-sed.md %})
- [awk]({% link docs/languages/shell/08-awk.md %})

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

# Miscellaneous for script

- escaping characters in bash (for JSON) [^1]
  - using `jq`
    - sample 1:
      - `echo -n '猫に小判' | jq -Rsa .`
        - _outputs_: `"\u732b\u306b\u5c0f\u5224"`
    - sample 2:
      - `cat tmp.txt | jq -Rsa .`
    - to explain:
      - `-R` means "raw input"
      - `-s` means "include linebreaks" (mnemonic: "slurp")
      - `-a` means "ascii output" (optional)
      - `.` means "output the root of the JSON document"
- [others]({% link docs/languages/shell/bash.md %}#miscellaneous-for-script)

----

[^1]: [Escaping characters in bash (for JSON) - _convert string to JSON_](https://stackoverflow.com/questions/10053678/escaping-characters-in-bash-for-json)
