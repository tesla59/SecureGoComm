package main

import (
	"bufio"
	"log"
	"net"
)

const (
	CONN_HOST  = "0.0.0.0"
	CONN_PORT  = "5202"
	CONN_PROTO = "tcp"
)

func main() {
	l, err := net.Listen(CONN_PROTO, CONN_HOST+":"+CONN_PORT)
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
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		log.Print(str)
	}
}
