package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/4nd3rs0n/oapi-test/pkg/api"
)

func main() {
	url := "http://127.0.0.1:9090"
	client, err := api.NewClient(url)
	if err != nil {
		log.Fatalf("error initializing the client (%s): %v", url, err)
	}
	resp, err := client.GetStatus(context.Background())
	if err != nil {
		log.Fatalf("error getting status: %v", err)
	}

	respMap := make(map[string]any)
	json.NewDecoder(resp.Body).Decode(&respMap)
	fmt.Printf("response code: %d; response: %+v", resp.StatusCode, respMap)
}
