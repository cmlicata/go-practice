package main

import "fmt"
import (
	"math/rand"
	"os"
)

func main() {
	var x int = rand.Int()

	// The syntax is different (from that in C): parentheses are not required and the body must always be brace-delimited.
	if x > 0 {
		fmt.Printf(" The value of x is greater than zero \n")
	} else {
		fmt.Printf(" The value of x is less than or equal to zero \n")
	}

	/*
		Mandatory braces encourage writing simple if statements on multiple lines. It is good
		style to do so anyway, especially when the body contains a control statement such as a
		return or break.
		Since if and switch accept an initialization statement, it’s common to see one used to
		set up a (local) variable.
	*/
	filename := "test.txt"
	if err := os.Chmod(filename, 0644); err != nil {
		fmt.Printf(string(err))
	}

	// You can use the logical operators (see table 1.1) as you would normally:
	if true && true {
		fmt.Println("true")
	}
	if !false {
		fmt.Println("true")
	}

	f, err := os.Open(filename)

	if err != nil {
		// If you use fmt.Printf(err) you will get a type error
		fmt.Println(err)
	}

	d, err := f.Stat()

	if err != nil {
		print(d)
	}

}

// Go has a goto statement — use it wisely. With goto you jump to a label which must be
// defined within the current function. For instance, a loop in disguise:
func myfunc() { // <-- Loop in disguise
	i := 0
Here: // <--- First word on a line ending with a colon is a label. The name of the label is case sensitive.
	println(i)
	i++
	goto Here // <-- Jump
}

// ----------------------------------------------------------
// -------------------------For Loops------------------------
// ----------------------------------------------------------

// Lower case characters in function names in Go mean they are private functions and an initial Capital Letter means that
func loops() {
	// The Go for loop has three forms, only one of which has semicolons.
	// for init; condition; post { } <--- like a C for loop
	// for condition {} <---- like a while loop
	// for {} <----- endless loop

	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1 // <--- short for sum = sum + i
	} // <--- i ceases to exist after the loop

	println(sum)
}

func cases(i int) {
	switch i {
	case 0:
		fallthrough
	case 1:
		println("The variable i is either 0 or 1")
	case 2:
		loops()
	default:
		println("This is the default case")
	}

}
