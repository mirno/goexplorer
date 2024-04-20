package drivers

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type StackClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewStackClient(baseURL string, cert tls.Certificate, caPool  *x509.CertPool ) (*StackClient, error) {
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs: caPool,
	}

	httpClient := &http.Client{
		Transport: &http.Transport{TLSClientConfig: tlsConfig},
	}

	return &StackClient{
		BaseURL: baseURL,
		HTTPClient: httpClient,
	}, nil
}


func (client *StackClient) PeekItem() (interface{}, error) {
    resp, err := client.HTTPClient.Get(client.BaseURL + "/stack")
    if err != nil {
        return nil, fmt.Errorf("error making request to peek item: %w", err)
    }
    defer resp.Body.Close()

    // Check for no content status
    if resp.StatusCode == http.StatusNoContent {
		log.Println("The stack is empty.")
        return nil, nil // Or consider a custom error/message indicating the stack is empty
    }

    var item struct {
        Value interface{} `json:"value"`
    }
    err = json.NewDecoder(resp.Body).Decode(&item)
    if err != nil {
        return nil, fmt.Errorf("error unmarshalling response body: %w", err)
    }
	log.Printf("Peeked item: %v\n", (item.Value))
    return item.Value, nil
}

func (client *StackClient) PushItem(value interface{}) error {
	item := map[string]interface{}{"value": value}
    jsonData, err := json.Marshal(item)
    if err != nil {
        log.Fatalf("Error marshalling JSON: %s", err)
    }

    // Create a new request with PUT method
    req, err := http.NewRequest(http.MethodPut, client.BaseURL+"/stack", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Fatalf("Error creating PUT request: %s", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Send the request using the client
    resp, err := client.HTTPClient.Do(req)
    if err != nil {
        log.Fatalf("Error pushing item with PUT request: %s", err)
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("received non-201 response status: %d", resp.StatusCode)
	}
	log.Println("Pushed item, status code:", resp.StatusCode)
	log.Printf("Item: %v, pushed successfully", (value))
	return nil
}

// Assuming POST is used for Pop for this example, adjust as needed
func (client *StackClient) PopItem() (interface{}, error) {
	resp, err := client.HTTPClient.Post(client.BaseURL+"/stack", "application/json", nil)
	if err != nil {
		return nil, fmt.Errorf("error making request to pop item: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// var item struct {
	// 	Value interface{} `json:"value"`
	// }
	// if err := json.Unmarshal(body, &item); err != nil {
	// 	return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	// }
    log.Printf("Popped item: %s\n", string(body))
	return body, nil
}