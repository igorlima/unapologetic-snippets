package main

import "fmt"

func main() {
	func() {
		//  pointer dereference examples:
		p0 := new(int)   // p0 points to a zero int value.
		fmt.Println(p0)  // (a hex address string)
		fmt.Println(*p0) // 0

		// x is a copy of the value at
		// the address stored in p0.
		x := *p0
		// Both take the address of x.
		// x, *p1 and *p2 represent the same value.
		p1, p2 := &x, &x
		fmt.Println(p1 == p2) // true

		fmt.Println(p0 == p1) // false
		p3 := &*p0            // <=> p3 := &(*p0) <=> p3 := p0
		// Now, p3 and p0 store the same address.
		fmt.Println(p0 == p3) // true
		*p0, *p1 = 123, 789
		fmt.Println(*p2, x, *p3) // 789 789 123

		fmt.Printf("%T, %T \n", *p0, x) // int, int
		fmt.Printf("%T, %T \n", p0, p1) // *int, *int
	}()

	func() {
		// why do we need pointers?
		func() {
			/*
				The double function is expected to modify the input argument by
				doubling it. However, it fails. Why? Because all value assignments,
				including function argument passing, are value copying in Go. What the
				double function modified is a copy (x) of variable a but not variable
				a.
			*/
			double := func(x int) {
				x += x
			}
			var a = 3
			double(a)
			fmt.Println(a) // 3
		}()
		func() {
			double := func(x *int) {
				*x += *x
				x = nil // the line is just for explanation purpose
			}
			var a = 3
			double(&a)
			fmt.Println(a) // 6
			p := &a
			double(p)
			fmt.Println(a, p == nil) // 12 false
		}()
	}()

	func() {
		a := int64(5)
		p := &a

		// The following two lines don't compile.
		/*
		   p++
		   p = (&a) + 8
		*/

		*p++
		fmt.Println(*p, a)   // 6 6
		fmt.Println(p == &a) // true

		*&a++
		*&*&a++
		**&p++
		*&*p++
		fmt.Println(*p, a) // 10 10
	}()

	func() {
		type MyInt int64
		type Ta *int64
		type Tb *MyInt

		// 4 nil pointers of different types.
		var pa0 Ta
		var pa1 *int64
		var pb0 Tb
		var pb1 *MyInt

		// The following 6 lines all compile okay.
		// The comparison results are all true.
		_ = pa0 == pa1
		_ = pb0 == pb1
		_ = pa0 == nil
		_ = pa1 == nil

		_ = pb0 == nil
		_ = pb1 == nil

		// None of the following 3 lines compile ok.
		/*
		   _ = pa0 == pb0
		   _ = pa1 == pb1
		   _ = pa0 == Tb(nil)
		*/
	}()

	fmt.Println("vim-go")
}
