---
layout: default
title: Writing Documentation for Golang
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/commands-get-private-repo
---


# How to fetch modules from private repository
__[back]({% link docs/languages/golang/commands.md %}#go-get)__
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

----

[^1]: [Today I Learned — Fix: go get private repository return error reading sum.golang.org/lookup … 410 gone](https://medium.com/mabar/today-i-learned-fix-go-get-private-repository-return-error-reading-sum-golang-org-lookup-93058a058dd8)
