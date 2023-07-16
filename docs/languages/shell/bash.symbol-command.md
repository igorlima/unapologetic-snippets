---
layout: default
title: Bash - Symbol & Command
nav_exclude: true
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/bash-symbol-command
---

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

## Redirection

One of the most basic but powerful features of Bash is the ability to redirect the input and output of commands. Here are some of the symbols you can use for redirection: [^1]

- `>`: Redirects the output of a command to a file, overwriting its contents if the file already exists. For example, `ls > file.txt` will write the output of the `ls` command to a file named `file.txt`.
- `>>`: Redirects the output of a command to a file, appending it to the end of the file if it already exists. For example, `echo "Hello, world!" >> file.txt` will append the string "Hello, world!" to the end of the file.
- `<`: Redirects the input of a command to come from a file instead of the keyboard. For example, `sort < file.txt` will sort the contents of `file.txt`.
- `2>`: Used to redirect the error output of a command to a file.
- `2>>`: Used to append the error output of a command to a file.
- `&>`: Used to redirect both the output and error output of a command to a file.
- `&>>`: Used to append both the output and error output of a command to a file.

## Wildcards and Globbing

Bash supports a powerful feature called globbing, which allows you to use wildcards to match patterns of filenames. Here are some of the most common symbols used in globbing:

- `*`: Matches any number of characters in a filename. For example, `ls *.txt` will list all files in the current directory that end with `.txt`.
- `?`: Matches any single character in a filename. For example, `ls file?.txt` will list all files in the current directory that start with `file` and have exactly one character after it, followed by `.txt`.
- `[ ]`: Matches any one of a set of characters. For example, `ls file[123].txt` will list all files in the current directory that start with `file` and end with either `1`, `2`, or `3`, followed by `.txt`.


## Regular Expressions

- `grep`: Used to search for a pattern in a file or stream.
- `sed`: Used to perform text transformations on a file or stream.
- `awk`: Used for pattern scanning and processing.


## Variable Expansion and Environment Variables

- `$`: Used to reference the value of a variable.
- `export`: Used to create or modify an environment variable in Bash.
- `${}`: Used to perform variable expansion.
- `$(())`: Used to perform arithmetic expansion.
- `env`: Used to display the current environment variables.
- `echo`: Used to display text or values of environment variables.

## Wildcards and Globbing

- `*`: Matches any number of characters in a filename.
- `?`: Matches any single character in a filename.

## Arithmetic Operations

- `+`: Addition operator in arithmetic expressions. For example, `echo $((2+3))` will output `5`.
- `-`: Subtraction operator in arithmetic expressions. For example, `echo $((10-5))` will output `5`.
- `*`: Multiplication operator in arithmetic expressions. For example, `echo $((4*3))` will output `12`.
- `/`: Division operator in arithmetic expressions. For example, `echo $((20/5))` will output `4`.
- `$((expression))`: Used to perform arithmetic operations in Bash. For example, `echo $((2 + 3))` will print the value `5`.
- `expr`: Used to perform arithmetic operations.

## Comparison

- `==`: Used to compare two values for equality in a conditional expression. For example, `if [ "$name" == "John" ]; then echo "Hello, John!"; fi` will output "Hello, John!" if the variable `name` has the value "John".
- `!=`: Used to compare two values for inequality in a conditional expression. For example, `if [ "$name" != "John" ]; then echo "You're not John!"; fi` will output "You're not John!" if the variable `name` does not have the value "John".

## Conditional Execution

- `&&`: Used to execute a command only if the previous command was successful. For example, `command1 && command2` will execute `command2` only if `command1` was successful.
- `||`: Used to execute a command only if the previous command failed. For example, `command1 || command2` will execute `command2` only if `command1` failed.

## Conditional Statements

- `if`: Used to check if a condition is true and execute commands accordingly. For example, `if [[ $foo == "bar" ]]; then echo "foo is bar"; fi` will check if the variable `foo` is equal to "bar" and print "foo is bar" if true.
- `elif`: Used to check additional conditions if the previous conditions in an `if` statement are false. For example, `if [[ $foo == "bar" ]]; then echo "foo is bar"; elif [[ $foo == "baz" ]]; then echo "foo is baz"; else echo "foo is something else"; fi` will check if `foo` is equal to "bar" and print "foo is bar" if true, otherwise check if `foo` is equal to "baz" and print "foo is baz" if true, and finally print "foo is something else" if both previous conditions are false.
- `[[ ]]`: Used to perform conditional tests in Bash.
- `case`: Used to match a variable against a list of patterns and execute commands accordingly. For example,

```bash
case $fruit in
  apple)
    echo "It's an apple"
    ;;
  banana)
    echo "It's a banana"
    ;;
  *)
    echo "It's something else"
    ;;
esac
```

will check the value of the variable `fruit` and print "It's an apple" if it equals "apple", "It's a banana" if it equals "banana", or "It's something else" for any other value.

## Loops

- `for`: Used to create a loop that iterates over a list of items. For example, `for fruit in apple banana orange; do echo $fruit; done` will output "apple", "banana", and "orange" on separate lines.
- `while`: Used to create a loop that continues executing commands as long as a condition is true. For example, `while true; do echo "Hello"; sleep 1; done` will output "Hello" repeatedly until the script is interrupted.
- `until`: Used to execute a command until a condition is true.

## Functions

- `function`: Used to define a function in Bash. For example, `function hello { echo "Hello, $1!"; }` defines a function called `hello` that takes one argument and prints "Hello, " followed by the argument.
- `$1`, `$2`, etc.: Used to refer to the first, second, third, etc. arguments passed to a function. For example, in the function `hello` defined above, `$1` refers to the first argument passed to the function.

## Arrays

- `[]`: Used to define an array in Bash. For example, `my_array=(apple banana orange)` defines an array called `my_array` containing the values "apple", "banana", and "orange".
- `${array[@]}`: Used to refer to all the elements in an array in Bash. For example, `echo ${my_array[@]}` will output "apple banana orange" on a single line.

## Job Control

- `jobs`: Used to display a list of background jobs.
- `fg`: Used to bring a background job to the foreground.
- `bg`: Used to move a suspended or stopped job to the background.
- `kill`: Used to terminate a running process.

----

[^1]: [Powering Up Your Command Line Skills: A Complete Guide to Bash and Linux Symbols and Commands](https://blog.devops.dev/powering-up-your-command-line-skills-a-complete-guide-to-bash-and-linux-symbols-and-commands-a68b5e077e0)

