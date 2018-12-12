# Golang

This repository contains files and instructions needed for the golang workshop of the module ESD at Fontys University of Applied Sciences in Venlo. The workshop is given by the authors Mathias Nawacki and Fabian Witeczek and will introduce you to golang and teach you it's basics.

## Prerequisites

Install golang on your machine. Download it [here](https://golang.org/dl/) or install the Linux repository: 

`sudo apt install golang-go`

[Instructions on how to install the binaries](https://golang.org/doc/install). To verify that the installation was successful, create a `hello.go` file with the following code:

```
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```


Save the file and run it with `go run hello.go`. Alternately you can use `go build` in `hello.go`'s directory to receive an executable named `hello`.

We recommend you to use Atom or Visual Studio Code (these are the IDEs we are using) since they come with easy to install golang extensions for code completion etc.

## Task 1: Reading a JSON file

The first task is to read the `inventory.json` file and store it as a string. After defining the `Product` struct the content of the JSON string should be migrated to be stored in an array of struct objects. These steps are listed as `TODO`s in the file as well.

## Task 2: Transforming input to JSON

Now we want to create a simple program which takes user input and transforms it into a JSON string. To accomplish this you may use the `Message` struct in order to use the `encoding/json`'s `Encode` method. For the sake of this task it is sufficient to generate a JSON string with a body string as the only attribute.

## Task 3: Lockstep counting

In the current program there is only one goroutine in use - so Joe and Ann count in lockstep. This means that they have to wait for the other party to finish printing before they can print themselves.

To resolve this lock implement the `fanIn` function and also change the main function, so that they can print their output whenever they want.


## Task 4: Listening and dialing through TCP

For this task there should be 2 components communicating with each other: The listener (`task4_listen.go`) and the dialer (`task4_dial.go`). The listener should listen on the port which is passed with the `-listen` argument. For port 8000 this would look like that: `go run task4_listen.go -listen=localhost:8000`. For every incoming connection, it should launch a goroutine that reads and decodes JSON-encoded messages from the connection and prints them to the standard output.
But before that can happen the dialer has to be able to translate user input to a JSON string (which is similar to the problem in task 2).
