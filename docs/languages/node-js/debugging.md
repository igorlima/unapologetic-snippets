---
layout: default
title: Debugging
nav_order: 5
parent: Node JS
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/node-js/debugging
---

There are three main pillars of being a good software engineer: writing code, debugging code and reviewing code.

> “The art of debugging is figuring out what you really told your program to do rather than what you thought you told it to do.” _Andrew Singer_

> "Debugging is a separate skill from writing code. Writing code is creative, its like art. Debugging is like science, you need to formulate a small hypothesis and test it, then keep doing that till you've fixed the problem".

other resources:
- not yet available

<br/>

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

----

## debugging on the terminal

```sh
node -h
-e, --eval script          evaluate script
-i, --interactive          always enter the REPL even if stdin
-r, --require=...          module to preload (option can be repeated)
```

```sh
# RUNning REPL with preloading a module
node -i -r "./build/main/index.js"
```

```sh
# RUNning REPL by injecting a script in advance
node -i -e '
function decode(text = ``) {
  const bufferedText = Buffer.from(text, `base64`)
  try {
    return JSON.parse(bufferedText.toString(`utf-8`))
  } catch (error) {
    console.error(`Error parsing decoded text: ${error}`)
    return {}
  }
}

function encode(text = {}) {
  const bufferedText = JSON.stringify(text)
  const encodedString = Buffer.from(bufferedText, `utf-8`).toString(`base64`)
  return encodedString
}
'
```

```sh
# DEBUGging a script in REPL mode
# use either `inspect` or `--inspect`
node inspect -i -e '
function decode(text = ``) {
  const bufferedText = Buffer.from(text, `base64`)
  debugger
  try {
    return JSON.parse(bufferedText.toString(`utf-8`))
  } catch (error) {
    console.error(`Error parsing decoded text: ${error}`)
    return {}
  }
}

function encode(text = {}) {
  debugger
  const bufferedText = JSON.stringify(text)
  const encodedString = Buffer.from(bufferedText, `utf-8`).toString(`base64`)
  return encodedString
}
'
```

## how to monitor the memory usage of Node.js?

```js
// How to monitor the memory usage of Node.js?
// https://stackoverflow.com/questions/20018588/how-to-monitor-the-memory-usage-of-node-js
const memoryUsage = () => {
  const formatMemoryUsage = (data) => `${Math.round(data / 1024 / 1024 * 100) / 100} MB`;
  const memoryData = process.memoryUsage();
  const memoryUsage = {
    rss: `${formatMemoryUsage(memoryData.rss)} -> Resident Set Size - total memory allocated for the process execution`,
    heapTotal: `${formatMemoryUsage(memoryData.heapTotal)} -> total size of the allocated heap`,
    heapUsed: `${formatMemoryUsage(memoryData.heapUsed)} -> actual memory used during the execution`,
    external: `${formatMemoryUsage(memoryData.external)} -> V8 external memory`,
  };
  return memoryUsage;
};
console.log(memoryUsage());
```

## how to take a heapdump of a Node.js process?

```js
const v8 = require("v8");
process.on('SIGUSR2', () => {
  const fileName = v8.writeHeapSnapshot();
  console.log(`Created heapdump file: ${fileName}`);
});
```
```sh
lsof -i :3000
kill -SIGUSR2 <the_process_id>
```
- **reference:**
  - [Preventing and Debugging Memory Leaks in Node.js](https://betterstack.com/community/guides/scaling-nodejs/high-performance-nodejs/nodejs-memory-leaks/) [<sup>+</sup>](https://github.com/igorlima/unapologetic-thoughts/blob/48130e8cb5d76a920cca105ae5d18be5037b6410/snippets/nodejs/heapdump.js)

## debugging performance on the terminal

```sh
# quote the label to prevent the backticks from being evaluated
cat << "EOF" > node-testing.js
return new Promise(resolve => {
  console.log(`doing some work...`)
  setTimeout(function() {
    callback()
    resolve(true)
  }, 1000)
}).then((data) => {
  console.log(data)
  debugger
})
EOF

# start a debug listener
node --inspect-brk=0.0.0.0:9235 node-testing.js

# connect to it
node inspect 0.0.0.0:9235
```

<details markdown="block">
  <summary>
    Performance Timing - Node.js v14.21.3
  </summary>

```js
// Node.js v14.21.3
// node --inspect-brk=0.0.0.0:9235 sample-measuring.js
const { performance, PerformanceObserver } = require('perf_hooks')
// https://nodejs.org/docs/latest-v14.x/api/perf_hooks.html

performance.clearMarks()
const observer = new PerformanceObserver((list, obs) => {
  // https://developer.mozilla.org/en-US/docs/Web/API/PerformanceObserverEntryList/getEntriesByName
  // https://nodejs.org/docs/latest-v14.x/api/perf_hooks.html#perf_hooks_performanceobserverentrylist_getentriesbyname_name_type

  // log entries named "measurement" with type "measure"
  list.getEntriesByName("measurement", "measure").forEach((entry) => {
    console.log(`${entry.name}'s duration: ${entry.duration}`)
    debugger
  })
})

// subscribe to various performance event types
observer.observe({
  // https://nodejs.org/api/perf_hooks.html#performanceentryentrytype
  // https://nodejs.org/api/perf_hooks.html#examples

  // entryTypes: ["mark", "measure", "function"],
  // entryTypes: ["mark", "measure"],
  entryTypes: ["measure"],
})

performance.mark('measurement start')
// a function for testing performance
function measurementTest(callback, iterations = 3) {
  return new Promise(resolve => {
    console.log(`doing some work... counting down ${iterations}...`)
    setTimeout(function() {
      if (iterations <= 0 ) {
        callback()
        resolve(true)
        return
      } else {
        resolve(measurementTest(callback, iterations - 1))
      }
    }, 1000)
  })
}

measurementTest(function() {
  performance.mark('measurement end')
  performance.measure('measurement', 'measurement start', 'measurement end')
}).then(() => {
  setTimeout(function() {
    debugger
    observer.disconnect()
  }, 10000)
})
```

<br/>
</details>

<details markdown="block">
  <summary>
    Debugging Worker thread - Node.js v14.21.3
  </summary>

```js
// worker-main.js
const { Worker } = require('worker_threads')

const runService = (workerData) => {
  return new Promise((resolve, reject) => {
    // https://nodejs.org/docs/v12.14.1/api/worker_threads.html#worker_threads_new_worker_filename_options
    // https://nodejs.org/docs/latest-v14.x/api/worker_threads.html#worker_threads_new_worker_filename_options
    const worker = new Worker('./worker-thread.js', {
      workerData,
      argv: ["--inspect-brk"],
      execArgv: ["--inspect-brk=0.0.0.0:9235"]
    })
    worker.on('message', resolve)
    worker.on('error', reject)
    worker.on('exit', (code) => {
      if (code !== 0)
        reject(new Error(`stopped with  ${code} exit code`))
    })
  })
}

const run = async () => {
  debugger
  const result = await runService('hello John Doe')
  debugger
  console.log(result)
}

run().catch(err => console.error(err))
```

```js
// worker-thread.js
const { workerData, parentPort } = require('worker_threads')
const inspector = require("inspector")

// https://github.com/nodejs/node/issues/26609
console.log(process.execArgv)
if ((process.execArgv || []).find(arg => arg.includes('--inspect'))) {
  console.log("--------xxxxxxx----------")
  inspector.open();
  inspector.waitForDebugger();
}

console.log("--------*******----------")
console.log(workerData)
console.log("--------*******----------")
debugger
parentPort.postMessage({ welcome: workerData })
```

```sh
node --inspect-brk worker-main.js
node inspect 0.0.0.0:9235
```

<br/>
</details>
