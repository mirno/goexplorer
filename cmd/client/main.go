package main

import (
	"fmt"
	"log"

	"github.com/mirno/goexplorer/pkg/config"
	"github.com/mirno/goexplorer/pkg/utils"
	"github.com/spf13/viper"
)

// var(
// 	caCertificatePath = ".private/ca.crt"
// 	caCertificatekeyPath = ".private/ca.key"
// 	serverCertificatePath = ".private/server.crt"
// 	serverCertificateKeyPath = ".private/server.key"
// )
func main() {
    config.LoadDefaultConfig() // Load default.yaml configurations

    // Access a default configuration directly
    proxy := viper.GetString("proxy")
    fmt.Println("Proxy from default.yaml:", proxy)

	conf, err := config.InitializeConfig("config.dt.yaml")
	utils.Assert(err)
    log.Printf("Loaded Config: %+v\n", conf)



	// if err := viper.ReadInConfig(); err != nil {
    //     log.Fatalf("Error reading config file: %s", err)
    // }

    // // Attempt to read the configuration
    // if err := viper.ReadInConfig(); err != nil {
    //     log.Fatalf("Error reading config file: %s", err)
    // }	
	// // Unmarshall the configuration into the Config struct
	// var conf config.Config

	// if err := viper.Unmarshall(&conf); err != nil {
	// 	log.Fatalf("Error unmarshalling config: %s", err)
	// }


	
	// caCert, err := os.ReadFile(caCertificatePath)
	// if err != nil{
	// 	log.Fatalf("Reading server certificate: %s", err)
	// }
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	// clientCert, err := tls.LoadX509KeyPair(serverCertificatePath, serverCertificateKeyPath)
	// if err != nil {
	// 	log.Fatalf("Error loading server key par: %s", err)
	// }

	// tlsConfig := &tls.Config{
	// 	RootCAs: caCertPool,
	// 	Certificates: []tls.Certificate{clientCert},
	// }

	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: tlsConfig,
	// 	},
	// }

	// resp, err := client.Get("https://localhost:8443")
	// if err != nil {
	// 	log.Fatalf("Reading response body: %s", err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     log.Fatalf("Reading response body: %s", err)
    // }
	// log.Printf("Server response: %s", body)
}