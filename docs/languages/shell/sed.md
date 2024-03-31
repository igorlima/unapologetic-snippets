---
layout: default
title: sed
nav_order: 3
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/sed
---

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

- [more about advanced text processing]({% link docs/languages/shell/text-processing.md %})


Sed stands for Stream Editor. It is very powerful tool to manipulate, filter, and transform text.

```sh
# SAMPLE
# the below employee database contains
# the following fields for every record:
# • Employee Id
# • Employee Name
# • Title


cat << "EOF" > employee.txt
101,John Doe,CEO
102,Jason Smith,IT Manager
103,Raj Reddy,Sysadmin
104,Anand Ram,Developer
105,Jane Miller,Sales Manager
EOF

vi employee.txt
```

## Command Syntax

__Basic sed syntax:__

Sed reads one line at a time from the `{input-file}` and executes the `{sed-commands}` on that particular line.
```sh
sed [options] {sed-commands} {input-file}
```

__Basic sed syntax for use with sed-command file:__
```sh
sed [options] -f {sed-commands-in-a-file} {input-file}
```
```sh
cat << "EOF" > test-script.sed
/^root/ p
/^nobody/ p
EOF

sed -n -f test-script.sed /etc/passwd
```

__Basic sed syntax using -e:__
```sh
sed [options] -e {sed-command-1} -e {sed-command-2} {input-file}
```
```sh
sed -n -e '/^root/ p' -e '/^nobody/ p' /etc/passwd
```

__Basic sed syntax using { }:__
```sh
sed [options] '{
sed-command-1
sed-command-2
}' input-file
```
```sh
sed -n '{
/^root/ p
/^nobody/ p
}' /etc/passwd
```

__basic sed alternative__
```sh
sed -i -e 's|^that starts with|the replace|' filename.ext
sed -i -e 's|^#precedence ::ffff:0:0/96  100|precedence ::ffff:0:0/96  100|' /etc/gai.conf
```

## Sed Scripting Flow

Sed will:
- __Read__ a line into the pattern space (an internal temporary sed buffer, where it places the line it reads from the input file).
- __Execute__ the sed command on the line in the sed pattern space. If there are more than one sed commands available, either via a sed script, -e options, or { }, it executes all the sed commands one by one in sequence on the line that is currently in the pattern space.
- __Print__ the line from the pattern space. After printing this line, the sed pattern space will be empty.
- __Repeat__ this again until the end of the input file is reached.


## other resources

- [more about advanced text processing]({% link docs/languages/shell/text-processing.md %})

