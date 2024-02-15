package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	CONN_HOST  = "0.0.0.0"
	CONN_PORT  = "5202"
	CONN_PROTO = "tcp"
)

func main() {
	conn, err := net.Dial(CONN_PROTO, fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	log.Println("Connection established with server")

	// reader for reading input from the user
	// rw for reading and writing to the server
	reader := bufio.NewReader(os.Stdin)
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		fmt.Printf("~ ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = rw.WriteString(text)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = rw.Flush()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
