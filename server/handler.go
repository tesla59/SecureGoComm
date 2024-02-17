package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log/slog"
	"strings"
)

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

func spawnWorker(workerQueue chan *tls.Conn) {
	for conn := range workerQueue {
		handleConn(conn)
	}
}
