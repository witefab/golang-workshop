package main

import (
	"fmt"
)

// TODO2: create a "Product" struct and adapt the fields to the JSON file's structure

func main() {

	var jsonString string
	// TODO1: read the inventory.json using the ioutil package and store the JSON string in the above variable
	fmt.Printf("%s\n", string(jsonString))

	var products []Product
	// TODO3: fill this array by Unmarshalling the JSON String

}
