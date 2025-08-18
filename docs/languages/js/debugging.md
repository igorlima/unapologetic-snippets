---
layout: default
title: Debugging - Browser JS
nav_order: 4
parent: Browser JS
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/js/debugging
---

There are three main pillars of being a good software engineer: writing code, debugging code and reviewing code.

> “The art of debugging is figuring out what you really told your program to do rather than what you thought you told it to do.” _Andrew Singer_

> "Debugging is a separate skill from writing code. Writing code is creative, its like art. Debugging is like science, you need to formulate a small hypothesis and test it, then keep doing that till you've fixed the problem".

- other resources:
  - not yet available
- other tricks
  - `$0` - refers to the last element you clicked in the Elements tab
  - `copy(someVar)` - copies a variable’s value to your clipboard
  - `$_` - gives you the result of your last expression
  - edit any text on the webpage:
    - ```js
      document.designMode = "on"
      ```

<br/>
<details markdown="block">
  <summary>
    <code>debug</code> function
  </summary>

- before
  - ```js
    console.log('Before call');
    someFunction();
    console.log('After call');
    ```
- after
  - ```js
    debug(someFunction);
    ```

Why it’s awesome: <sup>[+](https://javascript.plainenglish.io/top-3-chrome-devtools-tricks-i-wish-i-knew-earlier-84ec870623eb) </sup>
1. It works even for functions from external libraries
2. You get full control - check variables, call stack, scope, everything
3. You don’t need to even know where the function is defined


<br/>
<!-- debug function -->
</details>


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
<!-- JS Performance Timing -->
</details>


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
<!-- JS Performance Timing - wrapper fn -->
</details>

<details markdown="block">
  <summary>
    Console for mobile browsers
  </summary>

A Console for Mobile Browsers — If you’re in a situation where you have no access to DevTools, you can add Eruda to your page and it provides a sort of virtual devtools you can use from any browser, including on mobile.
- [github](https://github.com/liriliri/eruda)

Browse it on your phone:
- [ilima wiki](http://wiki.ilima.xyz/eruda.html)
- [eruda home page](https://eruda.liriliri.io/)

In order to try it for different sites, execute the script below on browser address bar.

```js
javascript:(function () { var script = document.createElement('script'); script.src="https://cdn.jsdelivr.net/npm/eruda"; document.body.append(script); script.onload = function () { eruda.init(); } })();
```

```html
<!DOCTYPE html>
<!--
https://www.jsdelivr.com/package/npm/eruda
https://www.jsdelivr.com/package/npm/eruda-code

https://cdn.jsdelivr.net/npm/eruda@3.0.1/eruda.min.js
https://cdn.jsdelivr.net/npm/eruda-code@2.1.0/eruda-code.min.js
-->
<html>
  <head>
    <script src="//cdn.jsdelivr.net/npm/eruda"></script>
    <script src="//cdn.jsdelivr.net/npm/eruda-code"></script>
  </head>
<body>

<button type="button" onClick="eruda.show()">Show</button>
<button type="button" onClick="eruda.hide()">Hide</button>
<button type="button" onClick="eruda.util.evalCss.setTheme('Dark')">Dark</button>
<button type="button" onClick="eruda.util.evalCss.setTheme('Light')">Light</button>
<h1>Console for mobile browsers</h1>
<ul>
  <li><a href="https://github.com/liriliri/eruda">Github</a></li>
  <li><a href="https://eruda.liriliri.io/">Documentation</a></li>
</ul>

<br/>
<code>
javascript:(function () { var script = document.createElement('script'); script.src="https://cdn.jsdelivr.net/npm/eruda"; document.body.append(script); script.onload = function () { eruda.init(); } })();
</code>

<br/><br/>
<code>python3 -m http.server 8000</code>

<script>
eruda.init();
// eruda.position({x:0, y: 0});
// eruda.util.evalCss.setTheme('Dark')
// eruda.util.evalCss.setTheme('Light')
eruda.add(erudaCode);
</script>
</body>
</html>
```

<br/>
<!-- Console for mobile browsers -->
</details>

