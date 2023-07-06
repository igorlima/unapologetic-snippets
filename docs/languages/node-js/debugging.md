---
layout: default
title: Debugging
nav_order: 5
parent: Node JS
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/node-js/debugging
---

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

## through the terminal

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
node inspect 0.0.0.0:9235 node-testing.js
```

<details markdown="block">
  <summary>
    Performance Timing - Node.js v14.21.3
  </summary>

```js
// Node.js v14.21.3
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
  debugger
  observer.disconnect()
})
```

<br/>
</details>
<br/>
