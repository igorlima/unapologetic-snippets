---
layout: default
title: Pipe data into Node.js
nav_order: 5
parent: Node JS
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/node-js/pipe-data-into-nodejs
---

__[back]({% link docs/languages/node-js/index.md %})__

```js
// https://igorlima.github.io/unapologetic-snippets/docs/languages/node-js/debugging

// how to pipe data into a Node.js script
// NodeJS pipe input
echo '
bla
foo bar
fo foo
' | \
node -e '
  const readline = require("node:readline");
  async function main() {
    const rl = readline.createInterface({
      input: process.stdin,
    });

    process.stdout.write("\n")
    for await (const line of rl) {
      // process a line at a time
      process.stdout.write(`line: ${line}\n`);
    }
    process.stdout.write("\n")
  }
  main();
'
```
