# Golang

This workshop will show you how to use golang and what it's benefits are.

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


## Task 1: Reading a JSON file

The first task is to read the `inventory.json` file and store it as a string. After defining the `Product` struct the content of the JSON string should be migrated to be stored in an array of struct objects. These steps are listed as `TODO`s in the file as well.
