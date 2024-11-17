# SED

Manipulates the content of a file or output sent to stdout.

```bash
sed [options] <manipulation instructions> <path/to/filename or stdout>
```

- Example:
  - The following example uses the `echo` command to send the string `Bobby is cool` to the `sed` command. The command `sed` uses the s subcommand to substitute the name `Teddy` where the name `Bobby` occurs. The output is then displayed:
    ```bash
    echo Bobby is cool | sed 's/Bobby/Teddy/'
    # Teddy is cool
    ```
