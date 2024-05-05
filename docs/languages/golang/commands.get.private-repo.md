---
layout: default
title: Writing Documentation for Golang
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/commands-get-private-repo
---


# How to fetch modules from private repository
__[back to `go get`]({% link docs/languages/golang/commands.md %}#go-get)__
<br/>

Fix: go get private repository return error reading sum.golang.org/lookup … 410 gone [^1]

Fixing error: go get gomodule on private module with error message pattern verifying git-host… reading sum.golang.org/lookup … 410 gone [^1]

## Using `GOPRIVATE`

As stated in the release doc of Go 1.13,

> The new GOPRIVATE environment variable indicates module paths that are not publicly available. It serves as the default value for the lower-level GONOPROXY and GONOSUMDB variables, which provide finer-grained control over which modules are fetched via proxy and verified using the checksum database.

```sh
export GOPRIVATE="gitlab.com/idmabar,bitbucket.org/idmabar,github.com/idmabar"
```

And to verify if this worked, you can do the `go env` command.

```sh
(GOPRIVATE="gitlab.com/idmabar,bitbucket.org/idmabar,github.com/idmabar" go env) | grep GOPRIVATE
```

So this `env-variable` will tell the `go get` command to use the private host proxy to retrieve the package.


## Using `GONOSUMDB`

Another solution, maybe using the GONOSUMDB variable.

So try to set this in as an environment variable:

```sh
export GONOSUMDB="gitlab.com/idmabar,bitbucket.org/idmabar,github.com/idmabar"
```

Actually this issue only happens in new Golang version 1.13 and afterward.


## Miscellaneous

- clone repo and run `GOINSECURE=true go get`
  - troubleshooting
    - `GOPRIVATE=gitlab.xyz.tech go get all`
      - [GOPRIVATE environment variable](https://stackoverflow.com/questions/58305567/how-to-set-goprivate-environment-variable)
        - [`GOPRIVATE=*.mydomain.com`](https://stackoverflow.com/questions/65755940/go-get-fails-to-download-a-go-package-git-repository-hosted-on-a-github-entre)
        - [`GOPRIVATE=*.gitlab.xyz.tech`](https://github.com/goproxyio/goproxy.io/blob/4de6666201d3b15e26ed53a8d03aa42ef31c5f34/content/docs/GOPRIVATE-env.md)
    - `go clean --modcache`
      - `GOPRIVATE=gitlab.xyz.tech,xyz-github.corp go get all`
      - _if any of those below..._
        - `verifying module - 404 Not Found`
        - `server response: not found`
        - `dial tcp: lookup - no such host`
    - `go get -insecure`
      - _`-insecure` flag is no longer supported; use `GOINSECURE` instead_

----

[^1]: [Today I Learned — Fix: go get private repository return error reading sum.golang.org/lookup … 410 gone](https://medium.com/mabar/today-i-learned-fix-go-get-private-repository-return-error-reading-sum-golang-org-lookup-93058a058dd8)
