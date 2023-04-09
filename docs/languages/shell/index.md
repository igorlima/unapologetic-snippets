---
layout: default
title: Shell Script
nav_order: 1
parent: Programming Languages
has_children: true
permalink: /docs/languages/shell
---

## Bash scripting syntax

setting variables
```sh
TEST_VAR='Hello World!'
export TEST_VAR='Hello World!'
```

accessing variable
```sh
echo $TEST_VAR
# if the variable is not specified, we can use the default value.
echo ${TEST_VAR:-hello world}
```

test command
```sh
# there are two syntaxes for using the test command
test EXPRESSION
[ EXPRESSION ]

# = and == are for string comparisons
# -eq is in the same family as -lt, -le, -gt, -ge, and -ne
if [ 1 -eq 1 ]; then echo 1; else echo 0; fi
if [ 1 -eq 0 ]; then echo 1; else echo 0; fi
if test 1 -eq 1; then echo 1; else echo 0; fi
if test 1 -eq 0; then echo 1; else echo 0; fi
```

exclamation mark
```sh
# exclamation mark in bash is often used to reverse the logic of the expression. if the expression is true, the ! operator will return false. and vice versa.
var="hello"
if [ ! "$var" = "hello" ]; then
  echo "var is not hello"
else
  echo "var is hello"
fi
```

bash variable $?
```sh
# '$?' is used to find the return value of the last executed command.
# if the last command was successful, then $? will be 0. if the previous command was unsuccessful, then $? will be non-zero.
ls xxx
if [ $? -eq 0 ]; then
  echo "The last command was successful."
else
  echo "The last command was unsuccessful."
fi
```

## Array

Using arrays in shell scripts

```sh
#!/bin/bash

fn() {
  numbers=(2 4 5 2)
  declare -a words

  # no space around the =
  words[0]='Orange'
  words[1]='Pineapple'

  echo ${numbers[@]} ${words[@]}
}

boo() {
  declare -a numbers=([0]="2" [1]="4" [2]="5" [3]="2" )
  declare -a words=([0]="Orange" [1]="Pineapple")
  echo ${numbers[@]} ${words[@]}
}

bar() {
  declare -a marks
  marks+=(75 65 80 102 26) # class A marks
  marks+=(103 68) # class B marks

  # remove invalid marks
  for i in "${!marks[@]}"; do
    if ((marks[i] > 100)); then
      unset "marks[$i]"
    fi
  done

  # prints marks
  echo ${marks[@]}

  # sort all marks
  marks_s=($(printf '%s\n' "${marks[@]}" | sort -nr))
  echo ${marks_s[@]}

  # prints the top-3
  echo ${marks_s[0]} ${marks_s[1]} ${marks_s[2]}
}

fn
boo
bar
```

## Map

Creating maps or dictionaries

Bash introduced the associative array feature for storing key-value data in version 4.0.

Use `!mapvar[@]` syntax to extract all dictionary keys as an array for iteration.

```sh
#!/bin/bash

fn() {
  # use `!mapvar[@]` syntax to extract all dictionary keys as an array for iteration.

  declare -A marks=([john]=75 [doe]=82 [ann]=83 [ava]=72)
  for key in "${!marks[@]}"; do
    printf "$key \t ${marks[$key]} \n"
  done
}

foo() {
  # `! -v` syntax as we typically use `not in` with Python dictionaries.

  read -p "Enter coords (i.e., [x]=10 [y]=12): " coords
  declare -A "coords=($coords)"

  if [ ! -v "coords[x]" ]; then
    coords[x]=5
  fi

  if [ ! -v "coords[y]" ]; then
    coords[y]=10
  fi

  for key in "${!coords[@]}"; do
    printf "$key = ${coords[$key]} \n"
  done
}

fn
foo
```

## Named arguments

Implementing named arguments support

```sh
#!/bin/bash

fn() {
  local title="Information"
  local message="Hello world"

  while getopts ":t:m:" option; do
    echo $option
    case "${option}" in
      t)
        title=${OPTARG}
        ;;
      m)
        message=${OPTARG}
        ;;
    esac
  done

  echo "$title $message"
}

fn $@
```

Bash also lets you create constants using the `declare -r` command. Whenever your script attempts to change a constant, Bash prints an error message on the screen.

## Pass-by-reference in functions

```sh
#!/bin/bash

function change_str_var() {
  # Using pass-by-reference in functions
  # If you use Bash 4.3 or a newer version, you can use --title name references to implement pass-by-reference in shell scripts.

  local str_new="Bash"
  local -n str_ref=$1
  echo "$str_ref -> $str_new" # Python -> Bash
  str_ref=$str_new
}

str="Python"
change_str_var str
echo $str
```
