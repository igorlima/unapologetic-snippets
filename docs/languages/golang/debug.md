---
layout: default
title: Debug
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/debug
---

<br/>
<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

## debuging go with `dlv`

```
 (dlv) break main.main
 (dlv) break ./main.go:10
 (dlv) condition <breakpoint name or id> <boolean expression>
 (dlv) list ./main.go:10
 (dlv) breakpoints
 (dlv) clear ./main.go:10
 (dlv) clear <breakpoint name or id>
 (dlv) clearall
 (dlv) clearall main.main
 (dlv) locals
 (dlv) whatis c
 (dlv) stack
```

to execute a file containing a list of delve commands:
- `(dlv) source dlv.cmd`

<details markdown="block">
  <summary> quick win... </summary>

Delve (`dlv`) is a CLI-based debugger for Go, tailored to the language’s concurrency model and runtime. It allows you to set breakpoints, inspect goroutines, and evaluate and manipulate variables in real-time.

- [quick tips](#dlv-tips)
- [keyboard shortcuts](#dlv-keyboard-shortcuts)

__Side Note__: Breakpoints `runtime-fatal-throw` and `unrecoverable-panic` do exactly what they sound like. Delve includes breakpoints on these critical failure points to debug a fatal runtime throw or an unrecoverable panic easily.

### dlv tips

| Task                       | Go Command                     | Delve Command                    |
| ----                       | ----------                     | -------------                    |
| Run Program                | `go run main.go`               | `dlv debug`                      |
| Run Tsts                   | `go test`                      | `dlv test`                       |
| Compile and Run Executable | `go build -o myapp && ./myapp` | `dlv exec ./myapp`               |
| Run Specific Test          | `go test -run TestName`        | `dlv test -- -test.run TestName` |

### dlv keyboard shortcuts

| Command Name | Shortcut | Description                                                                   |
| ------------ | -------- | -----------                                                                   |
| `break`      | `b`      | Sets a breakpoint at a specific source location.                              |
| `continue`   | `c`      | Continues execution until the next breakpoint.                                |
| `step`       | `s`      | Executes the current line and stops at the next line, entering functions.     |
| `next`       | `n`      | Executes the current line and stops at the next line in the current function. |
| `print`      | `p`      | Evaluates an expression and prints its value.                                 |
| `list`       | `l`      | Shows the source code around the current line.                                |
| `set`        |          | Changes the value of a variable.                                              |
| `locals`     |          | Prints all local variables in the current scope.                              |
| `clear`      |          | Removes a breakpoint.                                                         |
| `exit`       |          | Exits the debugger.                                                           |

-----
<!-- delve - quick win -->
</details>

<details markdown="block">
  <summary> more... </summary>
  
- [using breaking points in delve](#using-breaking-points-in-delve)
- [doing basic analysis with delve](#doing-basic-analysis-with-delve)
- [location specifiers in delve](#location-specifiers-in-delve)
- [delve config](#delve-config)

### using breaking points in delve
- set a breakpoint on the current line: `break +0` (abbr `b +0`)
- set a breakpoint on `line_number` in the current file:
  - `break <line_number> (abbr b <line_number>)`
- set a breakpoint on `line_number` in `file_name`:
  - `break <file_name>:<line_number>` (abbr `b <file_name>:<line_number>`)
- list all breakpoints: `breakpoints` (abbr `bp`)
- remove a breakpoint: `clear <breakpoint_id>`
- remove all breakpoints: `clearall`

### doing basic analysis with delve
- list the current source and `line_count` before and after: `list` (abbr `l`)
- run an expression in the current context: `print <expr>` (abbr `p <expr>`)
- show the current stacktrace: `stack` (abbr `bt`)
- list all local variables: `locals`
- search all local variables: `locals <search_regex>`
- show detailed information for a local var: `locals -v <search_regex>`
- print local variable data:
  - `print <local_variable_name>` (abbr `p <local_variable_name>`)

### location specifiers in delve
- `*<address>` specifies the location of memory address address. address can
  be specified as a decimal, hexadecimal or octal number
- `<filename>:<line>` specifies the line line in filename. filename can be the
  partial path to a file or even just the base name as long as the expression
  remains unambiguous.
- `<line>` specifies the line line in the current file
- `+<offset>` specifies the line offset lines after the current one
- `-<offset>` specifies the line offset lines before the current one
- `<function>[:<line>]` specifies the line line inside function. The full
  syntax for function is `<package>.(*<receiver type>).<function name>` however
  the only required element is the function name, everything else can be omitted
  as long as the expression remains unambiguous. For setting a breakpoint on an
  init function (ex: main.init), the `<filename>:<line>` syntax should be used
  to break in the correct init function at the correct location.
- `/<regex>/` specifies the location of all the functions matching regex


### delve config

```
 (dlv) config -list
 (dlv) config max-string-len 1000
 (dlv) config max-array-values 1000
 (dlv) config source-list-line-count 10
 (dlv) config -save
```

the config file on MacOS is by default in this following location:
- `cat $HOME/.dlv/config.yml`

-----
<!-- delve - more -->
</details>


## debuging go with `vim-go`

- basic
  - press `K` or `:GoDoc` or `:GoDocBrowser`
  - to trigger completion, while in INSERT mode
    - type `<ctrl+x><ctrl+o>`
  - `<leader>i` or `:GoInfo`
    - see function signature
    - prints the description of the identifier under the cursor
  - `ctrl+]`/`gd` or `:GoDef`
    - jump to a definition, locally or globally
    - jump to the type declaration
  - `ctrl+o` or `:GoDefType`
    - jump to a type definition from an instance of the type
  - `ctrl+t` or `GoDefPop`
    - jump back to the previous location
  - <details markdown="block"><summary>more...</summary>

    - `:GoAlternate`
      - switch between the file and test file
      - use `!` (bang) to create a file if not exists yet
    - `:GoDefStack`
      - see history of all your locations invoked via `:GoDef`
      - `:GoDefStackClear`
    - `:GoReferrers`
      - finds references to the selected identifier, scanning all necessary packages within the workspace
    - `:GoCallers`
      - see the callers of a given function
    </details>
- remote debugging:
  - `go build -gcflags "all=-N -l" main.2023a02m21d.go`
  - ```sh
    dlv --listen=0.0.0.0:8181 --headless=true --api-version=2 --accept-multiclient exec ./main.2023a02m21d --log
    ```
  - `:GoDebugConnect 0.0.0.0:8181`
    - `dlv connect 0.0.0.0:8181`
- to download Golang Makefile
  - ```sh
    (GIST_ID=e3be354c2c3aa78a2f7956d046b0bf3a && \
     curl -L \
     'https://gist.githubusercontent.com/igorlima/'$GIST_ID'/raw/Makefile' \
     -o Makefile)
    ```

<details markdown="block">
  <summary> more... </summary>

- [using vim-go while coding](#using-vim-go-while-coding)
- [debug commands](#debug-commands-in-vim-go)
- [other `vim-go` functions](#miscellaneous-vim-go-functions)
- [how-to](#how-to-session-with-vim-go)

### using vim-go while coding
- press `K` or `:GoDoc` or `:GoDocBrowser`
- to trigger completion, while in INSERT mode
   - type `<ctrl+x><ctrl+o>`
- supports browsing tags of the current file and overview its structure
   - `:TagbarToggle`
- `<leader>i` or `:GoInfo`
   - see function signature
   - prints the description of the identifier under the cursor
- `ctrl+]`/`gd` or `:GoDef`
   - jump to a definition, locally or globally
   - jump to the type declaration
- `ctrl+o` or `:GoDefType`
   - jump to a type definition from an instance of the type
- `ctrl+t` or `GoDefPop`
   - jump back to the previous location
- `:GoRun %`
   - run the current file
   - `:!clear` to clear output
- `:GoRun`
   - for whole package
- `:GoBuild`
   - jumps to the first error encountered
   - the quickfix view will open if any erro compilation appears
   - `:GoBuild!`
      - append the `!` (bang) if don't want to jump to the first error encountered
- `:call go#auto#template_autocreate()`
   - `:call go#template#create()`

### debug commands in vim-go

- intro
  - `pgrep dlv` a command to list all delve process
  - `lsof -i :8081` lists all process running at the given port
  - resize window debug layout size
    - `:vertical resize 35` - width
    - `:resize 40` - height
  - `runtime.Breakpoint()`
    - `import "runtime"`
- `<Leader>gds`  or `:GoDebugStart`
- `<Leader>gdc`  or `:GoDebugConnect 0.0.0.0:8181`
- `<Leader>gdss` or `:GoDebugStart main.go`
   - the command below `gdss` is so helpfull when there is a need to run a
     single file
- `<Leader>gdb`  or `:GoDebugBreakpoint`
- `<Leader>gn`  or `:GoDebugNext`
- `<Leader>gc`  or `:GoDebugContinue`
- `<Leader>gi`  or `:GoDebugStep`
   - to step in
- `<Leader>go`  or `:GoDebugStepOut`
- `<Leader>gp`  or `:GoDebugPrint`
   - print value should be in the register `@"`
- `<Leader>gP`  or `:GoDebugPrint call`
   - print value should be in the register `@"`
      - ```
        " REDIRECT/COPY OUTPUT OF A COMMAND TO A REGISTER
        :redir @*
        :redir END
        " SAMPLE COMMAND
        :redir @* | set guifont | redir END

        " `:redir` command redirects the output of a command to a register `(@*)`.
        " the register `@*` refers to the clipboard.
        ```
- `<Leader>gdS`  or `:GoDebugStop`
- `<Leader>gdH`  or `:GoDebugHalt`
- `:GoDebugSet {var} {value}`
   - set the variable `{var}` to `{value}`
      - `:GoDebugSet truth 42`
      - this only works for `float`, `int` and variants, `uint` and variants,
        `bool`, and pointers

### miscellaneous vim-go functions

- `:GoInstall`
- `:GoTest`
   - a quickfix will open if any error encounters
- `:GoCoverage`
- `:GoCoverageClear`
- `:GoCoverageToggle`
- `:GoFmt`
- `:GoLint`
- `:GoDescribe`
   - shows various properties of the selected syntax
- `:GoAlternate`
   - switch between the file and test file
   - use `!` (bang) to create a file if not exists yet
- `:GoDefStack`
   - see history of all your locations invoked via `:GoDef`
   - `:GoDefStackClear`
- `:GoReferrers`
   - finds references to the selected identifier, scanning all necessary
     packages within the workspace
- `:GoCallers`
   - see the callers of a given function
      - ways to edit a filename under the cursor
         - `gf` - edit existing file under cursor in same window
         - `C-W gf` - edit existing file under cursor in new tabpage
         - `C-W f` - edit existing file under cursor in split window
         - `C-W C-F` - edit existing file under cursor in split window
      - useful vim commnad
         - location file syntax
            - `:set syntax?`
               - `:set filetype?`
            - `:set syntax=qf`
               - `set filetype=qf`
         - move a window to a new tab
            - `ctrl-w` `T`
            - `ctrl-w` `Shift-T`
         - list buffers `:ls`
            - open buffer `:b1` - `b <buffer_number>`
            - `:tabnew | b 2`
            - `:tabnew | b2`
         - quickfix list
            - the location list behaves just like the quickfix list except
              that it is local to the current window instead of being global
              to the Vim session. So if you have five open windows, you can
              have up to five location lists, but only one quickfix list.
               - same alias for navigation list as well~
                 rather `q` uses `l` instead
            - `<leader>qX` clear the quickfix list
            - `<leader>qx` add a formated line in quickfix list
            - `<leader>qa` add a line in quickfix list described with the yank
               register
            - `<leader>qs` set the quickfix syntax
            - navigation list: `[[q`, `]]q`, `[q`, `[Q`, `]q`, `]Q`
               - `:copen`, `:lopen`
            - same alias for navigation list as well~
              rather `q` uses `l` instead
- `:GoCallstack`
   - shows an arbitrary path from the root of the call graph to the function
     containing the selection
- `:GoRename`
   - searches all packages under GOPATH and renames all identifiers that depend
     on the identifier
   - `:GoFreevars`
      - open a quickfix list with all the variables that are free variables
      - can also be used to understand the complexibility of a code. Just run
        it and see how many variables are dependent on it
- `:GoAddTags`
   - `:GoRemoveTags`
   - add or remove tags on struct fields


### how to session with vim-go

- refresh/update the window layout during the debug
   - set any variable by using `:GoDebugSet`
     - even trying to set an unexisting variable works `:GoDebugSet z 10`
     - a vim command to help see messages from vim-go
        - `:messages` `:messages clear`
- remote debugging
   - `go build -gcflags "all=-N -l" main.2023a02m21d.go`
   - ```sh
     dlv --listen=0.0.0.0:8181 --headless=true --api-version=2 --accept-multiclient exec ./main.2023a02m21d --log
     ```
   - `:GoDebugConnect 0.0.0.0:8181`
      - `dlv connect 0.0.0.0:8181`
- to download Golang Makefile
  - ```sh
    (GIST_ID=e3be354c2c3aa78a2f7956d046b0bf3a && \
     curl -L \
     'https://gist.githubusercontent.com/igorlima/'$GIST_ID'/raw/Makefile' \
     -o Makefile)
    ```

-----
<!-- vim-go - more -->
</details>


## further info

- [profiling]({% link docs/languages/golang/debug.profiling.md %})

----

