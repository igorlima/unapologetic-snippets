---
layout: default
title: Vim
nav_order: 2
parent: Programming Languages
has_children: true
permalink: /docs/languages/vim
---

## Search

- search through the folder of the current file path by using the absolute folder path.
  - `:lgrep -E "gz" -0 --color -n %:p:h/*`
- <details markdown="block"><summary>more...</summary>

  - `:vimgrep /regex/ %:p:h/**/*` -- search recursively by using regex
    - `:vimgrep "word" %:p:h/**/*` -- search recursively by using a single word
    - __location list__ `:lvimgrep "word" %:p:h/**/*` -- search recursively by using a single word
  - vim can interface with "grep" and grep-like programs
    - `:grep` `:lgr` `:lgrep`
      - `grep -E "test" -0 -nR src --exclude="*node_modules*"`
      - `lgrep -E "test" -0 -nR src --exclude="*node_modules*"`
    - the location list behaves just like the quickfix list except that it is local to the current window instead of being global to the Vim session. So if you have five open windows, you can have up to five location lists, but only one quickfix list.
      - __quickfix list__
        - `grep -E "test" -0 -nR src --exclude="*node_modules*"`
          - navigation: `[[q`, `]]q`, `[q`, `[Q`, `]q`, `]Q`
          - `<leader>qa`, `<leader>qX`, `<leader>qx`, `<leader>qs`
      - __location list__
        - `lgrep -E "test" -0 -nR src --exclude="*node_modules*"`
          - navigation: `[[l`, `]]l`, `[l`, `[L`, `]l`, `]L`
          - `<leader>la`, `<leader>lX`, `<leader>lx`, `<leader>ls`
  </details>

## How to...

__How to use VI to search for lines with a certain length?__

- [find any lines exceeding a certain length]({% link docs/languages/shell/text-processing.md %})

how to use VI to find lines longer than a certain number of characters:
```vim
/\%>80v.\+
```

to find lines shorter than 25 characters:
```vim
/^.\{,25}$/
```
