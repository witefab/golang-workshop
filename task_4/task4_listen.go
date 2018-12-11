// Run this program using:
// go run task4_listen.go -listen=localhost:8000
// Lines typed in the 'dial' terminal should appear as JSON objects in the 'listen' terminal
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
)

var listenAddr = flag.String("listen", "localhost:8000", "host:port to listen on")

type Message struct {
	Body string
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: accept incoming connections and launch a goroutine to handle each one
	c, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	serve(c)

}

func serve(c net.Conn) {
	defer c.Close()
	d := json.NewDecoder(c)
	for {
		var m Message
		err := d.Decode(&m)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("%#v\n", m)
	}
}
