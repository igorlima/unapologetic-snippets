---
layout: default
title: Bash Script
nav_order: 2
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/bash-script
---

Other things:
- [symbol & commands]({% link docs/languages/shell/bash.symbol-command.md %})

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

## Bash scripting syntax

<details markdown="block">
  <summary>
    getting started...
  </summary>

__What is Bash?__
Bash, short for _“Bourne Again Shell”_ is a command-line interface and scripting language used in Unix-like operating systems. It provides a way to interact with the operating system by running commands, and it also supports writing scripts to automate tasks.

__variables__
Bash uses variables to store and manipulate data.
```bash
name="John"
echo "Hello, $name!"
```

__input and output__
- echo : Print text to the screen.
- read : Read input from the user.
```bash
echo "What's your name?"
read name
echo "Hello, $name!"
```

__conditional statements__
Bash supports `if` statements for conditional execution.
_[See]({% link docs/languages/shell/bash.symbol-command.md %}#use-of-double-brackets) a difference between double brackets and when a condition is enclosed in a single bracket._
```bash
age=25
if [ "$age" -lt 18 ]; then
  echo "You're a minor."
else
  echo "You're an adult."
fi
```

__loops__
```bash
for fruit in apple banana orange; do
  echo "I like $fruit."
done
```

__functions__
```bash
greet() {
  echo "Hello, $1!"
}
greet "Alice"
```

__simple calculator script__
```bash
#!/bin/bash
read -p "Enter first number: " num1
read -p "Enter second number: " num2
sum=$((num1 + num2))
echo "Sum: $sum"
```

-----
<!-- getting started -->
</details>

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

## Strict Mode

Strict mode is a set of options that can be activated to improve the safety and dependability of the script.

The following is an explanation of the choices made in stringent mode:
1. The `-e` option means to “exit immediately if any command fails.” When enabled, if any command within it exits with a non-zero status, denoting an error or failure, the script will end immediately. This aids in spotting mistakes at an early stage and stops the script from running with unexpected or incorrect data.
2. The `-u` option interprets unset variables as an error. The script will end in error if any variable is utilized before receiving a value. This lowers the possibility of defects and unexpected behavior by identifying potential problems brought on by the use of uninitialized variables.
3. The `-o pipefail` option also takes pipeline failures into account. Typically, pipelines in Bash scripts report the exit status of the final command in the pipeline (commands connected using pipes are denoted by the symbol `|` ). When `pipefail` is enabled, the pipeline is deemed to have failed even if the last command succeeds if any command in the pipeline fails (returns a non-zero status). This makes sure that pipeline faults are handled correctly and are not disregarded.

```bash
#!/bin/bash
set -euo pipefail

# -e: Exit immediately if any command fails
# -u: Treat unset variables as an error
# -o pipefail: Consider pipeline failures as well
```

### Command failure

```bash
if ! command; then
  echo "Command failed." >&2
  exit 1
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

