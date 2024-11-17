# AWK

Finds, filters, or replaces text in a file or from stdout.

```bash
awk <processing instruction string>
```

- Example:
  - This example pipes the string _"Bobby is cool"_ to the awk command. The `awk` command invokes the subcommand named `sub` to find any occurrence of "Bobby" and change the string to _"Teddy"_. Then, the subcommand print outputs the result of the substitution:
    ```bash
    echo "Bobby is cool" | awk '{sub("Bobby","Teddy"); print}'
    # Teddy is cool
    ```
  - This example uses `awk` to filter output according to field position. First, the example shows the output of the `who` command, which lists the current users logged in to the computer. The who command displays four fields (columns). The fields are username, the terminal line number, the login time, and the machine from which access originated.
  - The second execution of `who` pipes the result to `awk`. Then, awk uses the `print %1` subcommand set to show only the first field name. The third execution of `who` pipes the result to `awk`, which then filters input to print the values in the second field:

```bash
who
# jaggermick pts/0 2022-01-19 09:14 (192.168.86.28)
# guest pts/1 2022-01-19 10:07 (192.168.86.20)

who | awk '{print $1}'
# jaggermick
# guest

who | awk '{print $2}'
# pts/0
# pts/1
```
