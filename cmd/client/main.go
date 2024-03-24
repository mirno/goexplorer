package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mirno/goexplorer/internal/drivers"
	"github.com/mirno/goexplorer/pkg/config"
	"github.com/mirno/goexplorer/pkg/utils"
	"github.com/spf13/viper"
)

const baseURL = "https://localhost:8443"


func main() {
    config.LoadDefaultConfig() // Load default.yaml configurations

    // Access a default configuration directly
    proxy := viper.GetString("proxy")
    fmt.Println("Proxy from default.yaml:", proxy)

	conf, err := config.InitializeConfig("config.dt.yaml")
	utils.Assert(err)
    // log.Printf("Loaded Config: %+v\n", conf)

	pool := config.NewCertificatePool(conf)
	cert := config.LoadMTLSClientCert(&conf.Server.AuthenticationType.MTLS)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs: pool,
	}
	httpClient := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: tlsConfig,
        },
    }

    // Push an item
    pushItem(httpClient, "test item")

    // Peek at the top item
    peekItem(httpClient)

    // Pop the top item
    popItem(httpClient)

    // Peek again to verify the pop
    peekItem(httpClient)
	// Driver implementation
	// Initialize the StackClient
	client, err := drivers.NewStackClient(baseURL, cert, pool)
	utils.Assert(err)

    err = client.PushItem("bla")
    utils.Assert(err)

	item, err := client.PeekItem()
    utils.Assert(err)
    fmt.Printf("Peeked item: %+v\n", item)

    err = client.PushItem(2)
    utils.Assert(err)

	item, err = client.PeekItem()
    utils.Assert(err)
    fmt.Printf("Peeked item: %+v\n", item)
}

func pushItem(client *http.Client, value interface{}) {
    item := map[string]interface{}{"value": value}
    jsonData, err := json.Marshal(item)
    if err != nil {
        log.Fatalf("Error marshalling JSON: %s", err)
    }

    // Create a new request with PUT method
    req, err := http.NewRequest(http.MethodPut, baseURL+"/stack", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Fatalf("Error creating PUT request: %s", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Send the request using the client
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error pushing item with PUT request: %s", err)
    }
    defer resp.Body.Close()

    // Check the response status code
    fmt.Println("Pushed item, status code:", resp.StatusCode)
}

func peekItem(client *http.Client) {
    resp, err := client.Get(baseURL + "/stack")
    if err != nil {
        log.Fatalf("Error peeking item: %s", err)
    }
    defer resp.Body.Close()

    fmt.Printf("Peek response status code: %d\n", resp.StatusCode)

    if resp.StatusCode == http.StatusNoContent {
        fmt.Println("The stack is empty.")
        return
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %s", err)
    }

    fmt.Printf("Peeked item: %s\n", string(body))
}

func popItem(client *http.Client) {
    req, err := http.NewRequest("POST", baseURL+"/stack", nil)
    if err != nil {
        log.Fatalf("Error creating request for pop: %s", err)
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error popping item: %s", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %s", err)
    }

    fmt.Printf("Popped item: %s\n", string(body))
}