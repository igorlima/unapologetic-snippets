# Special Variables in Bash Scripting

## List of special shell variables

| Variable | Description                                                                           |
| ---      | ---                                                                                   |
| `$0`     | The filename of the current script.                                                   |
| `$1-$9`  | Stores the names of the first 9 arguments or can be used as the arguments’ positions. |
| `$$`     | The process id of the current shell.                                                  |
| `$#`     | The number of arguments supplied to a script.                                         |
| `$*`     | Stores all the command line arguments by joining them together.                       |
| `$@`     | Stores the list of arguments as an array.                                             |
| `$?`     | Specifies the exit status of the last command or the most recent execution process.   |
| `$!`     | Shows the id of the last background command.                                          |
| `$-`     | Shows the options set used in the current bash shell.                                 |

__Sample Script__ [^1]

```bash
#!/bin/bash

# sample.sh

echo "Process id of shell = $$"
echo "Name of shell script = $0"
echo "Number of args = $#"
echo "Argument 1 = $1"
echo "Argument 2 = $2"

# $ ./sample.sh
# $ ./sample.sh 1 2 3
```

- `$?`: the `$?` represents the exit status of the previous command.
- `$!`: it returns the process id of the last background command.
- `$*` and `$@`
  - `$*` and `$@` special variables are used to take multiple arguments at a time. But there is some difference between them which is shown as:
    1. `$*` and `$@` when unquoted are identical and expand into the arguments.
    2. `$*` is a single word, comprising all the arguments to the shell, joined together with spaces. For example ‘1 2’ 3 becomes “1 2 3”.
    3. `$@` is identical to the arguments received by the shell, the resulting list of words completely matches what was given to the shell. For example ‘1 2’ 3 becomes “1 2” “3”.
    4. In short, `$@` when quoted (`$@`) breaks up the arguments if there are spaces in them. But `$*` does not break the arguments.
  - Let’s take an example to explain the above points.
    ```bash
    #!/bin/bash

    # variable arguments
    # varArgs.sh

    echo "argument length=${#@}"
    echo "argument length=${#*}"
    echo
    for arg in $*
    do
      echo "$arg "
    done
    echo
    for arg in "$*"
    do
      echo "$arg "
    done
    echo
    for arg in $@
    do
      echo "$arg "
    done
    echo
    for arg in "$@"
    do
      echo "$arg "
    done
    ```
    ```bash
    ./varArgs.sh 1 2 3 4 5
    ./varArgs.sh 1 2 3 '4 5'

    ```
- `$-`: (dollar hyphen) is a special variable that returns the flags used in the current Bash shell. $- contains the shell’s flags in use in the terminal.


# REFERENCES

- [^1]: [Special Variables in Bash Scripting - Knoldus Blogs](https://blog.knoldus.com/special-variables-in-bash-scripting/)
- [^2]: [Special Variables in Bash Scripting - Medium](https://medium.com/@knoldus/special-variables-in-bash-scripting-56cbd660714d)
