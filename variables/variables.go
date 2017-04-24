package main //All Go files start with package <something>, and package main is required for a standalone executable.

import (
	"fmt"
) //Implements formatted I/O.  A package other than main is commonly called a library, a familiar concept in many programming languages.

func main() {

	println("Declaring variables with = operator")
	var a int
	var b bool
	a = 15
	b = true

	fmt.Printf("Value assigned to var a int is: %d\n", a)  // the %d verb is for integers
	fmt.Printf("Value assigned to var b bool is: %t\n", b) // the %t verb is for booleans

	// This form of declaration/assignment may only be used inside functions
	hello := "hello"
	float := 1.27

	/*
		Width is specified by an optional decimal number immediately preceding the verb. If absent, the width is whatever is necessary to represent the value. Precision is specified after the (optional) width by a period followed by a decimal number. If no period is present, a default precision is used. A period with no following number specifies a precision of zero. Examples:

		%f     default width, default precision
		%9f    width 9, default precision
		%.2f   default width, precision 2
		%9.2f  width 9, precision 2
		%9.f   width 9, precision 0
	*/
	fmt.Printf("Value assigned to float := %.2f\n", float) // the %f verb is for basic decimal formatting
	fmt.Printf("Value assigned to hello := %s\n", hello)   // the %s verb is for standard string printing

	var (
		x int
		y float64
	)
	x = 10
	y = 7.4

	fmt.Printf("Value assigned to x := %d\n", x)       // the %d verb is for basic integer formatting
	fmt.Printf("Value assigned to float := %.2f\n", y) // the %f verb is for basic decimal formatting

	// ----------------------------------------------------------
	// -----------------------CONSTANTS--------------------------
	// ----------------------------------------------------------

	// Constants in Go are just that — constant. They are created at compile time, and can only
	// be numbers, strings or booleans;
	const newConstant = 42

	println(newConstant)

	// You can use iota to enumerate values.
	// The first use of iota will yield 0, so constantEx1 is equal to 0, whenever iota is used again on a
	// new line its value is incremented with 1, so constantEx2 has a value of 1.

	println("Specifying iota for constants on each new line, thereby incrementing the value of iota. \n")
	const (
		constantEx1 = iota
		constantEx2 = iota
	)

	fmt.Printf("Value assigned to constantEx1 := %d\n", constantEx1)
	fmt.Printf("Value assigned to constantEx2 := %d\n", constantEx2)

	println("Or you can let Go repeat the use of iota by only specifying the assignment of iota to the " +
		"first variable and then list the rest of the variables on new lines. \n")
	const (
		constantEx3 = iota
		// Implicitly constantEx4 = iota
		constantEx4
	)

	fmt.Printf("Value assigned to constantEx3 := %d\n", constantEx3)
	fmt.Printf("Value assigned to constantEx4 := %d\n", constantEx4)

	println("You may also explicity type a constant, if you need that \n")

	const (
		constantEx5 = 0 // is an int now
		constantEx6 = "0"
	)

	fmt.Printf("Value assigned to constantEx5 := %d\n", constantEx5)
	fmt.Printf("Value assigned to constantEx6 := %s\n", constantEx6)

	// Strings in Go are a sequence of UTF-8 characters enclosed in double quotes (”). If you use
	// the single quote (’) you mean one character (encoded in UTF-8) — which is not a string
	// in Go.

	// ----------------------------------------------------------
	// -------------------------Strings--------------------------
	// ----------------------------------------------------------

	var s string = "hello"
	// s[0] = 'c' <----- Change first char. to 'c', this is an error

	// To do this in Go you will need the following

	c := []rune(s)          // <---- converts to an array of runes
	c[0] = 'c'              // <---- change the first element of this array
	s2 := string(c)         // <---- create a new string s2 with the alteration
	fmt.Printf("%s \n", s2) // <---- print the string with fmt.Printf

	 /* Multi-line strings (See https://golang.org/doc/effective_go.html#semicolons)
		Due to the insertion of semicolons (see [8] section “Semicolons”), you need to
		be careful with using multi line strings. If you write:
		s := "Starting part"
		    + "Ending part"

	 	This is transformed into:
		s := "Starting part"; <-- notice the semicolons
		    + "Ending part";

	 	Which is not valid syntax, you need to write:
		s := "Starting part" +  <-- notice that the + operator is at the end of the first string and not at the beginning of the next one.
			"Ending part"

	 	Then Go will not insert the semicolons in the wrong places. Another way would be to
		use raw string literals by using backquotes (`):
		s := `Starting part
			Ending part`

	 Be aware that in this last example s now also contains the newline. Unlike interpreted
	 string literals the value of a raw string literal is composed of the uninterpreted characters
	 between the quotes */

	// ----------------------------------------------------------------------------------
	// -------------------------String Manipulation with Runes --------------------------
	// ----------------------------------------------------------------------------------

	/* Rune is an alias for int32. It is an UTF-8 encoded code point. When is this type useful?
	For instance, when iterating over characters in a string. You can loop over each byte
	(which is only equivalent to a character when strings are encoded in 8-bit ASCII, which
	they are not in Go!). So to get the actual characters you should use the rune type. */

	// A code unit is the number of bits an encoding uses. So UTF-8 would use 8 and UTF-16 would use 16 units.
	// A code point is a character and this is represented by one or more code units depending on the encoding.

}
