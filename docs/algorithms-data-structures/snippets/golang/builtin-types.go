package main

import "fmt"

func main() {
	/*
		ZERO VALUES
	*/
	var status bool
	var MyString string
	var Id uint64
	var real float32
	fmt.Println("Declaring Variables in Go")
	fmt.Println("-------------------------------")
	fmt.Printf("Type of status is %T and has value %v\n", status, status)
	fmt.Printf("Type of MyString is %T and has value %v\n", MyString, MyString)
	fmt.Printf("Type of Id is %T and has value %v\n", Id, Id)
	fmt.Printf("Type of real is %T and has value %v\n", real, real)
	fmt.Println("...............................\n")

	/*
	  type declarations
	*/
	var boolean bool = true
	var Text string = "Hello, World!"
	var U8 uint8 = 255
	var char rune = 'A'
	fmt.Println("Declaring Variables in Go")
	fmt.Println("-------------------------------")
	fmt.Printf("Type of boolean is %T and has value %v\n", boolean, boolean)
	fmt.Printf("Type of Text is %T and has value %v\n", Text, Text)
	fmt.Printf("Type of U8 is %T and has value %v\n", U8, U8)
	fmt.Printf("Type of char is %T and has value %v\n", char, char)
	fmt.Println("...............................\n")

	var booleanB = true
	var TextB = "Hello, World!"
	var U8B = 255
	fmt.Println("Declaring Variables in Go")
	fmt.Println("-------------------------------")
	fmt.Printf("Type of booleanB is %T and has value %v\n", booleanB, booleanB)
	fmt.Printf("Type of TextB is %T and has value %v\n", TextB, TextB)
	fmt.Printf("Type of U8B is %T and has value %v\n", U8B, U8B)
	fmt.Println("...............................\n")

	/*
		In the following declarations, the types of the multiple variables
		declared in the same declaration line can be different.
	*/
	// The types of the lang and dynamic variables
	// will be deduced as built-in types "string"
	// and "bool" by compilers, respectively.
	var lang, dynamic = "Go", false
	// The types of the compiled and announceYear
	// variables will be deduced as built-in
	// types "bool" and "int", respectively.
	var compiled, announceYear = true, 2009
	_, _, _, _ = lang, dynamic, compiled, announceYear // make lang, dynamic, compiled, and announceYear used.

	/*
		Multiple variables can be grouped into
		one standard form declaration by using ().
	*/
	var (
		announceAt, releaseAt int = 2009, 2012
		createdBy, website    string
	)
	_, _, _, _ = announceAt, releaseAt, createdBy, website // make announceAt, releaseAt, createdBy, and website used.

	/*
		SHORT VARIABLE DECLARATION FORMS
	*/
	/*
		Deductions in Go
		Go compilers will deduce the types for these values by context.
		Type deduction is also often called type inference.
	*/
	// The types of variables booleanC and TextC will be
	// deduced as the built-in type "bool" and "string",
	booleanC := true
	TextC := "Hello, World!"
	fmt.Println("Deduced Type Variables in Go")
	fmt.Println("-------------------------------")
	fmt.Printf("Type of booleanC is %T and has value %v\n", booleanC, booleanC)
	fmt.Printf("Type of TextC is %T and has value %v\n", TextC, TextC)
	fmt.Println("...............................\n")

	/*
	   There are two forms of string value literals, interpreted string literal
	   (double quotes form) and raw string literal (back quotes form). For example,
	   the following two string literals are equivalent:
	*/
	fmt.Println("Declaring String Literals in Go")
	fmt.Println("-------------------------------")
	var interpretedString = "Hello\nworld!\n\"你好世界\""
	var rawString = `Hello
	world!
	"你好世界"`
	fmt.Println("The interpreted string is: ", interpretedString)
	fmt.Println("The raw string is: ", rawString)
	fmt.Println("...............................\n")

	fmt.Println("Constants in Go")
	fmt.Println("-------------------------------")
	// Declare two individual constants. Yes,
	// non-ASCII letters can be used in identifiers.
	const π = 3.1416
	const Pi = π // <=> const Pi = 3.1416
	// Declare multiple constants in a group.
	const (
		Yes        = true
		No         = !Yes
		MaxDegrees = 360
		Unit       = "radian"
	)
	// typed named constants
	const X float32 = 3.14
	const (
		A, B int64   = -3, 5
		Y    float32 = 2.718
	)
	fmt.Println("The value of π is: ", π)
	fmt.Println("...............................\n")

}
