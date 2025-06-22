---
layout: default
title: Programming Languages
nav_order: 3
has_children: true
has_toc: false
permalink: /docs/languages
---

# Programming Languages

To make it as easy as possible to write documentation in plain Markdown, most UI components are styled using default Markdown elements with few additional CSS classes needed.
{: .fs-6 .fw-300 }


```html
<!-- make a piece of code collapsible -->
<details markdown="block">
  <summary>
    a full code snippet
  </summary>


<br/>
</details>
<br/>
```
```html
<!-- § reference sample § -->
[ref to ...](#ref-8c678ce8-89c1-45e1-834f-294729cb7d8e)
[another ref to ...]({% link docs/languages/languages.md#ref-8c678ce8-89c1-45e1-834f-294729cb7d8e %})

<!-- either -->
<details markdown="block">
  <summary>
    sample title
    <a href="#ref-8c678ce8-89c1-45e1-834f-294729cb7d8e">§</a>
  </summary>

  <a id="ref-8c678ce8-89c1-45e1-834f-294729cb7d8e"></a>
</details>

<!-- or -->
<a id="ref-8c678ce8-89c1-45e1-834f-294729cb7d8e" href="#ref-8c678ce8-89c1-45e1-834f-294729cb7d8e">§</a>
```
