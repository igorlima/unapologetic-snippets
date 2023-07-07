---
layout: default
title: Browser JS
nav_order: 4
parent: Programming Languages
has_children: true
permalink: /docs/languages/js
---


<br/>
<details markdown="block">
  <summary>
    JS Performance Timing
  </summary>

```js
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

// created profile can be seen in 'CPU PROFILES' part in
// 'Javascript Profiler' tab that can be opened from
// Developer Tools → Three dots → More tools.
console.profile("TEST :: console.profile")
measurementTest(function(){
  console.profileEnd("TEST :: console.profile")
})

// using console.timeStamp
// the console.timeStamp method adds a single marker to the browser's Performance tool
console.profile("profiling...")
console.timeStamp("TEST :: console.timeStamp started")
measurementTest(function(){
  console.timeStamp("TEST :: console.timeStamp ended")
  console.profileEnd("profiling...")
})

// using console.time
console.time("TEST :: console.time")
measurementTest(function(){
  console.timeEnd("TEST :: console.time")
})

// using perfomance.now
;(() => {
  const t0 = performance.now()
  measurementTest(function(){
    var t1 = performance.now()
    console.log(`it took ${t1 - t0} milliseconds.`)
  })
})();

// using performance.mark
performance.clearMarks()
performance.clearMeasures()
;(() => {
  performance.mark('measurement start')
  measurementTest(function(){
    performance.mark('measurement end')
  }).then(() => {
    performance.measure(
      'measurement',
      'measurement start',
      'measurement end'
    )
    const measures = performance.getEntriesByName('measurement')
    const [measure] = measures
    console.log(`performance timing was ${measure.duration}ms`)
  })
})();
```

<br/>
</details>


<br/>
<details markdown="block">
  <summary>
    JS Performance Timing - wrapper fn
  </summary>

```js
// sync
const perfduration = (metricName, func) => (...args) => {
  performance.mark(`${metricName} start`)
  const result = func(...args)
  performance.mark(`${metricName} end`)
  performance.measure(`${metricName}`, `${metricName} start`, `${metricName} end`)

  const measures = performance.getEntriesByName(`${metricName}`)
  // the reverse fn will get the lastest measurement via destructuring
  const [measure] = measures.reverse()
  console.log(`${metricName}: performance timing was ${measure.duration}ms`)
  return result
}


function measurementTest(iterations = 10) {
  for (let i = 0; i < iterations; i++) {
    console.log(i);
  }
}

perfduration('testing...', measurementTest)();
perfduration('testing...', measurementTest)(20);
perfduration('testing...', measurementTest)(100);
```

```js
// async
const perfduration = (metricName, func) => (...args) => {
  performance.mark(`${metricName} start`)
  return Promise.resolve(func(...args)).then((...data) => {
    performance.mark(`${metricName} end`)
    performance.measure(`${metricName}`, `${metricName} start`, `${metricName} end`)

    const measures = performance.getEntriesByName(`${metricName}`)
    // the reverse fn will get the lastest measurement via destructuring
    const [measure] = measures.reverse()
    console.log(`${metricName}: performance timing was ${measure.duration}ms`)
    return Promise.resolve(...data)
  })
}


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

perfduration('testing 1...', measurementTest)(function callback () {
  console.log('done 1')
}, 1);

perfduration('testing 2...', measurementTest)(function callback () {
  console.log('done 2')
}, 2);

perfduration('testing 4...', measurementTest)(function callback () {
  console.log('done 4')
}, 4);
```

<br/>
</details>

<br/>
