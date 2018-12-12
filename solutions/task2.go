package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {

	// TODO :
	// read lines from standard input (os.Stdin) using the bufio package
	// create Message object
	// encode each line as a JSON object, written to standard output (os.Stdout) using the encoding/json package
	s := bufio.NewScanner(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for s.Scan() {
		var m Message
		m = Message{Body: s.Text()}
		err := enc.Encode(m)
		if err != nil {
			log.Fatal(err)
		}
	}
}
