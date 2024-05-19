---
layout: default
title: Regex
nav_order: 10
parent: Browser JS
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/js/regex
---

# JavaScript Regex

- [further info]({% link docs/languages/regex/index.md %})

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

## Using Predefined Ranges

- digit character (`\d`)
- word character (`\w`)
- space character (`\s`)
- non-digit character (`\D`)
- non-word character (`\W`)
- non-space character (`\S`)
- single character (`.`)

## Regex With Quantifiers

The pattern `{n,m}` lets us match the preceding item at least `n` times and at most `m` times (`n` and `m` must be positive integers).

## Types of Quantifiers

There are six forms of quantifiers in regex: zero or more (`*`), one or more (`+`), zero or one (`?`), exactly `n` times `{n}`, at least `n` times `{n,}`, and from `n` to `m` times `{n,m}`.

## Capturing Group

A capturing group is a subpattern enclosed in parentheses. It captures the matched substring and stores it in a numbered group.

## Non-Capturing Group

A question mark (`?`) and a colon (`:`) right after the opening parentheses to create a non-capturing group:

```js
const re = /(\d{1,3})(?:st|nd|rd|th)/;
const str = "Tiger Woods sits 16th in the latest World Golf Ranking.";
const match = str.match(re);
if (match) {
  console.log("Player Rank: " + match[1]); // → Player Rank: 16
}
```

## Named Capturing Groups

```js
const re = /(?<hour>\d{2}):(?<min>\d{2}):(?<sec>\d{2})\s(?<period>\w{2})/;
const match = "09:30:00 AM".match(re);
if (match) {
  console.log("Hour: " + match.groups.hour); // → Hour: 09
  console.log("Minute: " + match.groups.min); // → Minute: 30
  console.log("Second: " + match.groups.sec); // → Second: 00
  console.log("Period: " + match.groups.period); // → Period: AM
}
```

__Named capturing groups__ is a syntax introduced in **ES2018**. A valid capturing group name must be an alphanumeric sequence starting with a letter. To avoid name collision with existing property names, JavaScript assigns all named groups to a separate object called `groups`.

## Special Replacement Characters

- `$<Name>`
  - ```js
    const phoneNum = "123-456-7890";
    function formatPhoneNumber(num) {
      // remove non-digits
      num = num.replace(/\D/g,"");
      // format the number
      num = num.replace(/(?<area>\d{3})(?<exchange>\d{3})(?<line>\d{4})/, "($<area>) $<exchange>-$<line>");
      return num;
    }
    formatPhoneNumber(phoneNum); // → "(123) 456-7890"
    ```
- `$n`
  - ```js
    const str = "cold & hot";
    const re = /(cold)\s&\s(hot)/;
    str.replace(re, "$2 & $1"); // → "hot & cold"
    ```
- `$&`
  - ```js
    const str = "FAT is a computer file system architecture";
    const re = /FAT/;
    str.replace(re, "$& (File Allocation Table)");
    // → "FAT (File Allocation Table) is a computer file system architecture"
    ```
- `$'`
  - ```js
    const str = "#3";
    const re = /#/;
    str.replace(re, "#$'"); // → "#33"
    ```
- `$\``
  - ```js
    const str = "1000 liter";
    const re = /liter/;
    str.replace(re, "liter = $`kg"); // → "1000 liter = 1000 kg"
    ```
- `$$`
  - ```js
    const str = "€700";
    const re = /€/;
    str.replace(re, "$$"); // → "$700"
    ```

## Using a Function as the Replacement Pattern

```js
const str = "3 Beds, 2.5 baths, 1,850 Sq. Ft";
// convert sqft to m2
function convertToM2(sqft) {
  // remove non-digit characters
  sqft = sqft.replace(/\D/g, "");
  // convert sqft to m2
  const m2 = sqft * 0.0929;
  // round to two decimal places and return it
  return m2.toFixed(2);
}
// match sqft in the string, have it converted to m2
// wrap parentheses around the result and append it to sqft
function appendM2ToSqft(str) {
  return str.replace(/\d+,?\d+\s(sqft|sq\.?\sft)/ig, (match) => {
    return `${match} (${convertToM2(match)} m2)`;
  });
}
appendM2ToSqft(str);
// → "3 Beds, 2.5 baths, 1,850 Sq. Ft (171.86 m2)"
```

The following table list all supported flags along with their corresponding properties.
Note that all these properties are read-only.

| Flag | Property |
| `d` | `hasIndices` |
| `g` | `global` |
| `i` | `ignoreCase` |
| `m` | `multiline` |
| `s` | `dotAll` |
| `u` | `unicode` |
| `y` | `sticky` |

Use the `flags` property to retrieve the flags of a regex object.


## Referencing a Matched String With the Backreference

```js
function dupWordRemover(str) {
  const re = /\b([-'\w]+)\s+\1\b/ig;
  return str.replace(re, "$1");
}
const str = "No no man has a a good enough memory to be a successful liar.";
dupWordRemover(str)
// → "No man has a good enough memory to be a successful liar."
```

Problem solved! Your text is now free of most duplicate words.

When there’s a capturing group in a pattern, the content inside the parentheses is bookmarked. A backreference provides a convenient way to reuse this content in the form of `\1`, `\2`, and so forth, where `\1` refers to the first captured group, `\2` refers to the second captured group, and so on.

```js
function dupWordRemover(str) {
  const re = /\b(?<dup>[-'\w]+)\s+\k<dup>\b/ig;
  return str.replace(re, "$1");
}
const str = "No no man has a a good enough memory to be a successful liar.";
dupWordRemover(str)
// → "No man has a good enough memory to be a successful liar."
```

This code achieves the same result as the solution in this recipe, except that it uses a named capturing group `(?<dup>[-'\w]+)`  and references the group with `\k<dup>`.

## Positive Lookahead

Use a positive lookahead denoted by `(?= ... )` to match the part of the string that you don’t want to include in the result:

```js
const str = `The Company posted a September quarter record revenue of $90.3 million, up 8 percent year over year.`;
const re = /\$(?=90\.3\smillion)/ig;
// replace only $ with €
str.replace(re, "€");
// → "The Company posted a September quarter record revenue \n
// of €90.3 million, up 8 percent year over year."
```

With the `(?= ... )` syntax, you require a pattern to appear after the regex match without including it in the match.

Lookaheads are non-capturing groups that allow us to match a pattern based on the substring that follows the pattern. For a positive lookahead match to be successful, it must match a pattern followed.

Notice the output of `match()` in the following code, which is only `“$”`

```js
const str = "$90.3 million";
const re = /\$(?=90\.3\smillion)/;
str.match(re)[0]; // → "$"
```

When using a lookahead, the subexpression is not included in the result. It also can’t be referenced with a backreference.

Don’t confuse capturing and matching. The positive lookahead assertion `(?= ... )` and the non-capturing group `(?: ... )` serve different purposes. While both don’t capture the substring they match, the non-capturing group includes the substring in the overall match, while the positive lookahead assertion does not.

The positive lookahead is denoted by the `(?= ... )` syntax, and it ensures that a pattern is followed by another pattern. On the other hand, negative lookahead is denoted by `(?! ...)`, and it guarantees that a pattern is not followed by another pattern.

## Negative Lookahead

A negative lookahead matches a pattern not followed by another pattern.

## Positive Lookbehind

```js
const re = /(?<=Question\s#\d{1,3}:\s).+?\./igs;
const str = `
Question #9: The Peloponnesian Wars were fought between __________.
Question #10: A ziggurat is __________.
`;
const questions = str.match(re);
console.log(questions);
// → [
//   "The Peloponnesian Wars were fought between __________.",
//   "A ziggurat is __________."
// ]
```

## Negative Lookbehind

```js
const re = /(?<!Dr\.\s)Smith/;
console.log(re.test("Dr. Smith"));  // → false
console.log(re.test("Mr. Smith"));  // → true
console.log(re.test("John Smith")); // → true
```

In addition to lookarounds, regular expressions provide several types of groups that are constructed using a pair of parentheses, with the opening parenthesis immediately followed by a question mark. For easier comparison, we’ve summarized the syntax for these groups in the following table. Keep this table bookmarked— it’s sure to come in handy:

| Metacharacter | Meaning |
| --- | --- |
| `( ... )` | Capturing group |
| `(?: ... )` | Non-capturing group |
| `(?= ... )` | Positive lookahead |
| `(?! ... )` | Negative lookahead |
| `(?<= ... )` | Positive lookbehind |
| `(?< ! ... )` | Negative lookbehind |


----

[^1]: [...](...)
