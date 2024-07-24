package main

import "fmt"

func main() {
	func() {
		// SAMPLE A
		i := 0

	Next: // here, a label is declared.
		fmt.Println(i)
		i++
		if i < 5 {
			goto Next // execution jumps
		}
	}()

	func() {
		// SAMPLE B
		i := 0
	Next:
		if i >= 5 {
			goto Exit
		}
		// Create an explicit code block to
		// shrink the scope of k.
		{
			k := i + i
			fmt.Println(k)
		}
		i++
		goto Next
	Exit:
	}()

	func() {
		// SAMPLE C
		var k int // move the declaration of k here.
		i := 0
	Next:
		if i >= 5 {
			goto Exit
		}
		k = i + i
		fmt.Println(k)
		i++
		goto Next
	Exit:
	}()

	func() {
		// SAMPLE D
		// break and continue Statements With Labels
		FindSmallestPrimeLargerThan := func(n int) int {
		Outer:
			for n++; ; n++ {
				for i := 2; ; i++ {
					switch {
					case i*i > n:
						break Outer
					case n%i == 0:
						continue Outer
					}
				}
			}
			return n
		}

		for i := 90; i < 100; i++ {
			n := FindSmallestPrimeLargerThan(i)
			fmt.Print("The smallest prime number larger than ")
			fmt.Println(i, "is", n)
		}
	}()

	fmt.Println("vim-go")
}
