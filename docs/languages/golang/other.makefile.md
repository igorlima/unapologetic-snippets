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

<br/>

```
$(info ************ Build Process with Makefiles in Golang ************)

# Define Go command and flags
GO = go
DELVE = dlv
DEBUGPORT = 8181
GODEBUGFLAGS = -gcflags "all=-N -l"
GOFLAGS = -ldflags="-s -w"

# Makefiles MUST be indented using TABs and not spaces or `make` will fail.

# Why does make think the target is up to date?
# It happens when you have a file with the same name as Makefile target name in the directory where the Makefile is present.

# Makefile Tutorial
# https://makefiletutorial.com/
# https://ftp.gnu.org/old-gnu/Manuals/make-3.79.1/html_chapter/make_6.html
# https://ftp.gnu.org/old-gnu/Manuals/make-3.79.1/html_chapter/make_7.html

# Define the target executable
TARGET = my_app

# Change the main file if needed via the variable MY_GO_FILE/file
# $> export MY_GO_FILE=main.go
# $> MY_GO_FILE=main.go make
# $> make file=main.go
ifneq ($(origin MY_GO_FILE), undefined)
	MAIN_FILE := $(MY_GO_FILE)
endif
ifneq ($(origin file), undefined)
	MAIN_FILE := $(file)
endif
ifeq ($(origin MAIN_FILE), undefined)
	MAIN_FILE := main.go
endif
$(info MAIN_FILE is $(MAIN_FILE))

# Default target: build the executable
all:
	@echo "\nprerequisites needed\nFILES:\n• $(TARGET)\n• $(MAIN_FILE) \n\nENV VAR:\n• MY_GO_FILE: $(MY_GO_FILE)"

# Rule to build the target executable
$(TARGET): $(MAIN_FILE)
	$(GO) build $(GOFLAGS) -o $(TARGET) $(MAIN_FILE)
# Rule to build the target executable
build: $(MAIN_FILE)
	$(GO) build $(GOFLAGS) -o $(TARGET) $(MAIN_FILE)

# Run target: build and run the target executable
run: $(TARGET) $(MAIN_FILE)
	./$(TARGET)

# Clean target: remove the target executable
clean: $(TARGET)
	rm -f $(TARGET)

# Test target: run Go tests for the project
test:
	$(GO) test ./...

debug-build:
	$(GO) build $(GODEBUGFLAGS) -o $(TARGET) $(MAIN_FILE)

debug-run: 
	$(DELVE) --listen=0.0.0.0:$(DEBUGPORT) --headless=true --api-version=2 --accept-multiclient exec ./$(TARGET) --log

debug-conn:
	$(DELVE) connect 0.0.0.0:$(DEBUGPORT)
```
