---
layout: default
title: Github CLI
nav_exclude: true
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/cli-gh
---

__[back]({% link docs/languages/shell/cli.md %})__


GitHub CLI brings the GitHub to the terminal[^1]. [GitHub CLI & hub](https://github.com/cli/cli/blob/586eb8b865ea539832716a3d8c5cb42b11ee784c/docs/gh-vs-hub.md).

- [https://cli.github.com/](https://cli.github.com/)

```bash
brew install gh
```

- `gh api <endpoint> [flags]`
  - [https://cli.github.com/manual/gh_api](https://cli.github.com/manual/gh_api)
  - `--hostname <string>`
    - _the GitHub hostname for the request (default "github.com")_
  - ```sh
    gh api graphql -F owner='{owner}' -F name='{repo}' -f query='
      query($name: String!, $owner: String!) {
        repository(owner: $owner, name: $name) {
          releases(last: 3) {
            nodes { tagName }
          }
        }
      }
    '
    ```
- `gh auth login`
- `gh repo --help`
- `gh issue create`
- `gh issue list`
- `gh label list [flags]`
  - `gh label list --sort name`
  - `gh label list --search bug`
- `gh pr`
  - [https://cli.github.com/manual/gh_pr](https://cli.github.com/manual/gh_pr)
    - `-R`, `--repo <[HOST/]OWNER/REPO>`
      - _select another repository using the [HOST/]OWNER/REPO format_
  - `gh pr list` `[flags]`
    - [https://cli.github.com/manual/gh_pr_list](https://cli.github.com/manual/gh_pr_list)
    - `-a`, `--assignee <string>`
    - `-A`, `--author <string>`
    - `-L`, `--limit <int>`
    - `-S`, `--search <query>`
    - `-s`, `--state <string>`
      - _filter by state: `{open|closed|merged|all}`_
  - `gh pr create`
  - `gh pr view [<number> | <url> | <branch>] [flags]`
    - `-c`, `--comments`
      - _view pull request comments_
  - `gh pr review` `[<number> | <url> | <branch>] [flags]`
    - `-a`, `--approve`
    - `-b`, `--body <string>`
      - _specify the body of a review_
    - `-c`, `--comment`
    - `-r`, `--request-changes`

----

[^1]: [GitHub CLI is Now Available: Here’s Why You Should Be Excited](https://blog.bitsrc.io/github-cli-is-now-available-heres-why-you-should-be-excited-91d8bdd81a51)
