---
layout: default
title: Bash - Sample
nav_exclude: true
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/bash-sample
---

# Bash Samples

__[back]({% link docs/languages/shell/bash.md %})__

<br/>
<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>


## Automatically Run a Command Every N Minutes

```bash
# Automatically Run a Command Every N Minutes
# How to Run or Repeat a Linux Command Every X Seconds
# Repeat a Unix command every x seconds
# https://askubuntu.com/questions/852070/automatically-run-a-command-every-5-minutes
# for (( i=0; i<=20; i++)); do echo "16733$i" ;sleep 1; done
for (( i=0; i<=20; i++)); do
  echo "$i" ;
  sleep 60;
done
```

----

[^1]: [...](https://www.google.com)

