package main

import "fmt"

func main() {
	func() {

		const Size = 32

		type Person struct {
			name string
			age  int
		}

		/* Array types */

		_ = [5]string{}
		_ = [Size]int{}
		// Element type is a slice type: []byte
		_ = [16][]byte{}
		// Element type is a struct type: Person
		_ = [100]Person{}

		/* Slice types */
		_ = []bool{}
		_ = []int64{}

		// Element type is a map type: map[int]bool
		_ = []map[int]bool{}
		// Element type is a pointer type: *int
		_ = []*int{}

		/* Map types */

		_ = map[string]int{}
		_ = map[int]bool{}
		// Element type is an array type: [6]string
		_ = map[int16][6]string{}
		// Element type is a slice type: []string
		_ = map[bool][]string{}
		// Element type is a pointer type: *int8,
		// and key type is a struct type.
		_ = map[struct{ x int }]*int8{}
	}()

	func() {
		// An array value containing four bool values.
		_ = [4]bool{false, true, true, false}

		// A slice value which contains three words.
		_ = []string{"break", "continue", "fallthrough"}

		// A map value containing some key-value pairs.
		_ = map[string]int{"C": 1972, "Python": 1991, "Go": 2009}
	}()

	func() {
		// variants for array and slice composite literals

		// The followings slice composite literals
		// are equivalent to each other.
		_ = []string{"break", "continue", "fallthrough"}
		_ = []string{0: "break", 1: "continue", 2: "fallthrough"}
		_ = []string{2: "fallthrough", 1: "continue", 0: "break"}
		_ = []string{2: "fallthrough", 0: "break", "continue"}

		// The followings array composite literals
		// are equivalent to each other.
		_ = [4]bool{false, true, true, false}
		_ = [4]bool{0: false, 1: true, 2: true, 3: false}
		_ = [4]bool{1: true, true}
		_ = [4]bool{2: true, 1: true}
		_ = [...]bool{false, true, true, false}
		_ = [...]bool{3: false, 1: true, true}
	}()

	func() {
		pm := &map[string]int{"C": 1972, "Go": 2009}
		ps := &[]string{"break", "continue"}
		pa := &[...]bool{false, true, true, false}
		fmt.Printf("%T\n", pm) // *map[string]int
		fmt.Printf("%T\n", ps) // *[]string
		fmt.Printf("%T\n", pa) // *[4]bool
	}()

	func() {
		func() {
			// Nested Composite Literals

			// A slice value of a type whose element type is
			// *[4]byte. The element type is a pointer type
			// whose base type is [4]byte. The base type is

			// an array type whose element type is "byte".
			var heads = []*[4]byte{
				&[4]byte{'P', 'N', 'G', ' '},
				&[4]byte{'G', 'I', 'F', ' '},
				&[4]byte{'J', 'P', 'E', 'G'},
			}
			println(heads)
			println(heads[0], heads[1], heads[2])
		}()
		func() {
			// can be simplified to
			var heads = []*[4]byte{
				{'P', 'N', 'G', ' '},
				{'G', 'I', 'F', ' '},
				{'J', 'P', 'E', 'G'},
			}
			println(heads)
			println(heads[0], heads[1], heads[2])
		}()
	}()

	func() {
		type language struct {
			name string
			year int
		}
		func() {
			var _ = [...]language{
				language{"C", 1972},
				language{"Python", 1991},
				language{"Go", 2009},
			}
		}()
		func() {
			// can be simplified to
			var _ = [...]language{
				{"C", 1972},
				{"Python", 1991},
				{"Go", 2009},
			}
		}()
	}()

	func() {
		type LangCategory struct {
			dynamic bool
			strong  bool
		}
		func() {
			// A value of map type whose key type is
			// a struct type and whose element type
			// is another map type "map[string]int".
			var _ = map[LangCategory]map[string]int{
				LangCategory{true, true}: map[string]int{
					"Python": 1991,
					"Erlang": 1986,
				},
				LangCategory{true, false}: map[string]int{
					"JavaScript": 1995,
				},
				LangCategory{false, true}: map[string]int{
					"Go":   2009,
					"Rust": 2010,
				},
				LangCategory{false, false}: map[string]int{
					"C": 1972,
				},
			}
		}()

		func() {
			// can be simplified to
			var _ = map[LangCategory]map[string]int{
				{true, true}: {
					"Python": 1991,
					"Erlang": 1986,
				},
				{true, false}: {
					"JavaScript": 1995,
				},
				{false, true}: {
					"Go":   2009,
					"Rust": 2010,
				},
				{false, false}: {
					"C": 1972,
				},
			}
		}()
	}()

	func() {
		var a [5]int
		fmt.Println(len(a), cap(a)) // 5 5
		var s []int
		fmt.Println(len(s), cap(s)) // 0 0
		s, s2 := []int{2, 3, 5}, []bool{}
		fmt.Println(len(s), cap(s))   // 3 3
		fmt.Println(len(s2), cap(s2)) // 0 0
		var m map[int]bool
		fmt.Println(len(m)) // 0
		m, m2 := map[int]bool{1: true, 0: false}, map[int]int{}
		fmt.Println(len(m), len(m2)) // 2 0
	}()

	func() {
		a := [3]int{-1, 0, 1}
		s := []bool{true, false}
		m := map[string]int{"abc": 123, "xyz": 789}
		fmt.Println(a[2], s[1], m["abc"])     // retrieve
		a[2], s[1], m["abc"] = 999, true, 567 // modify
		fmt.Println(a[2], s[1], m["abc"])     // retrieve

		n, present := m["hello"]
		fmt.Println(n, present, m["hello"]) // 0 false 0

		n, present = m["abc"]
		fmt.Println(n, present, m["abc"]) // 567 true 567
		m = nil
		fmt.Println(m["abc"]) // 0

		// The two lines fail to compile.
		/*
		   _ = a[3]  // index 3 out of bounds
		   _ = s[-1] // index must be non-negative
		*/

		return
		// Each of the following lines can cause a panic.
		_ = a[n]         // panic: index out of range
		_ = s[n]         // panic: index out of range
		m["hello"] = 555 // panic: assign to entry in nil map
	}()

	func() {
		// Like map assignments, if a slice is assigned to another slice, they will
		// share all (underlying) elements.

		m0 := map[int]int{0: 7, 1: 8, 2: 9}
		m1 := m0
		m1[0] = 2
		fmt.Println(m0, m1) // map[0:2 1:8 2:9] map[0:2 1:8 2:9]

		s0 := []int{7, 8, 9}
		s1 := s0
		s1[0] = 2
		fmt.Println(s0, s1) // [2 8 9] [2 8 9]

		a0 := [...]int{7, 8, 9}
		a1 := a0
		a1[0] = 2
		fmt.Println(a0, a1) // [7 8 9] [2 8 9]
	}()

	func() {
		//  delete entries from maps:

		m := map[string]int{"Go": 2007}
		m["C"] = 1972    // append
		m["Java"] = 1995 // append
		fmt.Println(m)   // map[C:1972 Go:2007 Java:1995]

		m["Go"] = 2009    // modify
		delete(m, "Java") // delete
		fmt.Println(m)    // map[C:1972 Go:2009]
	}()

	func() {
		// example showing how to use the `append` function:

		s0 := []int{2, 3, 5}
		fmt.Println(s0, cap(s0)) // [2 3 5] 3
		s1 := append(s0, 7)      // append one element
		fmt.Println(s1, cap(s1)) // [2 3 5 7] 6
		s2 := append(s1, 11, 13) // append two elements
		fmt.Println(s2, cap(s2)) // [2 3 5 7 11 13] 6
		s3 := append(s0)         // <=> s3 := s0
		fmt.Println(s3, cap(s3)) // [2 3 5] 3
		s4 := append(s0, s0...)  // double s0 as s4
		fmt.Println(s4, cap(s4)) // [2 3 5 2 3 5] 6

		s0[0], s1[0] = 99, 789
		fmt.Println(s2[0], s3[0], s4[0]) // 789 99 2
	}()

	func() {
		var s = append([]string(nil), "array", "slice")
		fmt.Println(s)      // [array slice]
		fmt.Println(cap(s)) // 2
		s = append(s, "map")
		fmt.Println(s)      // [array slice map]
		fmt.Println(cap(s)) // 4
		s = append(s, "channel")
		fmt.Println(s)      // [array slice map channel]
		fmt.Println(cap(s)) // 4
	}()

	func() {
		//  how to use the built-in make function to create maps and slices:
		// Make new maps.
		fmt.Println(make(map[string]int)) // map[]
		m := make(map[string]int, 3)
		fmt.Println(m, len(m)) // map[] 0
		m["C"] = 1972
		m["Go"] = 2009

		fmt.Println(m, len(m)) // map[C:1972 Go:2009] 2

		// Make new slices.
		s := make([]int, 3, 5)
		fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 5
		s = make([]int, 2)
		fmt.Println(s, len(s), cap(s)) // [0 0] 2 2
	}()

	func() {
		// Allocate Containers With the Built-in `new` Function
		m := *new(map[string]int)   // <=> var m map[string]int
		fmt.Println(m == nil)       // true
		s := *new([]int)            // <=> var s []int
		fmt.Println(s == nil)       // true
		a := *new([5]bool)          // <=> var a [5]bool
		fmt.Println(a == [5]bool{}) // true
	}()

	func() {
		a := [5]int{2, 3, 5, 7}
		s := make([]bool, 2)
		pa2, ps1 := &a[2], &s[1]
		fmt.Println(*pa2, *ps1) // 5 false
		a[2], s[1] = 99, true
		fmt.Println(*pa2, *ps1) // 99 true
		ps0 := &[]string{"Go", "C"}[0]
		fmt.Println(*ps0) // Go

		m := map[int]bool{1: true}
		_ = m
		// The following lines fail to compile.
		/*
		   _ = &[3]int{2, 3, 5}[0]
		   _ = &map[int]bool{1: true}[1]
		   _ = &m[1]
		*/
	}()

	func() {

		type T struct{ age int }
		mt := map[string]T{}
		mt["John"] = T{age: 29} // modify it as a whole
		ma := map[int][5]int{}
		ma[1] = [5]int{1: 789} // modify it as a whole

		// The following two lines fail to compile,
		// for map elements can be modified partially.
		/*
		   ma[1][1] = 123      // error
		   mt["John"].age = 30 // error
		*/

		// Accesses are okay.

		fmt.Println(ma[1][1])       // 789
		fmt.Println(mt["John"].age) // 29
	}()

	func() {
		type T struct{ age int }
		mt := map[string]T{}
		mt["John"] = T{age: 29}
		ma := map[int][5]int{}
		ma[1] = [5]int{1: 789}

		t := mt["John"] // a temporary copy
		t.age = 30
		mt["John"] = t // overwrite it back

		a := ma[1] // a temporary copy
		a[1] = 123
		ma[1] = a // overwrite it back

		fmt.Println(ma[1][1], mt["John"].age) // 123 30
	}()

	func() {
		//  subslice syntax forms:

		a := [...]int{0, 1, 2, 3, 4, 5, 6}
		s0 := a[:]     // <=> s0 := a[0:7:7]
		s1 := s0[:]    // <=> s1 := s0
		s2 := s1[1:3]  // <=> s2 := a[1:3]
		s3 := s1[3:]   // <=> s3 := s1[3:7]
		s4 := s0[3:5]  // <=> s4 := s0[3:5:7]
		s5 := s4[:2:2] // <=> s5 := s0[3:5:5]
		s6 := append(s4, 77)

		s7 := append(s5, 88)
		s8 := append(s7, 66)
		s3[1] = 99
		fmt.Println(len(s2), cap(s2), s2) // 2 6 [1 2]
		fmt.Println(len(s3), cap(s3), s3) // 4 4 [3 99 77 6]
		fmt.Println(len(s4), cap(s4), s4) // 2 4 [3 99]
		fmt.Println(len(s5), cap(s5), s5) // 2 2 [3 99]
		fmt.Println(len(s6), cap(s6), s6) // 3 4 [3 99 77]
		fmt.Println(len(s7), cap(s7), s7) // 3 4 [3 4 88]
		fmt.Println(len(s8), cap(s8), s8) // 4 4 [3 4 88 66]
	}()

	func() {
		// Convert Slice to Array Pointer

		type S []int
		type A [2]int
		type P *A

		func() {
			var x []int
			var y = make([]int, 0)
			var x0 = (*[0]int)(x) // okay, x0 == nil
			var y0 = (*[0]int)(y) // okay, y0 != nil
			_, _ = x0, y0

			var z = make([]int, 3, 5)
			var _ = (*[3]int)(z) // okay
			var _ = (*[2]int)(z) // okay
			var _ = (*A)(z)      // okay
			var _ = P(z)         // okay

			var w = S(z)
			var _ = (*[3]int)(w) // okay
			var _ = (*[2]int)(w) // okay
			var _ = (*A)(w)      // okay
			var _ = P(w)         // okay

			return
			var _ = (*[4]int)(z) // will panic
		}()

	}()

	func() {
		// Convert Slice to Array

		var s = []int{0, 1, 2, 3}
		var a = [3]int(s[1:])
		s[2] = 9
		fmt.Println(s) // [0 1 9 3]
		fmt.Println(a) // [1 2 3]

		return
		_ = [3]int(s[:2]) // panic
	}()

	func() {
		type Ta []int
		type Tb []int
		dest := Ta{1, 2, 3}
		src := Tb{5, 6, 7, 8, 9}
		n := copy(dest, src)
		fmt.Println(n, dest) // 3 [5 6 7]
		n = copy(dest[1:], dest)
		fmt.Println(n, dest) // 2 [5 5 6]

		a := [4]int{} // an array
		n = copy(a[:], src)
		fmt.Println(n, a) // 4 [5 6 7 8]
		n = copy(a[:], a[2:])
		fmt.Println(n, a) // 2 [7 8 7 8]
	}()

	func() {

		m := map[string]int{"C": 1972, "C++": 1983, "Go": 2009}
		for lang, year := range m {
			fmt.Printf("%v: %v \n", lang, year)
		}

		a := [...]int{2, 3, 5, 7, 11}
		for i, prime := range a {
			fmt.Printf("%v: %v \n", i, prime)
		}

		s := []string{"go", "defer", "goto", "var"}
		for i, keyword := range s {
			fmt.Printf("%v: %v \n", i, keyword)
		}
	}()

	func() {
		aContainer := []string{"go", "defer", "goto", "var"}

		// Ignore the key iteration variable.
		for _, element := range aContainer {
			// ...
			println(element)
		}

		// Ignore the element iteration variable.
		for key, _ := range aContainer {
			element := aContainer[key]
			// ...
			println(element)
		}

		// The element iteration variable is omitted.
		// This form is equivalent to the last one.
		for key := range aContainer {
			element := aContainer[key]
			// ...
			println(element)
		}

		// Ignore both the key and element iteration variables.
		for _, _ = range aContainer {
			// This variant is not much useful.
		}

		// Both the key and element iteration variables are
		// omitted. This form is equivalent to the last one.
		for range aContainer {
			// This variant is not much useful.
		}
	}()

	fmt.Println("vim-go")
}
