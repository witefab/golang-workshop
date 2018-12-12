package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	boring("boring!")
	// go 	boring("boring!")
	// fmt.Println("Im listening")
	// time.Sleep(2 * time.Second)
	// fmt.Println("This is boring")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
