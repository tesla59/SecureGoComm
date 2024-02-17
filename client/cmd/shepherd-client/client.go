package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/tesla59/shepherd/common"
)

func main() {
	conn, err := net.Dial(common.CONN_PROTO, common.CONN_ADDR)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer conn.Close()

	slog.SetDefault(slog.With("remote", conn.RemoteAddr()))

	slog.Info("Connection established with server")

	// reader for reading input from the user
	// rw for reading and writing to the server
	reader := bufio.NewReader(os.Stdin)
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		fmt.Printf("~ ")
		text, err := reader.ReadString('\n')
		if err != nil {
			slog.Warn(err.Error())
			os.Exit(1)
		}
		_, err = rw.WriteString(text)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		err = rw.Flush()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}
}
