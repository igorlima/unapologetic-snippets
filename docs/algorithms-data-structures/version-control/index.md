---
layout: default
title: Version Control
nav_order: 2
parent: Algorithms and Data Structures
has_children: false
has_toc: true
permalink: /docs/algorithms-and-data-structures/version-control
---

# Version Control

- [go back]({% link docs/algorithms-data-structures/index.md %})

## Git


- <details markdown="block"><summary>shorten git history</summary>
   
  ```sh
  # BURN git to the ground, keep only latest commit, shrink size:
  # https://gist.github.com/stucka/2973d0702f05003c5061e57b3df29f70
  # adapted from https://stackoverflow.com/questions/9683279/make-the-current-commit-the-only-initial-commit-in-a-git-repository
  git checkout --orphan newBranch
  git add -A  # Add all files and commit them
  git commit
  git branch -D master  # Deletes the master branch
  git branch -m master  # Rename the current branch to master
  git gc --aggressive --prune=all     # remove the old files
  git push -u -f origin master  # Force push master branch to github
  git gc --aggressive --prune=all     # remove the old files
  ```
  ```sh
  # REMOVE all but LAST N COMMITS from git history to save space
  # https://stackoverflow.com/questions/72995274/how-to-remove-all-but-last-n-commits-from-git-history-to-save-space
  git switch --detach <new-repo-root-commit-id>
  git reset --soft <old-repo-root-commit-id>
  git commit --amend --reuse-message=<new-repo-root-commit-id>
  git rebase <new-repo-root-commit-id> main --onto @
  ```
  - the process is:
    1. the branch name you wish to rewrite. _(this answer will assume it is called `main`.)_
    2. the root commit of your branch. Get it with: `git log main --reverse` _(this answer will assume it is called `<old-repo-root-commit-id>`.)_
    3. the oldest commit ID you want to keep, e.g. your forth commit from the top, which will become your new repo root commit. _(this answer will assume it is called `<new-repo-root-commit-id>`.)_
    4. your `git status` should be clean before you start. If it isn't, consider committing (or undoing) your most recent changes.

  ```sh
  # TRUNCATING git history
  # https://passingcuriosity.com/2017/truncating-git-history/
  git checkout --orphan temp e41d7f633c45c46bd42e97cecf93204191d9e4c9
  git commit -m "Truncate history"
  git rebase --onto temp e41d7f633c45c46bd42e97cecf93204191d9e4c9 master
  ```
  - the process is:
    1. make a new branch (I’ll call it `temp` below); and
    2. use `rebase` to replay every subsequent commit on top of the new branch.
  </details>
- <details markdown="block"><summary>find size of git repository</summary>
   
  ```sh
  # https://stackoverflow.com/questions/8185276/find-size-of-git-repository -
  git count-objects -vH
  ```
  </details>


------ ------

[^1]: [...)](...)
