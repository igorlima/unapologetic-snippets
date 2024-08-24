---
layout: default
title: Shell - CLI fzf
nav_exclude: true
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/cli-fzf
---

# fzf _(Fuzzy Finder)_

__[back]({% link docs/languages/shell/cli.md %})__


## custom functions

fzf is a powerful fuzzy finder that can significantly enhance your command-line
workflow. Here are three custom functions that leverage fzf to streamline
common tasks. [^1]

```sh
# interactive process killing
function pkill() {
  ps aux | fzf --height 40% --layout=reverse --prompt="Select process to kill: " | awk '{print $2}' | xargs -r sudo kill
}
```

```sh
# git log navigation and browsing
function logg() {
  git lg | fzf --ansi --no-sort \
    --preview 'echo {} | grep -o "[a-f0-9]\{7\}" | head -1 | xargs -I % git show % --color=always' \
    --preview-window=right:50%:wrap --height 100% \
    --bind 'enter:execute(echo {} | grep -o "[a-f0-9]\{7\}" | head -1 | xargs -I % sh -c "git show % | nim -c \"setlocal buftype=nofile bufhidden=wipe noswapfile nowrap\" -c \"nnoremap <buffer> q :q!<CR>\" -")' \
    --bind 'ctrl-e:execute(echo {} | grep -o "[a-f0-9]\{7\}" | head -1 | xargs -I % sh -c "gh browse %")'
}
```


----

[^1]: [10 CLI Tools That Made the Biggest Impact](https://itnext.io/10-cli-tools-that-made-the-biggest-impact-f8a2f4168434)
