# CHMOD

```bash
chmod 664 file_name
chmod u+w,go-w file_name
```

There are 3 types of permissions for a file: [^1]
- read
- write
- execute

On the other hand, there are 3 types of users:
- the owner of the file
- the group (including all its members) which owns the file
- other users (who are not the file’s owner and not members of the group)

Therefore, file permissions management in Linux is just about how to assign the 3 types of permissions to the 3 types of users.

Permission situation:
- `0` for no permission
- `1` for “execute”
- `2` for “write”
- `4` for “read”

So, as the following table shows, there are 8 possible situations in total:

| #   | Sum                | rwx   | Permission              |
| --- | -----              | ----- | ------------            |
| 7   | 4(r) + 2(w) + 1(x) | rwx   | read, write and execute |
| 6   | 4(r) + 2(w)        | rw-   | read and write          |
| 5   | 4(r)        + 1(x) | r-x   | read and execute        |
| 4   | 4(r)               | r--   | read only               |
| 3   | 2(w) + 1(x)        | -wx   | write and execute       |
| 2   | 2(w)               | -w-   | write only              |
| 1   |               1(x) | --x   | execute only            |
| 0   | 0                  | ---   | no permission           |

- `chmod 664 file_name`
  - the first `6` means the __owner__ of the file has the “read and write” permission.
  - the second `6` means the owning __group__ has the “read and write” permission.
  - the last `4` means the __other__ users only have the “read” permission.

As its name implies, the chmod command can use some symbols to represent different types of users and permissions.
- For users,
  - `u`: file owner
  - `g`: owning group
  - `o`: other users
  - `a`: all of the above (the same as ugo)
- For permissions,
  - `r`: read
  - `w`: write
  - `x`: execute
In addition, there are 3 types of operators:
  - `=`: to assign the exact permissions
  - `-`: to remove permissions
  - `+`: to add permissions

- `chmod u+w,go-w file_name`
  - it means giving the owner “write” permission and removing the “write” permission from the group and other users.
    - `u+w`: add the “write” permission to the owner of the file
    - `go-w`: remove the “write” permission from the owning group and other users




# REFERENCE
- [^1]: [5 Levels of Using Chmod Commands in Linux](https://medium.com/techtofreedom/5-levels-of-using-chmod-commands-in-linux-4a8431bf2a61)
