---
layout: default
title: Makefile
nav_order: 3
parent: Shell Script
grand_parent: Programming Languages
permalink: /docs/languages/shell/makefile
---

# Makefile

A Makefile is a simple text file that contains instructions for building a software project. The Makefile specifies the project’s dependencies, build rules, and targets, allowing the ‘make’ utility to compile and link your source files into an executable program.

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

# Samples

- [Build Process with Makefiles in Golang]({% link docs/languages/golang/other.makefile.md %})
