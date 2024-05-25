---
layout: default
title: concurrency
nav_order: 2
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/concurrency
---

# Concurrency in Shell Scripts

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

## background execution

In shell scripting, commands run in the background by appending an ampersand (`&`) to the end of the command. [^1]

```sh
ls -la &
./some_sript.sh &
```

## using `await`

The `wait` command is used to pause scrit execution until all specified background processes have completed.

```sh
./task1.sh &
./task2.sh &
./task3.sh &

wait
echo "All tasks are completed."
```

To efficiently manage multiple processes, it's useful to store the process IDs (PIDs), then explicitly wait for specific processes to complete, which provides finer control over task synchronization.

```sh
process1(){
  sleep 10
}
process2(){
  sleep 20
}
process1 &
pid1=$!
process2 &
pid2=$!

# wait for each process to complete
wait $pid1
wait $pid2

echo "process 1 and process 2 are completed"
```

## parallel processing with `xargs`

```sh
{
  process_aux() {
    if [ -z "$1" ]; then
      # no argument
      echo "sleeping for 1sec"
      sleep 1
    else
      echo "sleeping for $1 sec"
      sleep $1
    fi
  }
  # note that `export -f` works only in Bash
  # so subprocess bash can see it
  export -f process_aux
  # run commands in parallel across 4 processes
  # if you do not specify the `-e` or the `-E` flags,
  #  underscore (`_`) is assumed for the logical EOF string.
  echo {1..10} | xargs -n1 -P4 bash -c 'process_aux "$@"' _
}
```

## comparison of commands

- `&` and `wait`: this is the simplest method to achieve concurrency but may require manual process management.
- `xargs`: allows simple parallel execution with minimal system load, suitable for tasks that are not interdependent.

----

[^1]: [Implementing Concurrency in Shell Scripts](https://dev.to/siddhantkcode/implementing-concurrency-in-shell-scripts-521o)
