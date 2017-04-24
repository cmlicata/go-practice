package main

import "fmt"

func main() {

	// Initializing only specific elements in the array using a special array literal
	array := [5]int{1: 2, 2: 2}
	fmt.Printf("This is an array initalized with the special index-specific array literal: %v \n", array)

	newArray := [2]int{1, 2}
	fmt.Printf("This is an array initalized with the normal array literal: %v \n", newArray)

	reallyNewArray := [...]int{1, 2, 3, 4, 5, 6, 7, 7, 7, 8, 9}
	fmt.Printf("This is an array initalized with the undefined size array literal: %v \n", reallyNewArray)

	// Declare an integer array of five elements.
	// Initialize each element with a specific value.
	array = [5]int{10, 20, 30, 40, 50}
	fmt.Printf("This is a fixed-length array initialized so that each element in the array has a specific value.: %v \n", array)

	// Declare an integer array.
	// Initialize each element with a specific value.
	// Capacity is determined based on the number of values initialized.
	array = [...]int{10, 20, 30, 40, 50}
	fmt.Printf("This is a variable-lenth array where Go calculates the length of the array for us upon initialization: %v \n", array)

	// Declare an integer array of five elements.
	// Initialize index 1 and 2 with specific values.
	// The rest of the elements contain their zero value.
	array = [5]int{1: 10, 2: 20}
	fmt.Printf("This is another array initalized with the special index-specific array literal: %v \n", array)

	// Declare an integer pointer array of five elements.
	// Initialize index 0 and 1 of the array with integer pointers.
	pointerArray1 := [5]*int{0: new(int), 1: new(int)}

	fmt.Println(pointerArray1)

	// Assign values to index 0 and 1.
	*pointerArray1[0] = 10
	*pointerArray1[1] = 20

	fmt.Println(pointerArray1)

	// Declare a string array of five elements.
	var array1 [5]string

	fmt.Printf("This is the blank [5]string array1: %v  \n", array1)

	// Declare a second string array of five elements.
	// Initialize the array with colors.
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	fmt.Printf("This is the initialized [5]string array2: %v \n", array2)

	// Copy the values from array2 into array1.
	array1 = array2

	fmt.Printf("This is the [5]string array1 after copying array2: %v \n", array1)

	// ----------------------------------------------------------
	//-------------Assigning arrays of different types-----------
	// ----------------------------------------------------------

	/*
		The type of an array variable includes both the length and the type of data that can be stored in each element.
		Only arrays of the same type can be assigned.

			// Declare a string array of four elements.
			var array1 [4]string

			// Declare a second string array of five elements.
			// Initialize the array with colors.
			array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

			// Copy the values from array2 into array1.
			array1 = array2

			Compiler Error:
			cannot use array2 (type [5]string) as type [4]string in assignment
	*/

	// Copying an array of pointers copies the pointer values and not the values that the pointers are pointing to.
	// Declare a string pointer array of three elements.
	var array3 [3]*string
	fmt.Printf("This is the [3]*string array3: %v \n", array3)

	// Declare a second string pointer array of three elements.
	// Initialize the array with string pointers.
	array4 := [3]*string{new(string), new(string), new(string)}

	// Add colors to each element
	*array4[0] = "Red"
	*array4[1] = "Blue"
	*array4[2] = "Green"
	fmt.Printf("This is the [3]*string array4: %v \n", array4)

	// Copy the values from array2 into array1.
	array3 = array4

	fmt.Printf("This is the [3]*string array3 after copying array4: %v \n", array3)
	fmt.Println("After the copy, you have two arrays pointing to the same strings")

	// Copy index 1 of array1 into a new array of the same type.

	// Declare two different two dimensional integer arrays.
	var array5 [2][2]int
	var array6 [2][2]int

	// Add integer values to each individual element.
	array6[0][0] = 10
	array6[0][1] = 20
	array6[1][0] = 30
	array6[1][1] = 40

	// Copy the values from array2 into array1.
	array5 = array6
	var array7 [2]int = array5[1]

	// Copy the integer found in index 1 of the outer array
	// and index 0 of the interior array into a new variable of
	// type integer.
	var value int = array5[1][0]

	fmt.Printf("This is the array7 after copying everything from array5[1] in it: %v \n", array7)
	fmt.Printf("This is the value int variable after copying array5[1][0] to it: %v \n", value)
}
