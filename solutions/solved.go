package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Product struct {
	Name     string
	Price    int
	Brand    string
	Volume   int
	Quantity int
}

func main() {

	jsonString, e := ioutil.ReadFile("inventory.json")
	if e != nil {
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(jsonString))

	var products []Product
	json.Unmarshal([]byte(jsonString), &products)
	fmt.Printf("Products : %+v", products)

}
