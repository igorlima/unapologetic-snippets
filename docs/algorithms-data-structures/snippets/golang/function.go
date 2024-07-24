package main

import "fmt"

/*
Parameters and results are treated as local variables within their
corresponding function bodies.
*/

/*
In the below example, each parameter and result declaration is composed of a
name and a type (the type follows the name).
*/
func SquaresOfSumAndDiff1(a int64, b int64) (s int64, d int64) {
	x, y := a+b, a-b
	s = x * x
	d = y * y
	return // <=> return s, d
}

/*
When all the results in a function declaration are anonymous, then, within the
corresponding function body, the return keyword must be followed by a sequence
of return values, each of them corresponds to a result declaration of the
function declaration.
*/
func SquaresOfSumAndDiff2(a int64, b int64) (int64, int64) {
	return (a + b) * (a + b), (a - b) * (a - b)
}

/*
If the type portions of some successive parameters in a parameter declaration
list are identical, then these parameters could share the same type portion in
the parameter declaration list.  The same is for result declaration lists.
*/
func SquaresOfSumAndDiff3(a, b int64) (s, d int64) {
	return (a + b) * (a + b), (a - b) * (a - b)
	// The above line is equivalent
	// to the following line.
	/*
	   s = (a+b) * (a+b); d = (a-b) * (a-b); return
	*/
}

/*
Go doesn't support default parameter values.
The initial value of each result is the zero value of its type.
*/
func f() (x int, y bool) {
	println(x, y) // 0 false
	return
}

func main() {
	println(SquaresOfSumAndDiff1(3, 4)) // 49 1
	println(SquaresOfSumAndDiff2(3, 4)) // 49 1
	println(SquaresOfSumAndDiff3(3, 4)) // 49 1
	println(f())                        // 0 false

	/* ANONYMOUS FUNCTIONS */
	/*
		Go supports anonymous functions.  The definition of an anonymous function
		is almost the same as a function declaration, except there is no function
		name portion in the anonymous function definition.
	*/
	/*
	 An anonymous function can be called right after it is defined
	*/

	// this anonymous function has no parameters but has two results.
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // call it - no arguments are needed.
	// the following anonymous function have no results.
	func(a, b int) {
		// the following line prints: a*a + b*b = 25
		println("a*a + b*b =", a*a+b*b)
	}(x, y) // pass argument x and y to parameter a and b.
	func(x int) {
		// the parameter x shadows the outer x.
		// the following line prints: x*x + y*y = 32
		println("x*x + y*y =", x*x+y*y)
	}(y) // pass argument y to parameter x.
	func() {
		// the following line prints: x*x + y*y = 25
		println("x*x + y*y =", x*x+y*y)
	}() // no arguments are needed.

	fmt.Println("vim-go")
}
