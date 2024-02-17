package main

import (
	"crypto/tls"
	"crypto/x509"
	"log/slog"
	"os"
)

func GetTLSConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("../certificates/server/server.crt", "../certificates/server/server.key")
	if err != nil {
		slog.Error("Loading server certificate", "error", err.Error())
		return nil, err
	}
	caCert, err := os.ReadFile("../certificates/ca/ca.crt")
	if err != nil {
		slog.Error("Loading CA certificate", "error", err.Error())
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		slog.Error("Couldn't add CA cert to pool", "error", err.Error())
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
		ClientCAs: caCertPool,
	}, nil
}
