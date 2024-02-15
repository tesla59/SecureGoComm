package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST  = "0.0.0.0"
	CONN_PORT  = "5202"
	CONN_PROTO = "tcp"
)

func main() {
	l, err := net.Listen(CONN_PROTO, fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT))
	log.Println("Server Started")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	_, err := rw.WriteString(fmt.Sprintf("Connection established with %s\n", conn.RemoteAddr().String()))
	if err != nil {
		log.Println(err)
		return
	}
	err = rw.Flush()
	if err != nil {
		log.Println(err)
		return
	}
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		log.Print(str)
	}
}
