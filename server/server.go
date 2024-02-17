package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"strings"
)

const (
	CONN_HOST  = "0.0.0.0"
	CONN_PORT  = "5202"
	CONN_PROTO = "tcp"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../certificates/server/server.crt", "../certificates/server/server.key")
	if err != nil {
		slog.Error("Loading server certificate", "error", err.Error())
		return
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	l, err := tls.Listen(CONN_PROTO, fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT), config)
	if err != nil {
		slog.Error("Setting up listener", "error", err.Error())
	}
	defer l.Close()

	slog.Info("Server Started")

	for {
		conn, err := l.Accept()
		if err != nil {
			slog.Error("Accepting connection from client", "error", err.Error(), "client", conn.RemoteAddr())
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	slog.Info("Connection established with client", "client", conn.RemoteAddr())
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	_, err := rw.WriteString(fmt.Sprintf("Connection established with %s\n", conn.RemoteAddr().String()))
	if err != nil {
		slog.Warn("Cannot write initial byte to client", "error", err.Error())
	}
	err = rw.Flush()
	if err != nil {
		slog.Warn("Cannot flush writer to client", "error", err.Error())
	}
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			slog.Error("Cannot read string from client", "error", err.Error())
			return
		}
		slog.Info(strings.TrimRight(str, "\n"), "client", conn.RemoteAddr())
	}
}
