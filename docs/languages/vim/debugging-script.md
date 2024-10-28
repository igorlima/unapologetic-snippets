---
layout: default
title: Script for Debugging
nav_order: 2
parent: Vim
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/vim/script-for-debugging
---

# Script for Debugging in Vim

- [go back]({% link docs/languages/vim/index.md %})

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

## Folding method

<details markdown="block">
  <summary>
    toggle between two folding methods
  </summary>

- filename: `vimrc`
```vim
" vim -N -u vimrc
" vim -N -u path/to/your/vimrc
"
" `:so vimrc`
" `:so %`
function! ToggleFoldMethod()
  " The 'foldmethod' option tells Vim what method to use when folding text.
  " https://vim.fandom.com/wiki/Folding
  " https://vimdoc.sourceforge.net/htmldoc/fold.html#zf
  "
  " There are two main methods of folding text: 'manual' and 'indent'.
  " • Manual
  "   - The 'manual' method is the simplest, and is the default.
  "   - It uses markers in the text to define the start and end of each fold.
  "     - `zf` to create
  "     - `zd` to delete
  "     - `zR` to open all folds
  "     - `zM` to close all folds
  "     - `zE` to delete all folds
  " • Indent
  "   - The 'indent' method is more sophisticated, and is often more convenient.
  "   - It uses the indentation of the text to define the start and end of each fold.
  let l:foldmethod = &foldmethod
  echom "fold method was " . l:foldmethod
  if l:foldmethod == 'manual'
    set foldmethod=indent
    echom "fold method is now indent"
  else
    set foldmethod=manual
    echom "fold method is now manual"
  endif

  " VIM FOLDING CHEATSHEET
  " `zj` moves the cursor to the next fold.
  " `zk` moves the cursor to the previous fold.
  " `zo` opens a fold at the cursor.
  " `zO` opens all folds at the cursor.
  " `zm` increases the foldlevel by one.
  " `zM` closes all open folds.
  " `zr` decreases the foldlevel by one.
  " `zR` decreases the foldlevel to zero -- all folds will be open.
  " `zd` deletes the fold at the cursor.
  " `zE` deletes all folds.
  " `[z` move to start of open fold.
  " `]z` move to end of open fold.
endfunc
" `:set foldmethod?`
```
</details>
