---
layout: default
title: Makefile
nav_order: 10
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/makefile
---

# Makefile

A Makefile is a text file that specifies the dependencies between files and the
commands needed to create or update those files.

A Makefile is a simple text file that contains instructions for building a
software project. The Makefile specifies the project’s dependencies, build
rules, and targets, allowing the ‘make’ utility to compile and link your source
files into an executable program.

Each rule specifies a target file, a list of dependency files, and a set of
commands to be executed to create or update the target file. [^1]

If the module is missing, or if it is older than the dependency file, make
considers it to be out of date, and issues the commands necessary to rebuild
it. A module can be treated as out of date if the commands used to build it
have changed.

## Structure

A makefile consists of __three sections__: _target, dependencies, and rules_. The __target__ is normally either an executable or object file name. The __dependencies__ are source code or other things needed to make the target. The __rules__ are the commands needed to make the target.

A simple makefile might be structured like this:

```Makefile
# comments are preceded by the hash symbol
target: dependencies
        command 1
        command 2
        ...
        command n
```

## Conditional Directives

The conditional directives are:

- __ifeq__ - the `ifeq` directive begins the conditional, and specifies the condition. It contains two arguments, separated by a comma and surrounded by parentheses. Variable substitution is performed on both arguments and then they are compared. The lines of the makefile following the ifeq are obeyed if the two arguments match; otherwise they are ignored.
- __ifneq__ - the `ifneq` directive begins the conditional, and specifies the condition. It contains two arguments, separated by a comma and surrounded by parentheses. Variable substitution is performed on both arguments and then they are compared. The lines of the makefile following the ifneq are obeyed if the two arguments do not match; otherwise they are ignored.
- __ifdef__ - the `ifdef` directive begins the conditional, and specifies the condition. It contains single argument. If the given argument is true then condition becomes true.
- __ifndef__ - the `ifndef` directive begins the conditional, and specifies the condition. It contains single argument. If the given argument is false then condition becomes true.
- __else__ - the `else` directive causes the following lines to be obeyed if the previous conditional failed. In the example above this means the second alternative linking command is used whenever the first alternative is not used. It is optional to have an else in a conditional.
- __endif__ - the `endif` directive ends the conditional. Every conditional must end with an endif.

## Syntax of Conditionals Directives

The syntax of a simple conditional with no else is as follows.

```
conditional-directive
    text-if-true
endif
```

The syntax of a complex conditional is as follows.

```
conditional-directive
    text-if-true
else
    text-if-false
endif
```

## Passing additional variables from command line to make

The command args overwrite environment variable.

Makefile:
```
send:
    echo $(MESSAGE1) $(MESSAGE2)
```

Example run:

```
$ MESSAGE1=YES MESSAGE2=NG  make send MESSAGE2=OK
echo YES OK
YES OK
```

## Function

Make has a decent amount of [builtin functions](https://www.gnu.org/software/make/manual/html_node/Functions.html).

### The shell function

shell - This calls the shell, but it replaces newlines with spaces!

```
all:
	@echo $(shell ls -la) # Very ugly because the newlines are gone!
```

## Other

### PHONY

- `.PHONY: all build test`: Specifies that `all`, `build`, and `test` are phony
  targets, **meaning that they don't correspond to actual files.** [^1]

In a Makefile, `.PHONY` is a special target that is used to declare phony
targets. A phony target is a target that does not represent a physical file but
instead represents a command that needs to be executed.  When a phony target is
declared using `.PHONY`, it tells Make that the target should always be
considered out-of-date and should be rebuilt every time it is requested,
regardless of whether a file with the same name as the target exists or not.

# Samples

- [Build Process with Makefiles in Golang]({% link docs/languages/golang/other.makefile.md %})


[^1]: [A Comprehensive Guide for Using Makefile In Golang Projects](https://levelup.gitconnected.com/a-comprehensive-guide-for-using-makefile-in-golang-projects-c89edebcbe6e)
