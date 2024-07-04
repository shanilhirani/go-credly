// go-credly app
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/shanilhirani/go-credly/pkgs/types"
)

// HTTPClient create a client connection to the URL parameter
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// Client structure
type Client struct{}

// Get implements HTTPClient.
func (c *Client) Get(url string) (*http.Response, error) {
	return http.Get(url) // #nosec G107 -- required for username/user id
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go [Credly Username/ID]")
	}
	arg := os.Args[1]

	client := &Client{}
	result, err := fetchData(client, arg)
	if err != nil {
		log.Fatalf("Error fetching data: %v\n", err)
	}

	formattedJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Failed to format JSON output: %v\n", err)
	}

	fmt.Println(string(formattedJSON))
}

func fetchData(client HTTPClient, param string) (*types.CredlyData, error) {
	url := fmt.Sprintf("https://api.credly.com/users/%s/badges.json/", param)
	response, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GET request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result *types.CredlyData

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return result, nil
}
