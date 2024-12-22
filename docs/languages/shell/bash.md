---
layout: default
title: bash script
nav_order: 2
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/bash-script
---

Other things:
- [symbol & commands]({% link docs/languages/shell/bash.symbol-command.md %})
- [bash sample]({% link docs/languages/shell/bash.sample.md %})
  - _automatically run a command every n minutes_
- [file descriptor]({% link docs/languages/shell/bash.file-descriptor.md %})
- [further more...]({% link docs/languages/shell/bash.further-more.md %})

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

__What is Bash?__[^1]
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

- [more about functions]({% link docs/languages/shell/bash.symbol-command.md %}#functions)

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

<details markdown="block"> <summary> input/output redirection </summary>

```sh
cat << EOF
Hey Everyone
Hope your'e liking this guide
if so, feel free to share your thoughts
Thanks
EOF
```

--------
<!-- input/output redirection -->
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

<details markdown="block"> <summary> <strong>positional arguments</strong> </summary>

to take the positional Arguments with `$1` and `$2`:

```sh
#!/bin/bash
echo "First argument: $1"
echo "Second argument: $2"
```

run the executable with two arguments:

```sh
./positional_arguments.sh Arindam Magician
```

--------
<!-- positional arguments -->
</details>

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

## Control Flow Statements

<details markdown="block"><summary> <strong>if/else statements</strong> </summary>

```bash
if condition
then
  code_to_execute
fi

my_variable="Welcome to a new world!"
if [ "$my_variable" == "Welcome to a new world!" ]
then
  echo "This variable says Welcome to a new world!"
if
```

```sh
if (condition);then
  statement
elif(condition);then
  statement
else
  statement
fi
```

```sh
#!/bin/bash
if [[ "${1,,}" = "arindam" ]]; then
  echo "Hey Admin, Welcome to our Bash Tutorial."
elif [[ "${1,,}" = "magician" ]]; then
  echo "Hey Magician, You don't have Admin Access"
else
  echo "You Don't have Admin access, login as an Admin"
fi
```

--------
<!-- if statements -->
</details>

<details markdown="block"><summary> <strong>case statements</strong> </summary>

```sh
case EXPRESSION in
  PATTERN_1)
    STATEMENTS
    ;;
  PATTERN_2)
    STATEMENTS
    ;;
  PATTERN_N)
    STATEMENTS
    ;;
  *)
    STATEMENTS
    ;;
esac
```

```sh
#!/bin/bash
case ${1,,} in
  "arindam")
    echo "Hey Admin, Welcome to our Bash Tutorial."
    ;;
  "magician")
    echo "Hey Magician, You don't have Admin Access"
    ;;
  *)
    echo "You Don't have Admin access, login as an Admin"
    ;;
esac
```

--------
<!-- case statements -->
</details>


<details markdown="block"><summary> <strong>for loops</strong> </summary>

```bash
for variable_name in list
do
  code_to_execute
done

fruits=("apple" "banana" "orange")
for fruit in "${fruits[@]}"
do
  echo $fruit
done
```

--------
<!-- for loops -->
</details>

<details markdown="block"><summary> <strong>while loops</strong> </summary>

```bash
while condition
do
  code_to_execute
done

while [ "$my_variable" == "Hello, world!" ]
do
  echo "The variable is equal to Hello, world!"
  my_variable="Goodbye, world!"
done
```

--------
<!-- while loops -->
</details>

## Other Fundamental Concepts

<details markdown="block"><summary> <strong>string manipulation</strong> </summary>

<details markdown="block"><summary> AWK </summary> <a id="awk"></a>

AWK is to filter file contents or output of a command in such a way that we get the essential part of the output and the way we like it.

```sh
# this will print the whole content of the file
awk '{ print }' /path/to/file

# if we want to print specific text like the first/second word, then we can use `$1` and `$2`.
awk '{print $1}' temp.txt

# use different characters as separators
awk -F, {print $3} temp.csv
```
<!-- AWK -->
</details>

<details markdown="block"><summary> SED </summary> <a id="sed"></a>

SED is a command-line tool that lets us modify certain values in a text file using regular expressions.

```sh
sed 's/text_to_change/to_this_text/g' filename
```
<!-- SED -->
</details>

<!-- string manipulation -->
</details>

<details markdown="block"><summary> <strong>json manipulation</strong> </summary>

<details markdown="block"><summary> JQ </summary> <a id="jq"></a>

`jq` is a command line tool for working with JSON.

```sh
curl 'https://jsonplaceholder.typicode.com/users/1' | jq
```

<!-- SED -->
</details>
<!-- json manipulation -->
</details>

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

## Functions

```sh
function_name () {
  commands
}
```

```sh
#!/bin/bash
first_fn () {
   echo 'This is my first Bash function'
}
first_fn
```

```sh
#!/bin/bash
first_fn () {
   echo 'This is my first Bash function'
}
first_arg(){
   echo "this is the first argument: $1"
}
first_fn
first_arg Arindam
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


## Cron System [^2]

The `cron system` is a time-based job scheduling mechanism in Unix-like operating systems. It allows you to automate repetitive tasks by scheduling them to run at specific intervals. The term “cron” comes from the Greek word “chronos,” meaning time.

__Components of the Cron System__
- Cron Daemon (cron): The cron daemon is a background service that manages the scheduling and execution of cron jobs. It constantly checks the system's crontab files to determine when to run scheduled tasks.
- Cron Jobs: A cron job is a command or script that you want to run automatically at a specified time or interval. Each cron job consists of a command and a schedule defined by cron syntax.

Cron Syntax: Cron uses a specific syntax to define when a job should be executed. The syntax consists of five fields (plus an optional command field):
```
* * * * * command_to_run
| | | | |
| | | | +----- Day of the week (0 - 6) (Sunday = 0)
| | | +------- Month (1 - 12)
| | +--------- Day of the month (1 - 31)
| +----------- Hour (0 - 23)
+------------- Minute (0 - 59)
```

Cron Samples:
- run a script every night at midnight
  - `0 0 * * * /path/to/script.sh`
- send system statistics every hour
  - `0 * * * * /usr/bin/collect_stats.sh`

## Lynx [^1]

`lynx` is a text-based web browser that allows users to browse the internet and view web pages from the command line in Unix-like operating systems. It's a versatile tool for accessing information from websites without requiring a graphical user interface.

While it doesn’t render images or complex layouts like modern graphical browsers, it’s useful for quick access to text-based content, checking links, and accessing web resources in a terminal environment.

- basic usage
  - `lynx [options] [URL]`
- lynx samples
  - viewing a web page
    - `lynx https://www.example.com`
  - viewing page source
    - `lynx -source https://www.example.com`
  - viewing plain text version
    - `lynx -dump https://www.example.com`

## Miscellaneous for script

- escaping characters in bash (for JSON) [^3]
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
- [others]({% link docs/languages/shell/07-text-processing.md %}#miscellaneous-for-script)

----

[^1]: [Bash, the magic that empowers the Cloud](https://dev.to/sm0ke/bash-the-magic-that-empowers-the-cloud-5394)
[^2]: [Cron System by DeployPro](https://www.docs.deploypro.dev/script-languages/bash/cron-system)
[^3]: [Escaping characters in bash (for JSON) - _convert string to JSON_](https://stackoverflow.com/questions/10053678/escaping-characters-in-bash-for-json)
