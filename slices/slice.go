package main

import (
	"fmt"
)

func main() {
	println()
	println("------------")
	println("------------")
	println("Slices in Go")
	println("------------")
	println("------------\n")
	// One way to create a slice is to use the built-in function make.
	// When you use make, one option you have is to specify the length of the slice.

	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	slice := make([]string, 5)
	fmt.Printf("This is a string slice created using the built-in function make with only length specified: %v \n", slice)

	// When you only specify the length (see above) the capacity of the slice is the same.
	// But, you can also specify the length and capacity separately (see below).

	// Create a slice of integers.
	// Contains a length of 3 and has a capacity of 5 elements.
	slice1 := make([]int, 3, 5)
	fmt.Printf("This is a slice created using the built-in function make with both length and capacity specified: %v \n", slice1)

	/*
		Create a slice of integers.
		Make the length larger than the capacity.
		slice := make([]int, 5, 3)

		Compiler Error:
		len larger than cap in make([]int)
	*/

	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	slice2 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Printf("Slice created using a slice literal: %v \n", slice2)

	// When using a slice literal, you can set the initial length and capacity.
	// All you need to do is initialize the index that represents the length and capacity you need

	// The following syntax will create a slice with a length and capacity of 100 elements.
	// Create a slice of strings.
	// Initialize the 100th element with an empty string.
	slice3 := []string{99: ""}
	fmt.Printf("Slice created using a slice literal that initialized the index representing the length and capacity of said slice: %v \n\n\n", slice3)

	// Sometimes in your programs you may need to declare a nil slice. A nil slice is created by declaring a slice without any initialization.

	// Create a nil slice of integers.
	// A nil slice is the most common way you create slices in Go.
	// They’re useful when you want to represent a slice that doesn’t exist, such as when an exception occurs in a function that returns a slice.
	var slice4 []int
	fmt.Printf("Simply declare an empty slice of integers.: %v \n", slice4)

	// Use make to create an empty slice of integers.
	slice5 := make([]int, 0)
	fmt.Printf("Use make to create an empty slice of integers.: %v \n", slice5)

	// Use a slice literal to create an empty slice of integers.
	slice6 := []int{}
	fmt.Printf("Use a slice literal to create an empty slice of integers: %v \n\n", slice6)

	// An empty slice contains a zero-element underlying array that allocates no storage.

	println("-----------------")
	println("-----------------")
	println("Working w/ slices")
	println("-----------------")
	println("-----------------\n")

	// Create a slice of integers.
	// Contains a length and capacity of 5 elements.
	slice6 = []int{10, 20, 30, 40, 50}

	// Change the value of index 1.
	slice6[1] = 25
	fmt.Printf("Declaring a slice using a slice literal and changing one of the values using its index %v \n\n\n", slice6)

	// Taking a slice of a slice
	// Create a slice of integers.
	// Contains a length and capacity of 5 elements.
	slice7 := []int{10, 20, 30, 40, 50}
	fmt.Printf("slice7 created using a slice literal: %v \n", slice7)

	/*
		The original slice views the underlying array as having a capacity of five elements, but the view of newSlice is different.
		For newSlice, the underlying array has a capacity of four elements.
		newSlice can’t access the elements of the underlying array that are prior to its pointer.
		As far as newSlice is concerned, those elements don’t even exist.
	*/

	// Create a new slice.
	// Contains a length of 2 and capacity of 4 elements.
	newSlice := slice7[1:3]
	fmt.Printf("newSlice created using a slice [1:3] of slice7: %v \n", newSlice)

	/*
		How len
		For slice[i:j] with an underlying array of capacity k

		Length:   j - i
		Capacity: k - i
	*/

	println("Changing the value of the element at the first index of newSlice.  Let's see if that effects slice7.")
	newSlice[1] = 24

	fmt.Printf("slice7 after the reassignment of the element at the first index of newSlice: %v \n", slice7)

	slice7[1] = 1000
	fmt.Printf("newSlice List after the reassignment of the element at the first index of slice7: %v \n", newSlice)
	fmt.Printf("slice7 after the reassignment of the element at the first index of slice7: %v \n\n", slice7)

	println("--------------")
	println("--------------")
	println("Growing Slices")
	println("--------------")
	println("--------------\n")

	// Create a slice of integers.
	// Contains a length and capacity of 5 elements.
	slice8 := []int{10, 20, 30, 40, 50}
	fmt.Printf("slice8 created using a slice literal: %v \n", slice8)

	// Create a new slice.
	// Contains a length of 2 and capacity of 4 elements.
	newSlice2 := slice8[1:3]
	fmt.Printf("newSlice2 created using a slice [1:3] of slice8: %v \n", newSlice2)

	// Allocate a new element from capacity.
	// Assign the value of 60 to the new element.
	newSlice2 = append(newSlice2, 60)
	fmt.Printf("newSlice2 after appending 60: %v \n", newSlice2)
	fmt.Printf("slice8 after appending 60 to newSlice2: %v \n", slice8)

	// Because there was available capacity in the underlying array for newSlice2,
	// the append operation incorporated the available element into the slice’s
	// length and assigned the value. Since the original slice is sharing the underlying array,
	// slice8 also sees the changes in index 3.
	newSlice2 = append(newSlice2, 100)
	fmt.Printf("newSlice2 after appending 100: %v \n", newSlice2)
	fmt.Printf("slice8 after appending 100 to newSlice2: %v \n", slice8)

	// Because there was available capacity in the underlying array for newSlice2,
	// the append operation incorporated the available element into the slice’s
	// length and assigned the value. Since the original slice is sharing the underlying array,
	// slice8 also sees the changes in index 4.
	newSlice2 = append(newSlice2, 77)

	// After this append operation, newSlice2 is given its own underlying array,
	// and the capacity of the array is doubled from its original size
	fmt.Printf("newSlice2 after appending 77: %v \n", newSlice2)
	fmt.Printf("slice8 after appending 77 to newSlice2: %v \n\n", slice8)

	fmt.Printf("newSlice2 capacity after appending 77: %d \n", cap(newSlice2))
	fmt.Printf("slice8 capacity after appending 77: %d \n", cap(slice8))
	fmt.Println("The reason for this is that we used a slice literal to declare and initialize slice8 with a length and capacity 5.")
	fmt.Println("The reason why newSlice2 doubled its capacity and size when we appended 77 is because we never explicity declared it's capacity and size." +
		"After this append operation, newSlice2 is given its own underlying array, and the capacity of the underlying array is doubled from its original size")

	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Printf("source created using a slice literal: %v \n\n", source)

	// Slice the third element and restrict the capacity.
	// Contains a length of 1 element and capacity of 2 elements.
	slice9 := source[2:3:4]
	fmt.Printf("slice9 created using a slice [2:3:4] of source: %v \n", slice9)
	fmt.Printf("slice9's capacity: %d \n", cap(slice9))

	slice9 = append(slice9, "Mango")
	slice9 = append(slice9, "Coconut")
	fmt.Printf("slice9 after appending Mango and Coconut: %v \n", slice9)
	fmt.Printf("slice9's capacity after appending Coconut and Mango: %d \n", cap(slice9))
	// For slice[i:j:k]  or  [2:3:4]
	// Length:   j - i  or  3 - 2 = 1
	// Capacity: k - i  or  4 - 2 = 2

	// This will cause a **runtime** error
	// slice10 := source[2:3:6]

	/*
		It’s easy to forget which slices are sharing the same underlying array.
		When this happens,making changes to a slice can result in random and odd-looking bugs.
		By having the option to set the capacity of a new slice to be the same as the length,
		you can force the first append operation to detach the new slice from the underlying array.
		Detaching the new slice from its original source array makes it safe to change.
	*/

	// Slice the third element and restrict the capacity.
	// Contains a length and capacity of 1 element.
	slice10 := source[2:3:3]
	fmt.Printf("slice10 created using a slice [2:3:3] of source: %v \n", slice10)

	// Append a new string to the slice.
	slice10 = append(slice10, "Kiwi")

	// Without this third index - 3 , appending Kiwi to our slice would’ve changed the value of Banana in index 3
	// of the underlying array, because all of the remaining capacity would still belong to the slice.

	/*
		The built-in function append is also a variadic function. This means you can pass multiple values to be
		 appended in a single slice call. If you use the ... operator, you can append all the elements of one
		 slice into another.
	*/

	s1 := []int{1, 2}
	s2 := []int{3, 4}

	fmt.Printf("%v\n", append(s1, s2...))

	println("---------------------")
	println("---------------------")
	println("Iterating over slices")
	println("---------------------")
	println("---------------------\n")

	/*
		range, when iterating over slices returns not only the index of the element, but also a **copy** of the element
		at that index.

		The following code was copied from https://www.goinggo.net/2013/09/iterating-over-slices-in-go.html.
		Thanks Will for always explaining things in the best and most understandable way.

	*/

	jackie := Dog{
		Name: "Jackie",
		Age:  19,
	}

	fmt.Printf("Jackie Addr: %p\n", &jackie)

	sammy := Dog{
		Name: "Sammy",
		Age:  10,
	}

	fmt.Printf("Sammy Addr: %p\n", &sammy)

	dogs := []Dog{jackie, sammy}

	fmt.Println("")

	for _, dog := range dogs {
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n", &dog)

		fmt.Println("")
	}

	allDogs := []*Dog{}

	for _, dog := range dogs {
		allDogs = append(allDogs, &dog)
	}

	for _, dog := range allDogs {
		fmt.Printf("Name: %s Age: %d\n\n", dog.Name, dog.Age)
	}

	/*
		Just remember that when you are iterating over a slice, you are getting a copy of each element of the slice.
		If that happens to be an object, you are getting a copy of that object.
		Don’t ever use the address of the local variable in the range loop.
		That is a local variable that contains a copy of the slice element and only has local context.
		Don’t make the same mistake I made.
	*/

	// The working code that produces the intended results is here.

	danny := &Dog{
		Name: "Danny",
		Age:  19,
	}

	fmt.Printf("Danny Addr: %p\n", danny)

	tammy := &Dog{
		Name: "Tammy",
		Age:  10,
	}

	fmt.Printf("Tammy Addr: %p\n\n", tammy)

	theDogs := []*Dog{danny, tammy}

	for _, dog := range theDogs {
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n\n", dog)
	}

	// If you don't need the index value while you are iterating over the slice,
	// you can use the underscore character to discard the value.

	// Create a slice of integers.// Contains a length and capacity of 4 elements.
	numberSlice := []int{10, 20, 30, 40}
	println("Output when using range keyword:")

	// Iterate over each element and display each value.
	for _, value := range numberSlice {
		fmt.Printf("Value: %d\n", value)
	}

	println("Output when using a traditional for loop starting at index 2:")
	// Iterate over each element starting at element 3.
	for index := 2; index < len(numberSlice); index++ {
		fmt.Printf("Index: %d  Value: %d\n", index, numberSlice[index])
	}

	slice11 := [][]int{{10}, {100, 200}}

	slice11[0] = append(slice11[0], 155)

	slice11[0] = append(slice11[0], 654321)
	for _, value := range slice11 {
		fmt.Printf("Value: %d\n", value)
	}
	println(cap(slice11[0]))
	println(cap(slice11))

	slice11 = append(slice11, []int{10, 22})
	for _, value := range slice11 {
		fmt.Printf("Value: %d\n", value)
	}
	println(cap(slice11))

}

type Dog struct {
	Name string
	Age  int
}
