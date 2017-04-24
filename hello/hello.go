package main //All Go files start with package <something>, and package main is required for a standalone executable.

import "fmt" //Implements formatted I/O.  A package other than main is commonly called a library, a familiar concept in many programming languages (see ).

/* Print something

When your Go program is executed, the first function called will be main.main(), which mimics the behavior from C. Here we declare that function
*/
func main() {
	// Finally we call a function from the package fmt to print a string to the screen. The string is enclosed with " and may contain non-ASCII characters
	fmt.Printf("Hello, world.")
}
