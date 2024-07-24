package main

import "fmt"

/*
At run time, each init function will be (sequentially) invoked once and only
once (before invoking the main entry function).
*/

func init() {
	fmt.Println("hi,", bob)
}

func main() {
	fmt.Println("bye")

	/*
		f := func() int {
			return z + y
		}
		g := func() int {
			return y / 2
		}
		var (
			w       = x
			x, y, z = f(), 123, g()
		)
	*/
}

func init() {
	fmt.Println("hello,", smith)
}

func titledName(who string) string {
	return "Mr. " + who
}

var bob, smith = titledName("Bob"), titledName("Smith")

/*
The output of this program:
```
hi, Mr. Bob
hello, Mr. Smith
bye
```
*/

/*
Go runtime will try to initialize package-level variables in a package by their
declaration order, but a package-level variable will be initialized after all
of its depended variables for sure.  For example, in the following code
snippet, the initializations the four package-level variables happen in the
order y, z, x, and w.
*/
func f() int {
	return z + y
}
func g() int {
	return y / 2
}

var (
	w       = x
	x, y, z = f(), 123, g()
)
