package main

import (
	"crypto/tls"
	"fmt"
	"log/slog"
)

const (
	CONN_HOST  = "0.0.0.0"
	CONN_PORT  = "5202"
	CONN_PROTO = "tcp"
	WORKERS    = 5
)

var workerQueue = make(chan *tls.Conn, WORKERS)

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

	spawnWorker()

	for {
		conn, err := l.Accept()
		if err != nil {
			slog.Error("Accepting connection from client", "error", err.Error(), "client", conn.RemoteAddr())
			continue
		}
		// go handleConn(conn.(*tls.Conn))
		workerQueue <- conn.(*tls.Conn)
	}
}
