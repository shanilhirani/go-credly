// Package fetch performs API requests
package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/shanilhirani/go-credly/pkgs/types"
)

// HTTPClient is an interface for making HTTP requests
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client is a struct that holds an HTTPClient implementation
type Client struct {
	transport HTTPClient
}

// NewClient creates a new instance of the Client struct with the provided HTTPClient
// If no HTTPClient is provided, http.DefaultClient is used
func NewClient(transport HTTPClient) *Client {
	if transport == nil {
		transport = http.DefaultClient
	}
	return &Client{transport: transport}
}

// Fetch performs a GET request to the specified URL and returns the response data as a CredlyData struct
// It takes a parameter 'param' which is the username or user ID for which the data needs to be fetched
// It returns a pointer to a CredlyData struct and an error if any occurs during the request or unmarshalling process
// For example: client.Fetch(param).DoSomething()
func (c *Client) Fetch(param string) (*types.CredlyData, error) {
	url := fmt.Sprintf("https://api.credly.com/users/%s/badges.json", param)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	resp, err := c.transport.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Read the response body for error message
		return nil, fmt.Errorf("received non-OK response code: %d, body: %s", resp.StatusCode, body)
	}

	var result types.CredlyData
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		var syntaxError *json.SyntaxError
		if errors.As(err, &syntaxError) {
			return nil, fmt.Errorf("failed to unmarshal JSON response at byte offset %d: %w", syntaxError.Offset, err)
		}
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return &result, nil
}

type FilteredBadge struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	URL         string `json:"url"`
}

// FilterData filters the CredlyData struct to include only the required fields
// It takes a parameter 'param' which is the username or user ID for which the data needs to be filtered
// It returns a slice of filtered Badge structs and an error if any occurs during the filtering process
// For example: filteredBadges, err := FilterData("john_doe")
func FilterData(param string, data *types.CredlyData) ([]FilteredBadge, error) {
	var filteredBadges []FilteredBadge

	// Iterate over the slice of anonymous structs in CredlyData
	for _, badge := range data.Data {
		// Check if the badge is issued to the specified user
		if badge.EarnerPath == fmt.Sprintf("/users/%s", param) {
			filteredBadge := FilteredBadge{
				Name:        badge.BadgeTemplate.Name,
				Description: badge.BadgeTemplate.Description,
				ImageURL:    badge.BadgeTemplate.ImageURL,
				URL:         badge.BadgeTemplate.URL,
			}
			filteredBadges = append(filteredBadges, filteredBadge)
		}
	}

	if len(filteredBadges) == 0 {
		return nil, fmt.Errorf("no badges found for user %s", param)
	}

	return filteredBadges, nil
}
