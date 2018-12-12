package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// TODO: Change the main function so that it doesnt count in locksteps
	// you can call the fanIn function to let whosoever is ready to talk
	// and prints out the names/strings
	joe := boring("Joe!")
	ann := boring("Ann!")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-joe)
		fmt.Printf("You say: %q\n", <-ann)
	}
	fmt.Println("The End")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	// TODO: magic
}
