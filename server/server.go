package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log/slog"
	"strings"
)

const (
	CONN_HOST  = "0.0.0.0"
	CONN_PORT  = "5202"
	CONN_PROTO = "tcp"
)

func main() {
	config, err := GetTLSConfig()
	if err != nil {
		slog.Error("Getting TLS config", "error", err.Error())
		return
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
		go handleConn(conn.(*tls.Conn))
	}
}

func handleConn(conn *tls.Conn) {
	defer conn.Close()
	if err := conn.Handshake(); err != nil {
		slog.Error("Handshaking with client", "error", err.Error())
		return
	}
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
