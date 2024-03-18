package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var(
	caCertificatePath = ".private/ca.crt"
	caCertificatekeyPath = ".private/ca.key"
	serverCertificatePath = ".private/server.crt"
	serverCertificateKeyPath = ".private/server.key"
)
func main() {
	caCert, err := os.ReadFile(caCertificatePath)
	if err != nil{
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair(serverCertificatePath, serverCertificateKeyPath)
	if err != nil {
		log.Fatalf("Error loading server key par: %s", err)
	}

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
		Certificates: []tls.Certificate{clientCert},
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Fatalf("Reading response body: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Reading response body: %s", err)
    }
	log.Printf("Server response: %s", body)
}