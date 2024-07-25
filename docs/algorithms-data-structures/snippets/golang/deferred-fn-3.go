package main

import "fmt"

func Triple(n int) (r int) {
	defer func() {
		r += n // modify the return value
	}()
	return n + n // <=> r = n + n; return
}

func main() {
	// deferred function calls can modify the
	// named return results of nesting functions

	fmt.Println(Triple(5)) // 15
	fmt.Println("vim-go")
}
