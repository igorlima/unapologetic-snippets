---
layout: default
title: Other Makefile
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/other-makefile
---

__[back]({% link docs/languages/golang/other.md %})__

- further info:
  - [Makefile]({% link docs/languages/shell/makefile.md %})


## Makefile sample

Any **update** in this Makefile will be [in this gist](https://gist.github.com/igorlima/e3be354c2c3aa78a2f7956d046b0bf3a).

<br/>

```
$(info ************ Build Process with Makefiles in Golang ************)

# Makefiles MUST be indented using TABs and not spaces or `make` will fail.

# Makefile Tutorial
# https://makefiletutorial.com/
# https://ftp.gnu.org/old-gnu/Manuals/make-3.79.1/html_chapter/make_6.html
# https://ftp.gnu.org/old-gnu/Manuals/make-3.79.1/html_chapter/make_7.html

# VARIABLES
GO_MAKEFILE = makefile.go.json
GIST_ID = e3be354c2c3aa78a2f7956d046b0bf3a

# default target
all:
	@echo "\nrun 'make update-makefile'\nto have the latest change from:\n• https://gist.github.com/igorlima/$(GIST_ID)"

update-makefile:
	curl https://gist.githubusercontent.com/igorlima/$(GIST_ID)/raw/Makefile -o Makefile
ifeq (, $(shell which jq))
	@echo "no jq installed, install it to make possible to have the Makefile Go JSON"
else
	$(eval COMMITS = $(shell curl -L \
		-H "Accept: application/vnd.github+json" \
		-H "X-GitHub-Api-Version: 2022-11-28" \
		https://api.github.com/gists/$(GIST_ID)/commits?per_page=10))
	$(shell echo '$(COMMITS)' | jq -r '.[0] | {version, committed_at}' > $(GO_MAKEFILE))
endif

# https://www.gnu.org/software/make/manual/html_node/Multi_002dLine.html
# https://docs.github.com/en/rest/gists/gists?apiVersion=2022-11-28#list-gist-commits
```

Shorten links for Makefile GoLang:
- http://makefile-go-download.ilima.xyz
- http://makefile-go-gist.ilima.xyz
