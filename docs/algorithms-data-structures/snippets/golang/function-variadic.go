package main

import "fmt"

func main() {
	func() {
		// Sum and return the input numbers.
		Sum := func(values ...int64) (sum int64) {
			// The type of values is []int64.
			sum = 0
			for _, v := range values {
				sum += v
			}
			return
		}

		// An inefficient string concatenation function.
		Concat := func(sep string, tokens ...string) string {
			// The type of tokens is []string.
			r := ""
			for i, t := range tokens {
				if i != 0 {
					r += sep

				}
				r += t
			}
			return r
		}

		println(Sum(1, 2, 3, 4, 5))
		println(Concat(" ", "a", "b", "c", "d"))
	}()

	func() {
		Sum := func(values ...int64) (sum int64) {
			sum = 0
			for _, v := range values {
				sum += v
			}
			return
		}
		a0 := Sum()
		a1 := Sum(2)
		a3 := Sum(2, 3, 5)
		// The above three lines are equivalent to
		// the following three respective lines.
		b0 := Sum([]int64{}...) // <=> Sum(nil...)
		b1 := Sum([]int64{2}...)
		b3 := Sum([]int64{2, 3, 5}...)
		fmt.Println(a0, a1, a3) // 0 2 10
		fmt.Println(b0, b1, b3) // 0 2 10
	}()

	func() {

		Concat := func(sep string, tokens ...string) (r string) {
			for i, t := range tokens {
				if i != 0 {
					r += sep
				}
				r += t
			}
			return
		}

		tokens := []string{"Go", "C", "Rust"}
		// manner 1
		langsA := Concat(",", tokens...)
		// manner 2
		langsB := Concat(",", "Go", "C", "Rust")
		fmt.Println(langsA == langsB) // true
	}()

	func() {
		HalfAndNegative := func(n int) (int, int) {
			return n / 2, -n
		}
		AddSub := func(a, b int) (int, int) {
			return a + b, a - b
		}
		Dummy := func(values ...int) {}

		// These lines compile okay.
		AddSub(HalfAndNegative(6))
		AddSub(AddSub(AddSub(7, 5)))
		AddSub(AddSub(HalfAndNegative(6)))
		Dummy(HalfAndNegative(6))
		_, _ = AddSub(7, 5)

		// The following lines fail to compile.
		/*
			_, _, _ = 6, AddSub(7, 5)
			Dummy(AddSub(7, 5), 9)
			Dummy(AddSub(7, 5), HalfAndNegative(6))
		*/
	}()

	func() {
		Double := func(n int) int {
			return n + n
		}
		Apply := func(n int, f func(int) int) int {
			return f(n) // the type of f is "func(int) int"
		}

		fmt.Printf("%T\n", Double) // func(int) int

		// Double = nil // error: Double is immutable.

		var f func(n int) int // default value is nil.
		f = Double
		g := Apply            // let compile deduce the type of g
		fmt.Printf("%T\n", g) // func(int, func(int) int) int

		fmt.Println(f(9))         // 18
		fmt.Println(g(6, Double)) // 12
		fmt.Println(Apply(6, f))  // 12
	}()

	func() {
		// This function returns a function (a closure).
		isMultipleOfX := func(x int) func(int) bool {
			return func(n int) bool {
				return n%x == 0
			}
		}

		var isMultipleOf3 = isMultipleOfX(3)
		var isMultipleOf5 = isMultipleOfX(5)
		fmt.Println(isMultipleOf3(6))  // true
		fmt.Println(isMultipleOf3(8))  // false
		fmt.Println(isMultipleOf5(10)) // true
		fmt.Println(isMultipleOf5(12)) // false

		isMultipleOf15 := func(n int) bool {
			return isMultipleOf3(n) && isMultipleOf5(n)
		}
		fmt.Println(isMultipleOf15(32)) // false
		fmt.Println(isMultipleOf15(60)) // true
	}()

	fmt.Println("vim-go")
}
