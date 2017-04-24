package main

import "fmt"

func main() {
	// Create a map with a key of type string and a value of type int.
	dict1 := make(map[string]int)
	fmt.Printf("This is what a declared map looks like: %v \n", dict1)

	// Create a map with a key and value of type string.
	// Initialize the map with 2 key/value pairs.
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Printf("This is how an initialized map looks: %v \n", dict2)

	/*
	   The map key can be a value from any built-in or struct type as long as the value can be used in an expression with the == operator.
	   Slices, functions, and struct types that contain slices can’t be used as map keys. This will produce a compiler error.

	   However, there is nothing stopping you from using a slice as a map value.
	*/
	// Create a map using a slice of strings as the key.
	// dict := map[[]string]int{}
	//
	// Compiler Exception:
	// invalid map key type []string

	// Create a map using a slice of strings as the value.
	//dict3 := map[int][]string{}

	// Create an empty map to store colors and their color codes.
	colors := map[string]string{}

	// Add the Red color code to the map.
	colors["Red"] = "#da1337"

	// You can create a nil map by declaring a map without any initialization.
	// A nil map can’t be used to store key/value pairs. Trying will produce a runtime error.
	// Create a nil map by just declaring the map.
	// var colors map[string]string
	//
	// Add the Red color code to the map.
	// colors["Red"] = "#da1337"
	//
	// Runtime Error:
	// panic: runtime error: assignment to entry in nil map

	// Retrieve the value for the key "Blue".
	value, _ := colors["Blue"]

	// Did this key exist?
	fmt.Printf("Does the key Blue exist: %v \n", value)

	// Retrieve the value for the key "Blue".
	_, exists := colors["Red"]

	// Did this key exist?
	if exists {
		fmt.Printf("Does the key Red exist: %t \n\n\n", exists)
	}

	// Create a map of colors and color hex codes.
	colors = map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
	println("\n\n\n")

	//println("Deleting AliceBlue from the map.")
	//delete(colors, "AliceBlue")

	// Call the function to remove the specified key.
	println("Using the function removeColor to deleting Coral from the map.")
	removeColor(colors, "Coral")
	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}

}

func removeColor(colorMap map[string]string, key string) {
	delete(colorMap, key)
}
