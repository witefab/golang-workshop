// Run this program using:
// go run task4_dial.go -dial=localhost:8000
// Lines typed in the 'dial' terminal should appear as JSON objects in the 'listen' terminal
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
)

var dialAddr = flag.String("dial", "localhost:8000", "host:port to dial")

type Message struct {
	Body string
}

func main() {
	flag.Parse()

	c, err := net.Dial("tcp", *dialAddr)
	if err != nil {
		log.Fatal(err)
	}

	// TODO :
	// read lines from standard input (os.Stdin)
	// create Message object
	// encode each line as a JSON object
	s := bufio.NewScanner(os.Stdin)
	e := json.NewEncoder(c)
	for s.Scan() {
		m := Message{Body: s.Text()}
		err := e.Encode(m)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
