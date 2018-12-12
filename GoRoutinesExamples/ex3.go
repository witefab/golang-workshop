package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Our boring function returns a channel that lets us communicate with
// the boring service it provides.
// We can have more instances of the services.

func main() {
	joe := boring("joe!")
	ann := boring("ann!")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-joe)
		fmt.Printf("You say: %q\n", <-ann)
	}
	fmt.Println("Im leaving")
}

func boring(msg string) <-chan string { //Returns receive-only channel of strings
	c := make(chan string)
	go func() { //We launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c //Return the channel to the caller
}
