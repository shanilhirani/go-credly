// Package fetch performs API requests
package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/shanilhirani/go-credly/pkgs/types"
)

var (
	ErrMissingRequiredParam = func(errorType string) error {
		return fmt.Errorf("missing %s parameter", errorType)
	}
	ErrFailedToParse = func(field string) error {
		return fmt.Errorf("failed to parse %s", field)
	}
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
// It takes a parameter 'username' which is the username or user ID for which the data needs to be fetched
// It returns a pointer to a CredlyData struct and an error if any occurs during the request or unmarshalling process
// For example: client.Fetch(param).DoSomething()
func (c *Client) Fetch(username string) (*types.CredlyData, error) {
	url := fmt.Sprintf("https://api.credly.com/users/%s/badges.json", username)
	req, err := http.NewRequest("GET", url, nil) //nolint:noctx // context is not needed for this example
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	resp, err := c.transport.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GET request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Read the response body for error message
		return nil, fmt.Errorf("received non-OK response code: %d, body: %s", resp.StatusCode, body)
	}

	defer resp.Body.Close() //nolint:errcheck // no comment

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

// FilteredBadge represents a filtered version of the Badge struct
// It contains only the required fields from the Badge struct
// This struct is used for filtering the data returned by the Credly API
// For example: filteredBadges, err := FilterData("john_doe") // returns a slice of FilteredBadge structs
type FilteredBadge struct {
	BadgeName          string `json:"name"`
	BadgeDescription   string `json:"description"`
	BadgeImageURL      string `json:"image_url"`
	BadgeURL           string `json:"url"`
	BadgeExpiresAtDate string `json:"expires_at_date"`
}

// FilterData filters the CredlyData struct to include only the required fields
// It takes a parameter 'username' which is the username or user ID for which the data needs to be filtered
// It returns a slice of filtered Badge structs and an error if any occurs during the filtering process
// For example: filteredBadges, err := FilterData("john_doe")
func FilterData(username string, data *types.CredlyData, includeExpired bool) ([]FilteredBadge, error) {
	var filteredBadges []FilteredBadge
	now := time.Now()

	// Iterate over the slice of anonymous structs in CredlyData
	for _, badge := range data.Data {
		// Check if the badge is issued to the specified user
		if badge.EarnerPath != fmt.Sprintf("/users/%s", username) {
			continue
		}

		// Parse the ExpiresAtDate string into a time.Time value
		expiresAtDate, err := parseExpiresAtDate(badge.ExpiresAtDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse ExpiresAtDate for badge ID %s: %v", badge.ID, err)
		}

		// Check if the badge is expired or if we want to include expired badges
		if includeExpired || expiresAtDate.After(now) {
			filteredBadge := FilteredBadge{
				BadgeName:          badge.BadgeTemplate.Name,
				BadgeDescription:   badge.BadgeTemplate.Description,
				BadgeImageURL:      badge.BadgeTemplate.ImageURL,
				BadgeURL:           badge.BadgeTemplate.URL,
				BadgeExpiresAtDate: badge.ExpiresAtDate,
			}
			filteredBadges = append(filteredBadges, filteredBadge)
		}
	}

	if len(filteredBadges) == 0 {
		return nil, fmt.Errorf("no badges found for user %s", username)
	}

	return filteredBadges, nil
}

var utcLoc *time.Location

func parseExpiresAtDate(expiresAtDate string) (time.Time, error) {
	if expiresAtDate == "" {
		return time.Time{}, ErrMissingRequiredParam("expiresAtDate")
	}

	if utcLoc == nil {
		var err error
		utcLoc, err = time.LoadLocation("UTC")
		if err != nil {
			return time.Time{}, fmt.Errorf("failed to load UTC location: %w", err)
		}
	}

	parsedTime, err := time.ParseInLocation("2006-01-02", expiresAtDate, utcLoc)
	if err != nil {
		var zeroTime time.Time
		return zeroTime, ErrFailedToParse("expiresAtDate")
	}

	return parsedTime, nil
}
