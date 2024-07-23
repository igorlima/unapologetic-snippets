package main

import "fmt"

func main() {
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

	booleanC := true
	TextC := "Hello, World!"
	fmt.Println("Declaring Variables in Go")
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
	var interpretedString = "Hello\nworld!\n\"你好世界\""
	var rawString = `Hello
	world!
	"你好世界"`
	fmt.Println("The interpreted string is: ", interpretedString)
	fmt.Println("The raw string is: ", rawString)
}
