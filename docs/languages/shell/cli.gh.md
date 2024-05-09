---
layout: default
title: Shell - CLI Github
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

- ```sh
  # list of prs involving me
  gh search prs --state=open --involves=@me
  # search by PR title
  gh pr list -S "state:open type:pr in:name IPM"
  # looking for things related to me
  gh pr list -S "is:open type:pr involves:@me"
  # view PR using a pager
  PAGER=less gh pr view 13992
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
- ```sh
  # ALIAS
  gh alias set all-with-me 'search prs --state=open --involves=@me'
  gh alias set with-me 'pr list -S "is:open type:pr involves:@me"'
  ```

## GitHub Search Queries

<br>
*__GitHub Search Tips__*

to exclude something prepend a hyphen (`-`):
```
foo_library -repo:owner1/repoX -repo:owner2/repo
```

- basic
  - `is:issue`
  - `is:open`
  - `is:public`
  - `is:private`
  - `label:beginner `
  - `type:issue`
  - `type:pr`
  - `no:project`
  - `no:milestone`
  - `no:label`
  - `no:assignee`
- <details markdown="block"><summary><i>further more...</i></summary>

  - find by people
    - `involves:octocat`
    - `author:octocat`
    - `user:octocat`
    - `commenter:octocat`
    - `answered-by:octocat`
    - `assignee:octocat`
  - find by content through name, description/readme, and so on
    - `in:name`
    - `in:title`
    - `in:description`
    - `in:readme`
    - `in:topic`
    - `in:body`
    - `in:comments`
  - find by date: `YYYY-MM-DD`
    - use `<`, `>`, `>=`, `<=` and `..`
    - `created:<2022-10-01`
    - `updated:>2022-10-01`
    - `merged:>=2022-10-01`
    - `closed:<=2022-10-01`
    - `closed:2024-05-01..2024-05-09`
  - find numbers through stars, forks, and so on
    - to further narrow your search use: `<`, `>`, `<=`, `>=` & `..`
    - `stars:n`
    - `stars:1000`
    - `forks:n`
    - `forks:<100`
    - `comments:n`
  - miscellaneous
    - `org:freecodecamp`
    - `language:LANGUAGE`
    - `language:PHP`
    - `license:LICENSE_KEYWORD`
    - `license:MIT`
    - `in:path user.rb` or `filename:user.rb`
    - `sort:reactions-+1-desc`
    - `state:open`
    - `draft:false`
  - additional resources:
    - [github search documentation](https://docs.github.com/en/search-github)
      - [understanding code search syntax](https://docs.github.com/en/search-github/github-code-search/understanding-github-code-search-syntax)
        - [using boolean operations](https://docs.github.com/en/search-github/github-code-search/understanding-github-code-search-syntax#using-boolean-operations)
        - [using qualifiers](https://docs.github.com/en/search-github/github-code-search/understanding-github-code-search-syntax#using-qualifiers)
        - [using regular expressions](https://docs.github.com/en/search-github/github-code-search/understanding-github-code-search-syntax#using-regular-expressions)
    - searching on GitHub
      - [searching discussions](https://docs.github.com/en/search-github/searching-on-github/searching-discussions)
      - [searching code](https://docs.github.com/en/search-github/searching-on-github/searching-code)

  <br/>
  </details>

----

[^1]: [GitHub CLI is Now Available: Here’s Why You Should Be Excited](https://blog.bitsrc.io/github-cli-is-now-available-heres-why-you-should-be-excited-91d8bdd81a51)
