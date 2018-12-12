package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	joe := boring3("Alice!")
// 	ann := boring3("Bob!")

// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("You say: %q\n", <-joe)
// 		fmt.Printf("You say: %q\n", <-ann) /// bob still has to wait for joe
// 	}
// 	fmt.Println("The End")
// }

// func boring3(msg string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 0; ; i++ {
// 			c <- fmt.Sprintf("%s %d", msg, i)
// 			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 		}
// 	}()
// 	print("combo killer")
// 	return c
// }

func main() {
	c := fanIn(boring("Ann"), boring("Bob"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("you both are boring")
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
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
