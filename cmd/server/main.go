package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"strconv"
)
var(
	caCertificatePath = ".private/ca.crt"
	caCertificatekeyPath = ".private/ca.key"
	serverCertificatePath = ".private/server.crt"
	serverCertificateKeyPath = ".private/server.key"

	port = 8443

)
func main() {
	caCert, err := os.ReadFile(caCertificatePath)
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	serverCert, err:= tls.LoadX509KeyPair(serverCertificatePath, serverCertificateKeyPath)
	if err != nil {
		log.Fatalf("Error loading server key par: %s", err)
	}

	tlsConfig := &tls.Config{
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{serverCert},
	}
	// tlsConfig.BuildNameToCertificate()


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, mTLS world!"))
	})
	server := &http.Server{
		Addr: ":"+ strconv.Itoa(port),
		TLSConfig: tlsConfig,	
	}

    log.Printf("Starting Secure Server on https://localhost:" + strconv.Itoa(port))
	// Start the server with TLS, this uses the default ServeMux
	log.Fatal(server.ListenAndServeTLS("", ""))  // Empty strings will use the certificates from TLSConfig

}