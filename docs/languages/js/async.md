---
layout: default
title: Async Tasks
nav_order: 10
parent: Browser JS
grand_parent: Programming Languages
has_children: false
permalink: /docs/languages/js/async
---

# Asynchronous JavaScript

- [further info](https://pragprog.com/titles/fkajs/modern-asynchronous-javascript/)

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>


<br/>
<details markdown="block">
  <summary>
    Creating a Custom Iterator
  </summary>

To be classified as an iterable, an object must come with a `Symbol.iterator` property and specify the return value for each iteration.

```js
const collection = {
  a: 10,
  b: 20,
  c: 30,
  [Symbol.iterator]() {
    let i = 0;
    const values = Object.keys(this);
    return {
      next: ()=>{
        return {
          value: this[values[i++]],
          done: i > values.length
        }
      }
    }
  }
}
const iterator = collection[Symbol.iterator]();
console.log(iterator.next()); // ⇒ {value: 10, done: false}
console.log(iterator.next()); // ⇒ {value: 20, done: false}
console.log(iterator.next()); // ⇒ {value: 30, done: false}
console.log(iterator.next()); // ⇒ {value: undefined, done: true}
for (const value of collection) {
  console.log(value)
  // logs:
  // ⇒ 10
  // ⇒ 20
  // ⇒ 30
}
```
</details>

<details markdown="block">
  <summary>
    Creating a Custom Asynchronous Iterator
  </summary>

```js
const collection = {
  a: 10,
  b: 20,
  c: 30,
  [Symbol.asyncIterator]() {
    const keys = Object.keys(this);
    let i = 0;
    return {
      next: ()=>{
        return new Promise((resolve,reject)=>{
          setTimeout(()=>{
            resolve({
              value: this[keys[i++]],
              done: i > keys.length
            });
          }
          , 1000);
        });
      }
    }
  }
}
const iterator = collection[Symbol.asyncIterator]();
iterator.next().then(result => {
  console.log(result); // ⇒ {value: 10, done: false}
});
iterator.next().then(result => {
  console.log(result); // ⇒ {value: 20, done: false}
});
iterator.next().then(result => {
  console.log(result); // ⇒ {value: 30, done: false}
});
iterator.next().then(result => {
  console.log(result); // ⇒ {value: undefined, done: true}
});
```
</details>


<details markdown="block">
  <summary>
    Async Iterables with <code>for...await...of</code>
  </summary>

```js
const arr = [
  {"firstName":"John", "lastName":"Doe"},
  {"firstName":"Anna", "lastName":"Smith"},
  {"firstName":"Peter", "lastName":"Jones"},
];
arr[Symbol.asyncIterator] = function() {
  let i = 0;
  return {
    async next() {
      if (i === arr.length) {
        return {
          done: true
        };
      }
      const obj = arr[i++];
      const response = await Promise.resolve(obj);
      return {
        value: response,
        done: false
      };
    }
  };
};

(async function() {
  for await (const obj of arr) {
    console.log(obj.firstName)
    // logs:
    // ⇒ John
    // ⇒ Peter
    // ⇒ Anna
  }
})();
```
</details>


<details markdown="block">
  <summary>
    Using a Generator to Define a Custom Iterator
  </summary>

```js
const collection = {
  a: 10,
  b: 20,
  c: 30,
  [Symbol.iterator]: function*() {
    for (let key in this) {
      yield this[key]
    }
  }
}

const iterator = collection[Symbol.iterator]();
console.log(iterator.next()); // ⇒ {value: 10, done: false}
console.log(iterator.next()); // ⇒ {value: 20, done: false}
console.log(iterator.next()); // ⇒ {value: 30, done: false}
console.log(iterator.next()); // ⇒ {value: undefined, done: true}
```
</details>

<sup>Generators enhance the process of creating iterables by providing an iterative algorithm. An async generator is similar to a sync generator except that it returns a promise rather than a plain object. Use a generator function when you don’t want to manipulate the state-maintaining behavior of the object.</sup>

<details markdown="block">
  <summary>
    Creating an Asynchronous Generator
  </summary>

```js
const arr = [{
  "firstName": "John",
  "lastName": "Doe"
}, {
  "firstName": "Anna",
  "lastName": "Smith"
}, {
  "firstName": "Peter",
  "lastName": "Jones"
}];

arr[Symbol.asyncIterator] = async function*() {
  let i = 0;
  for (const obj of this) {
    const response = await Promise.resolve(obj)
    yield response
  }
}

const iterator = arr[Symbol.asyncIterator]()
iterator.next().then(result => {
  console.log(result.value.firstName); // ⇒ John
})
iterator.next().then(result => {
  console.log(result.value.firstName); // ⇒ Peter
})
iterator.next().then(result => {
  console.log(result.value.firstName); // ⇒ Anna
})
```
</details>


<details markdown="block">
  <summary>
    Iterating over Paginated Data
  </summary>

One situation we want to use asynchronous iteration over synchronous is when working with web APIs that provide paginated data. By using an asyn- chronous iterator, we can seamlessly make multiple network requests and iterate over the results.

**an asynchronous generator and program it to handle the pagination**

```js
// data:text/html,<html contenteditable>
async function* generator(repo) {
  for (;;) {
    const response = await fetch(repo)
    const data = await response.json()
    for (let commit of data) {
      yield commit
    }
    const link = response.headers.get('Link')
    repo = /<(.*?)>; rel="next"/.exec(link)?. [1];
    if (repo === undefined) {
      break
    }
  }
}

async function getCommits(repo) {
  let i = 0
  for await (const commit of generator(repo)) {
    console.log(commit)
    if (++i >= 10) {
      break
    }
  }
}
getCommits('https://api.github.com/repos/tc39/proposal-temporal/commits')
```
</details>

<sub>Choose `Promise.all` when you want to fail fast or only care about fulfillment values. Opt for `Promise.allSettled` when you want to know the outcome of all promises and handle both fulfillments and rejections in a comprehensive way.</sub>


<details markdown="block">
  <summary>
    Using the <code>Promise.any()</code> Method
  </summary>

Remember, `Promise.any()` only uses the result of the first promise that fulfills, so the other result is ignored.

```js
const apis = ['https://eloux.com/todos/1', 'https://jsonplaceholder.typicode.com/todos/1'];
async function fetchData(api) {
  const response = await fetch(api);
  if (response.ok) {
    return response.json();
  } else {
    return Promise.reject(new Error('Request failed'));
  }
}
function getData() {
  return Promise.any([fetchData(apis[0]), fetchData(apis[1])]);
}
getData().then((response)=>console.log(response.title));
```
</details>


<details markdown="block">
  <summary>
    Comparing <code>Promise.race()</code> to <code>Promise.any()</code>
  </summary>

When it come to rejection, `Promise.race()` is completely different: it settles as soon as one of the given promises rejects. In other words, while `Promise.any()` rejects if all of the given promises reject, `Promise.race()` rejects if the first promise that settles is rejected.

```js
function loadFromCache() {
  const data = {
    "userId": 1,
    "id": 1,
    "title": "delectus aut autem",
    "completed": false
  };
  return new Promise((resolve)=>{
    resolve(data);
  }
  )
}
function fetchData() {
  const timeOut = 2000;
  // two seconds
  const cache = loadFromCache().then((data)=>{
    return new Promise((resolve)=>{
      setTimeout(()=>{
        resolve(data);
      }
      , timeOut);
    }
    );
  }
  );
  const freshData = fetch('https://jsonplaceholder.typicode.com/todos/1');
  return Promise.race([freshData, cache]);
}
fetchData().then(async (response)=>{
  console.log(response.json ? await response.json() : response);
}).catch((error)=>{
  console.error(error);
});
```
</details>


<details markdown="block">
  <summary>
    Canceling Pending Async Requests
  </summary>

- <details markdown="block">
    <summary>
      Canceling Async Tasks After a Period of Time
    </summary>
  
  ```js
  const controller = new AbortController();
  const signal = controller.signal;
  fetch('https://eloux.com/todos/1', {
    signal
  }).then(response=>{
    return response.json();
  }).then(response=>{
    console.log(response);
  });
  setTimeout(()=>controller.abort(), 2000);
  ```
  ```js
  signal.addEventListener('abort', () => {
    console.log(signal.aborted);
  });
  ```
  ```js
  function fetchWithTimeout(url, settings, timeout) {
    // If the timeout argument doesn't exists
    if (timeout === undefined) {
      return fetch(url, settings);
    }
    // if timeout isn't an integer, throw an error
    if (!Number.isInteger(timeout)) {
      throw new TypeError('The third argument is not an integer')
    }
    const controller = new AbortController();
    setTimeout(()=>controller.abort(), timeout);
    settings.signal = controller.signal;
    return fetch(url, settings);
  }
  ```
  ```js
  signal.addEventListener('abort', () => {
    console.log(signal.aborted);
  });
  ```
  </details>
- <details markdown="block">
    <summary>
      Removing Multiple Event Listeners
    </summary>

  ```js
  // data:text/html,<html contenteditable><body> <div class="container">Mouse over me!</div> </body>
  const container = document.querySelector('.container');
  const controller = new AbortController();
  const signal = controller.signal;
  function sayHello() {
    container.textContent = 'Hello';
  }
  function sayBye() {
    container.textContent = 'Bye!';
  }
  function depress() {
    container.style.backgroundColor = 'aqua';
  }
  function release() {
    container.style.backgroundColor = 'transparent';
  }
  container.addEventListener('mouseenter', sayHello, {signal});
  container.addEventListener('mouseout', sayBye, {signal});
  container.addEventListener('mousedown', depress, {signal});
  container.addEventListener('mouseup', release, {signal});
  ```
  </details>
- <details markdown="block">
    <summary>
      Making a User-Cancelable Async Request
    </summary>
  ```js
  // data:text/html,<html contenteditable><body> <image class="image"> <span class="result"></span> <button class="loadBtn">Load Photo</button> <button class="abortBtn" disabled="disabled">Cancel Loading</button> </body>
  const loadBtn = document.querySelector('.loadBtn');
  const abortBtn = document.querySelector('.abortBtn');
  const image = document.querySelector('.image');
  const result = document.querySelector('.result');
  const controller = new AbortController();
  // abort the request
  
  abortBtn.addEventListener('click', ()=>controller.abort());
  
  // load the image
  loadBtn.addEventListener('click', async()=>{
    loadBtn.disabled = true;
    abortBtn.disabled = false;
    result.textContent = 'Loading...';
    try {
      const response = await fetch(`https://upload.wikimedia.org/wikipedia/commons/a/a3/Kayakistas_en_Glaciar_Grey.jpg`, {
        signal: controller.signal
      });
      const blob = await response.blob();
      image.src = URL.createObjectURL(blob);
      // remove the "Loading.." text
      result.textContent = '';
    } catch (err) {
      if (err.name === 'AbortError') {
        result.textContent = 'Request successfully canceled';
      } else {
        result.textContent = 'An error occurred!'
        console.error(err);
      }
    }
  
    loadBtn.disabled = false;
    abortBtn.disabled = true;
  });
  ```
  </details>
- <details markdown="block">
    <summary>
      Aborting Multiple Fetch Requests with One Signal
    </summary>
  ```js
  // data:text/html,<html contenteditable><body> <div class="gallery"></div> <span class="result"></span> <button class="loadBtn">Load Photos</button> <button class="abortBtn" disabled="disabled">Cancel Loading</button> </body>
  const loadBtn = document.querySelector('.loadBtn');
  const abortBtn = document.querySelector('.abortBtn');
  const gallery = document.querySelector('.gallery');
  const result = document.querySelector('.result');
  
  const controller = new AbortController();
  
  const urls = [
    `https://upload.wikimedia.org/wikipedia/commons/thumb/7/70/Por_do_Sol_em_Baixa_Grande.jpg/320px-Por_do_Sol_em_Baixa_Grande.jpg`,
    `https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Zebrasoma_flavescens_Luc_Viatour.jpg/320px-Zebrasoma_flavescens_Luc_Viatour.jpg`,
    `https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Domestic_goat_kid_in_capeweed.jpg/320px-Domestic_goat_kid_in_capeweed.jpg`
  ];
  
  abortBtn.addEventListener('click', () => controller.abort());
  loadBtn.addEventListener('click', async () => {
    loadBtn.disabled = true;
    abortBtn.disabled = false;
    result.textContent = 'Loading...';
    const tasks = urls.map(url => fetch(url, {signal: controller.signal}));
    try{
      const response = await Promise.all(tasks);
      response.forEach(async (r) => {
        const img = document.createElement('img');
        const blob = await r.blob();
        img.src = URL.createObjectURL(blob); gallery.appendChild(img);
      });
      result.textContent = '';
    } catch (err) {
      if (err.name === 'AbortError') {
        result.textContent = 'Request successfully canceled';
      } else {
        result.textContent = 'An error occurred!'
        console.error(err);
      }
    }
    loadBtn.disabled = false;
    abortBtn.disabled = true;
  });
  ```
  </details>

</details>

<details markdown="block">
  <summary>
    Importing libraries
  </summary>

```js
const CDNs = [
  import('https://cdnjs.cloudflare.com/ajax/libs/lodash.js/4.17.21/lodash.js'),
  import('https://cdn.jsdelivr.net/npm/lodash@4.17.21/lodash.min.js'),
];
const lodash = await Promise.any(CDNs);
console.log(_.camelCase('bla bla bla'))
```
</details>

----

[^1]: [...](...)
