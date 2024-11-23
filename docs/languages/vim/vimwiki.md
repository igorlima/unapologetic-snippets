---
layout: default
title: Vim Wiki
nav_order: 2
parent: Vim
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/vim/vimwiki
---

# Vim Wiki

- [go back]({% link docs/languages/vim/index.md %})


- [vimwiki doc](https://raw.githubusercontent.com/vimwiki/vimwiki/master/doc/vimwiki.txt) _<sup>a personal wiki for Vim</sup>_


<details markdown="block"><summary>advanced <strong>Search</strong> and <strong>Replace</strong> in Vim</summary>

- <details markdown="block"><summary>replacing in multiple files</summary>

  **Example Workflow:**

  Here's a concise example of the entire process:

  ```vim
  :grep foo **/*.js      " Search for 'foo' in JavaScript files
  :cdo s/foo/bar/gc      " Replace 'foo' with 'bar' (with confirmation)
  :cfdo up               " Save changes
  ```

  This guide provides a friendly walkthrough for performing powerful search and replace operations across multiple files in Vim.  We'll break down the process into smaller, manageable steps.

  1. <details markdown="block"><summary>populating the Quickfix list:</summary>

     The first step is to identify all occurrences of your search pattern.  Several commands can achieve this by populating Vim's "quickfix list." Here are a few options:

     * **`:grep`:**  A built-in command for searching.  For example, to search for "foo" in all JavaScript files within the current directory and its subdirectories, use:

        ```vim
        :grep foo **/*.js
        ```

     * **`:vimgrep`:** Similar to `:grep`, but offers more advanced pattern matching capabilities.

     You can view the quickfix list with `:cwindow`.
     </details>

  2. <details markdown="block"><summary>performing the replacement:</summary>

     Once the quickfix list is populated, you can perform the replacement using either `:cdo` (operate on each line in the quickfix list) or `:cfdo` (operate on each file containing matches).

     * **`:cdo` (Line-wise Replacement):**

        ```vim
        :cdo s/foo/bar/gc
        ```
        This replaces "foo" with "bar" on each line listed in the quickfix list, prompting for confirmation (`c`) before each substitution.

     * **`:cfdo` (File-wise Replacement):**

        ```vim
        :cfdo %s/foo/bar/gc
        ```
        This replaces all instances of "foo" with "bar" within each *file* listed in the quickfix list, also prompting for confirmation (`c`) before each substitution.  The `%` ensures the substitution applies to the entire file.

     * **Skipping Confirmation:**  If you're confident in your replacement, you can omit the `c` flag for automatic substitution. See `:help :s_flags` for more details on substitution flags.
     </details>

  3. <details markdown="block"><summary>saving changes:</summary>

     After performing the replacements, save the changes to disk using `:cfdo update` (or its abbreviated form, `:cfdo up`):

     ```vim
     :cfdo update
     ```
     This command writes only the modified files to disk.

     </details>

  - <details markdown="block"><summary>deleting lines with <code>:global</code>:</summary>

    You can also use the quickfix list to delete lines matching a pattern.  For example, to delete all lines containing "my-grep-pattern" within the files in the quickfix list:

    ```vim
    :cfdo %g/my-grep-pattern/d
    ```
    </details>

  <details markdown="block"><summary><strong>Helpful Resources</strong></summary>

  For more in-depth information, check out:
  - Vim's built-in help:
    - `:help :cfdo`
    - `:help :cdo`
  - online resources:
    - [StackOverflow - Find and replace all instances of specific string in multiple files in vim](https://stackoverflow.com/questions/70003193/find-and-replace-all-instances-of-specific-string-in-multiple-files-in-vim)
    - [beezwax - Advanced Search and Replace with Vim](https://blog.beezwax.net/advanced-search-and-replace-with-vim/)
  </details>

  <!-- Replacing in Multiple Files -->
  ---------
  </details>

<!--
  Advanced Search and Replace in Vim
-->
</details>
