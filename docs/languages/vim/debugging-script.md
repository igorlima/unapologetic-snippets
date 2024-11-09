---
layout: default
title: Script for Debugging
nav_order: 9
parent: Vim
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/vim/script-for-debugging
---

# Script for Debugging in Vim

- [go back]({% link docs/languages/vim/index.md %})
- [vim snippets]({% link docs/algorithms-data-structures/snippets/index.md %})
  <sup>[+](https://github.com/igorlima/unapologetic-snippets/tree/main/docs/algorithms-data-structures/snippets/vim)</sup>

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

## Playground

<details markdown="block">
  <summary>
    date format / register / url regex / if condition
  </summary>

- filename: `playground.vimrc`

```vim
" vim -e playground.vimrc
" vim -N -u playground.vimrc
" vim -N -u path/to/your/vimrc
"

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
</details>

<details markdown="block">
  <summary>
    counting the number of occurrences of each word in the selection
  </summary>

```vim
" This function will count the number of occurrences of each word in the selection
" vim -N -u path/to/your/vimrc
" vim -N -u test.vimrc
function SGrep0()
  " if the execute comes from visual mode, it will yank the selection in register t
  normal! gv"ty
  " let l:selection = @* " read it from clipboard
  " let l:selection = @@ " read it from yank register
  let l:selection = @t " read it from the register 't'
  let l:splitted = split(l:selection, " ")
  " let l:splitted = map(l:splitted[:], 'substitute(v:val, "\s", "", "g")')
  let l:splitted = map(l:splitted[:], 'trim(v:val)')
  let l:splitted = filter(l:splitted, 'v:val != ""')
  let l:counter = {}
  for l:tag in l:splitted
    let l:pos = search(l:tag, 'W')
    while l:pos != 0
      let l:counter[l:tag] = get(l:counter, l:tag, 0) + 1
      let l:pos = search(l:tag, 'W')
    endwhile
  endfor
  redir @z
  for [key, value] in items(l:counter)
    echom key . " " . value
  endfor
  redir END
  echom "coped it into `\"z`, paste it by pressing `\"zp` in normal mode"
endfunction
nnoremap <c-k>s0 :call SGrep0()<CR>
vnoremap <c-k>s0 :call SGrep0()<CR>
```

</details>

<details markdown="block">
  <summary>
    setting markers using search via vim script
  </summary>

- filename: `playground.vimrc` <a id="playground-setting-marker"></a>

```vim
" This function creates a marker for each pattern in the list
" `vim -N -u path/to/your/vimrc`
" `vim -N -u playground.vimrc`
"
function! GenerateVimwikiMarkerTemplateA()
  let l:markers = [
    \ ['a', '^\#\# Meetings$'],
    \ ['b', '^\#\# Messages$'],
    \ ['c', '^\#\# Emails$'],
    \ ['d', '^\#\# Notes$'],
    \ ['e', '^\#\# Energy Level and Mood Tracker$'],
  \]
  for l:marker in l:markers
    let l:letter = l:marker[0]
    let l:pattern = l:marker[1]
    let l:pos = search(l:pattern, 'W')
    " https://vimdoc.sourceforge.net/htmldoc/motion.html#%3amark
    execute ':' . l:pos . 'mark ' . l:letter
  endfor
  echom "markers created: 'a' for Meetings, 'b' for Messages, 'c' for Emails, 'd' for Notes, 'e' for Mood Tracker"
endfunction

nnoremap <c-k><c-m> :call GenerateVimwikiMarkerTemplateA()<CR>
inoremap <c-k><c-m> <esc>:call GenerateVimwikiMarkerTemplateA()<CR>a

" REFERENCES:
" • Looping through Vimscript: iterate over an array or dictionary
"  - https://learnvimscriptthehardway.stevelosh.com/chapters/36.html
"  - https://learnvimscriptthehardway.stevelosh.com/chapters/37.html
"  - https://renenyffenegger.ch/notes/development/vim/script/vimscript/arrays-and-hashes/iterate_over_array
" • search() function
"  - How can I tell if a search was successful in vimscript?
"   - https://stackoverflow.com/questions/27660221/how-can-i-tell-if-a-search-was-successful-in-vimscript
"  - How can I make a search in vimscript let n and N look for more?
"   - https://vi.stackexchange.com/questions/3655/how-can-i-make-a-search-in-vimscript-let-n-and-n-look-for-more
"  - VIM check for search pattern match in vim script
"   - https://stackoverflow.com/questions/39930288/vim-check-for-search-pattern-match-in-vim-script
" • set mark
"  - Vim documentation: motion
"   - https://vimdoc.sourceforge.net/htmldoc/motion.html#%3amark
"  - set marker (tag) in vim command
"   - https://stackoverflow.com/questions/10034645/set-marker-tag-in-vim-command
" • Using markers for folding
"  - Writing a custom fold expression
"   - http://vimcasts.org/transcripts/38/en/
"  - How can I automatically fold a file when start and end markers are the same?
"   - https://superuser.com/questions/557502/how-can-i-automatically-fold-a-file-when-start-and-end-markers-are-the-same
"  - vim: set fix jump marker in a source comment of the current file
"   - https://stackoverflow.com/questions/50985045/vim-set-fix-jump-marker-in-a-source-comment-of-the-current-file
"
" ALTERNATIVE SCRIPT:
" GO TO:
" - either `/ctrl-r 0` or `:execute '/' . @0`
"   - `/ctrl-r"` or `:execute '/' . @"`
"   - `:@"` or `:@0`
"     - `:execute '/.\+Mettings$'`
"     - `:execute '/.\+Messages$'`
"     - `:execute '/.\+Emails$'`
"     - `:execute '/.\+Notes$'`
"   - `:[range]ma[rk] {a-zA-Z'}`
"     - `https://vimdoc.sourceforge.net/htmldoc/motion.html#%3amark`
"     - `'a`: Mettings
"     - `'b`: Messages
"     - `'c`: Emails
"     - `'d`: Notes
"     - `let @a = search('^\#\# Meetings$')`
```
</details>
