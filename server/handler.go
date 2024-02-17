package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log/slog"
	"strings"
)

// handleConn is a function that handles the incoming connection
// It reads the incoming data from the client and logs it
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

// spawnWorker creates a pool of workers to handle incoming connections
// Which are then passed to the workerQueue
// workerQueue is tracked by the goroutines and they handle the connections
func spawnWorker() {
	for range WORKERS {
		go func() {
			for conn := range workerQueue {
				handleConn(conn)
			}
		}()
	}
}
