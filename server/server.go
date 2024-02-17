package main

import (
	"crypto/tls"
	"log/slog"

	"github.com/tesla59/shepherd/common"
)

var workerQueue = make(chan *tls.Conn, common.SERVER_WORKERS)

func main() {
	// Get the tls.Config object to establish a secure tcp listener
	config, err := GetTLSConfig()
	if err != nil {
		slog.Error("Getting TLS config", "error", err.Error())
		return
	}

	l, err := tls.Listen(common.CONN_PROTO, common.CONN_ADDR, config)
	if err != nil {
		slog.Error("Setting up listener", "error", err.Error())
	}
	defer l.Close()

	slog.Info("Server Started")

	// Spawn a pool of workers to handle incoming connections
	spawnWorker()

	for {
		conn, err := l.Accept()
		if err != nil {
			slog.Error("Accepting connection from client", "error", err.Error(), "client", conn.RemoteAddr())
			continue
		}
		// Pass the tls.Conn to the workerQueue
		// The workerQueue is tracked by the goroutines
		workerQueue <- conn.(*tls.Conn)
	}
}
