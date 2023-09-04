---
layout: default
title: React Server Side Rendering
nav_exclude: true
parent: Node JS
grand_parent: Programming Languages
permalink: /docs/languages/node-js/framework-react-ssr
---

__[back]({% link docs/languages/node-js/framework.md %})__

## React - Server-side Rendering (SSR)


<details markdown="block">
  <summary>
    code snippet
  </summary>

A sample from a DEV Community post [^1].

```bash
# init
npm init

# https://koajs.com/
npm install @koa/router@12.0.0
npm install koa@2.14.2
npm install koa-bodyparser@4.4.1

# https://react.dev/
npm install react@18.2.0
npm install react-dom@18.2.0

# https://babeljs.io/docs/babel-cli
# https://babeljs.io/docs/babel-register
# https://www.npmjs.com/package/babel-plugin-transform-react-jsx
npm install --save-dev @babel/cli@7.22.10
npm install --save-dev babel-plugin-transform-react-jsx@6.24.1
npm install --save-dev @babel/register@7.22.5
```

<details markdown="block">
  <summary>
    <sup>.babelrc</sup>
  </summary>
```json
{
  "plugins": ["transform-react-jsx"]
}
```
------
<!-- .babelrc -->
</details>

<details markdown="block">
  <summary>
    <sup>app.js</sup>
  </summary>
```js
const React = require('react');
const App = () => {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <title>React SSR</title>
      </head>
      <body>
        <div id="root">
          <h1>Hello, world!</h1>
        </div>
      </body>
    </html>
  );
}

module.exports = App
```
------
<!-- app.js -->
</details>

<details markdown="block">
  <summary>
    <sup>client.js</sup>
  </summary>
```js
const React = require('react');

const { hydrateRoot } = require('react-dom/client');

const App = require('./app');

hydrateRoot(document.getElementById('root'), <App />);
```
------
<!-- client.js -->
</details>

<details markdown="block">
  <summary>
    <sup>router-ssr.js</sup>
  </summary>
```js
const React = require('react');
const { renderToPipeableStream } = require('react-dom/server');
const App = require('./app');

const router = async (ctx) => {
  let didError = false;
  try {
    // Wraps into a promise to force Koa to wait for the render to finish
    return new Promise((_resolve, reject) => {
      const { pipe, abort } = renderToPipeableStream(
        <App />,
        {
          bootstrapModules: ['./client.js'],
          onShellReady() {
            ctx.respond = false;
            ctx.status = didError ? 500 : 200;
            ctx.set('Content-Type', 'text/html');
            pipe(ctx.res);
            ctx.res.end();
          },
          onShellError() {
            ctx.status = 500;
            abort();
            didError = true;
            ctx.set('Content-Type', 'text/html');
            ctx.body = '<!doctype html><p>Loading...</p><script src="clientrender.js"></script>';
            reject();
          },
          onError(error) {
            didError = true;
            console.error(error);
            reject();
          }
        },
      );

      setTimeout(() => {
        abort();
      }, 10_000);
    })
  } catch (err) {
    console.log(err);
    ctx.status = 500;
    ctx.body = 'Internal Server Error';
  }
};

module.exports = router
```
------
<!-- router-ssr.js -->
</details>

<details markdown="block">
  <summary>
    <sup>server.js</sup>
  </summary>
```js
const Router = require('@koa/router');
const Koa = require('koa');
const bodyparser = require('koa-bodyparser');
const ssr = require('./router-ssr')

const router = new Router();
const app = new Koa();

router.get('/app', ssr);

router.get('/(.*)', async (ctx) => {
  ctx.status = 200;
  ctx.body = 'OK';
});

app.use(bodyparser());
app.use(router.routes());
app.use(router.allowedMethods());

module.exports = app;
```
------
<!-- server.js -->
</details>

<details markdown="block">
  <summary>
    <sup>index.js</sup>
  </summary>
```js
require("@babel/register");

const http = require('http');
const _server = require('./server');

const currentHandler = _server.callback();
const server = http.createServer(_server.callback());

server.listen(4001, (error) => {
  console.log('listening...');
  console.error(error);
});
```
------
<!-- index.js -->
</details>

After creating all files, the repo should look like below.

```
 |
 |-/.babelrc
 |-/app.js
 |-/client.js
 |-/router-ssr.js
 |-/server.js
 |-/index.js
 |-/package-lock.json
 |-/package.json
```

```bash
# run the server
node index.js
```

```bash
# GET request
curl localhost:4001/app
```

</details>


----

[^1]: [Server-side Rendering (SSR) From Scratch with React](https://dev.to/woovi/server-side-rendering-ssr-from-scratch-with-react-19jm)
