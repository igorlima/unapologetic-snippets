---
layout: default
title: Debugging
nav_order: 8
parent: Vim
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/vim/debugging
---

# Debugging Vim

- [go back]({% link docs/languages/vim/index.md %})

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

## Vim Modes

- Normal
- Insert
- Visual
- Ex mode
  - `:help Ex-mode`
  - to enter Ex mode type either `Q` or `gQ`
  - when entering Ex mode. Type `visual` to go to Normal mode.
  - browse your command-line history
    - `q:`, or `:<C-f>`, instead provides a way to browse command-line history and edit it like a normal buffer. This makes it easy to find a previous command you ran, edit it with normal Vim commands, and then run the modified command. The `q/` and `q?` commands exist to provide the same functionality for the search history.

## Debugging Vimscript

```vim
" debug on a line
debug echo "here"

" debug a command
debug CommandName

" debug a fucntion
debug call FunctionName(arg)
```

Interact with the debugger:
- `cont`: continue execution until the next breakpoint (if one exists)
- `quit`: stop current execution, but still stops at the next breakpoint
- `step`: execute the current command and come back to debug mode when it is finished
- `next`: like `step` except it also steps over function calls and sourced files
- `interrupt`: like quit, but returns to debug mode for the next command
- `finish`: finishes the current script or function and returns to debug mode for the next command

## Vimscript's Debugger

You can start up Vim in its debug mode with the `-D` argument.

```sh
# opens vim in debug mode
vim -D somefile.txt
```

This is useful for seeing what is run when Vim initializes and might help give you a better idea of everything that happens when you open Vim. You can use this to walk through the Vimscript in your `vimrc` that runs when an instance of Vim initializes. I will explain how to use the debugger momentarily.

If you are interested in profiling your Vim's startup, you may be interested in starting Vim with the `startuptime` option, which will write to a file you specify with timestamps and information about what Vim is loading. You can use this to find bottlenecks or problematic plugins that might be causing Vim to be slow.

```sh
# this will write startup info to a file: vim.log
vim --startuptime vim.log
```

## **How to Use the Debugger**

Before we jump into how to use the debugger, you should also know that you can do more than just start up Vim in debug mode as that would not be terribly useful. You can also call the `debug` command from within Vim to debug Vimscript functions or commands that have already been sourced.

```vim
" debug a command
debug CommandName

" debug a fucntion
debug call FunctionName(arg)
```

Once you are in debug mode, you will see a prompt (`>`) in your status bar. This is where you can control and interact with the debugger.

- `cont`: continue execution until the next breakpoint (if one exists)
- `quit`: stop current execution, but still stops at the next breakpoint
- `step`: execute the current command and come back to debug mode when it is finished
- `next`: like `step` except it also steps over function calls and sourced files
- `interrupt`: like quit, but returns to debug mode for the next command
- `finish`: finishes the current script or function and returns to debug mode for the next command

You should note that each of these debug commands has the shorthand of the first letter of the word `[c, q, s, n, i, f]` and hitting enter (`<CR>`) will execute the previous command again.

The last thing you need to know is how to manage breakpoints to your Vimscript. You can add a breakpoint with the `breakadd` command, delete it with the `breakdel` command, and see all breakpoints with `breaklist`.

_**Adding Breakpoints**_

```vim
" bracketed [] arguments are optional

" add breakpoint to function
breakadd func [lineNumber] functionName
" add breakpoint to file
breakadd file [lineNumber] fileName
" add breakpoint to current line of current file
breakadd here
```

_**Deleting Breakpoints**_

```vim
" bracketed [] arguments are optional

" delete breakpoint number from breaklist output
breakdel number
" delete all breakpoints
breakdel *
" delete breakpoint on function
breakdel func [lineNumber] functionName
" delete breakpoint on file
breakdel file [lineNumber] fileName
" delete breakpoint at current line of current file
breakdel here
```

Read the documentation for `debug`:
- `:help debug`
- http://vimdoc.sourceforge.net/htmldoc/repeat.html#:debug

## General Debugging with Verbose

Other than the `debug` command, Vim also has a `verbose` command that provides logging when actions take place in Vim, which can also be useful for debugging, especially things like autocommands. I recently used verbose to determine what autocommands I could leverage to hook into netrw ([see post](http://inlehmansterms.net/2014/09/04/sane-vim-working-directories/)). Using verbose is as simple as:

```vim
" set verbose command
set verbose=9

" set verbose on startup
vim -V9 file.text
```

You can set verbose to any value from 0-15, where 0 is no extra output, and 15 gives you output for everything Vim is doing. Each number progressively adds more output and insight into what Vim is doing:

```
" >= 1  When the viminfo file is read or written.
" >= 2  When a file is ":source"'ed.
" >= 5  Every searched tags file and include file.
" >= 8  Files for which a group of autocommands is executed.
" >= 9  Every executed autocommand.
" >= 12 Every executed function.
" >= 13 When an exception is thrown, caught, finished, or discarded.
" >= 14 Anything pending in a ":finally" clause.
" >= 15 Every executed Ex command (truncated at 200 characters).
```

It will not take much playing around with verbose to become overwhelmed by the amount of output showing up in the status bar. Each message interrupts the flow of Vim, and may make it difficult to do anything. To get around this you can send verbose output to a file:

```vim
set verbosefile=filename.txt
```

Lastly, you can run verbose for just a single command using `[count]verbose`, which is useful if you only want to see what Vim is doing behind the scenes for one small action.

```
:12verbose call MyFunction()
```

Read the documentation for verbose:
- `:help verbose`
- http://vimdoc.sourceforge.net/htmldoc/options.html#'verbose'

Now you should have a good idea of where to begin when debugging and logging output in Vim. Though Vimscript may not be the most enjoyable language to work with, it comes with tools like `verbose` and `debug` to help you fight the bane of each Vim developer's existence, Vimscript.


## Play with Vimscript

<details markdown="block">
  <summary>
    a file to play with
  </summary>

```sh
vim -e playground.vimrc
```

```vim
" we can enter Ex mode from Normal mode by
" pressing `Shift-Q`. This switches to Ex mode
" and displays a prompt at the bottom of the
" screen where we can enter commands:
"
" alternatively, we can start Vim in Ex mode
" by using the -e or -E options when
" launching Vim from the command line:
" $> vim -e

" a function to format the date in the same way I use to do on Workflowy
" https://vim.fandom.com/wiki/Insert_current_date_or_time
" https://stackoverflow.com/questions/56052/best-way-to-insert-timestamp-in-vim
function! MyDate()
  let l:date = strftime('%Ya%mm%dd %Hh%Mm%Ss')
  return l:date
endfunction
echo MyDate()

" let l:selection = @* " read it from clipboard
" let l:selection = @@ " read it from yank register
" let @y = 'hello'
function! GetRegValue()
  " let l:selection = @y " read it from the register 'y'
  let l:selection = getreg('y') " read it from the register 'y'
  return l:selection
endfunction
echom GetRegValue()

" REGEX
" For more information, see:
" :h matchstr()
" :h /\c
" :h /\zs
" :h /\{-
let test = 'a href="http://www.google.com">www.google.com</a;'
let url = matchstr(test, '\ca href=\([''"]\)\zs.\{-}\ze\1')
echo url
if empty(url) 
  throw "no url recognized into ``".test."''"
endif
" IF condition
echo '12345678-1234-1234-1234-123456789012' =~ '^[0-9a-z]\{8}-[0-9a-z]\{4}-[0-9a-z]\{4}-[0-9a-z]\{4}-[0-9a-z]\{12}'
echo '2024a' =~ '^[0-9a-z]\{8}-[0-9a-z]\{4}-[0-9a-z]\{4}-[0-9a-z]\{4}-[0-9a-z]\{12}'
```

------

</details>
