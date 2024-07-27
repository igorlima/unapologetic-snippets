package main

import "fmt"

func main() {

	func() {
		type Book struct {
			title, author string
			pages         int
		}

		book := Book{"Go 101", "Tapir", 256}
		fmt.Println(book) // {Go 101 Tapir 256}

		// Create a book value with another form.
		// All of the three fields are specified.
		book = Book{author: "Tapir", pages: 256, title: "Go 101"}

		// None of the fields are specified. The title and
		// author fields are both "", pages field is 0.
		book = Book{}

		// Only specify the author field. The title field
		// is "" and the pages field is 0.
		book = Book{author: "Tapir"}

		// Initialize a struct value by using selectors.
		var book2 Book // <=> book2 := Book{}
		book2.author = "Tapir Liu"
		book2.pages = 300
		fmt.Println(book2.pages) // 300

		func() {
			var _ = Book{
				author: "Tapir",
				pages:  256,
				title:  "Go 101", // here, the "," must be present
			}
			// The last "," in the following line is optional.
			var _ = Book{author: "Tapir", pages: 256, title: "Go 101"}
		}()

		func() {
			/*
				When a struct value is assigned to another struct value,
				the effect is the same as assigning each field one by one.
			*/

			book1 := Book{pages: 300}
			book2 := Book{"Go 101", "Tapir", 256}

			book2 = book1
			// The above line is equivalent to the
			// following lines.
			book2.title = book1.title
			book2.author = book1.author
			book2.pages = book1.pages
		}()
	}()

	func() {
		// Struct Field Addressability
		func() {
			type Book struct {
				Pages int
			}
			var book = Book{} // book is addressable
			p := &book.Pages
			*p = 123
			fmt.Println(book) // {123}

			// The following two lines fail to compile, for
			// Book{} is unaddressable, so is Book{}.Pages.
			/*
			   Book{}.Pages = 123
			   p = &(Book{}.Pages) // <=> p = &Book{}.Pages
			*/
		}()
		func() {
			type Book struct {
				Pages int
			}
			// Book{100} is unaddressable but can
			// be taken address.
			p := &Book{100} // <=> tmp := Book{100}; p := &tmp
			p.Pages = 200
		}()

		func() {
			//  Dereferences of Receivers
			type Book struct {
				pages int
			}
			book1 := &Book{100} // book1 is a struct pointer
			book2 := new(Book)  // book2 is another struct pointer
			// Use struct pointers as structs.
			book2.pages = book1.pages
			// This last line is equivalent to the above line.
			// In other words, if the receiver is a pointer,
			// it will be implicitly dereferenced.
			(*book2).pages = (*book1).pages
		}()

		func() {
			/*
				Struct Value Conversions

				Values of two struct types S1 and S2 can be converted
				to each other's types, if S1 and S2 share
				the identical underlying type (by ignoring field tags).
			*/

			type S0 struct {
				y int "foo"
				x bool
			}
			// S1 is an alias of an unnamed type.
			type S1 = struct {
				x int "foo"
				y bool
			}
			// S2 is also an alias of an unnamed type.
			type S2 = struct {
				x int "bar"
				y bool
			}

			// If field tags are ignored, the underlying
			// types of S3(S4) and S1 are same. If field
			// tags are considered, the underlying types
			// of S3(S4) and S1 are different.
			type S3 S2 // S3 is a defined (so named) type
			type S4 S3 // S4 is a defined (so named) type

			var v0, v1, v2, v3, v4 = S0{}, S1{}, S2{}, S3{}, S4{}
			func() {
				v1 = S1(v2)
				v2 = S2(v1)
				v1 = S1(v3)
				v3 = S3(v1)
				v1 = S1(v4)
				v4 = S4(v1)
				v2 = v3
				v3 = v2 // the conversions can be implicit
				v2 = v4
				v4 = v2 // the conversions can be implicit

				v3 = S3(v4)
				v4 = S4(v3)
				/*
					In fact, two struct values can be assigned (or compared) to each
					other only if one of them can be implicitly converted to the type of
					the other.
				*/
			}()
		}()

		func() {
			// Anonymous Struct Types
			// can be used in field declarations

			var aBook = struct {
				// The type of the author field is
				// an anonymous struct type.
				author struct {
					firstName, lastName string
					gender              bool
				}
				title string
				pages int
			}{
				author: struct { // an anonymous struct type
					firstName, lastName string
					gender              bool
				}{
					firstName: "Mark",
					lastName:  "Twain",
				},
				title: "The Million Pound Note",
				pages: 96,
			}
		}()
	}()

	fmt.Println("vim-go")
}
