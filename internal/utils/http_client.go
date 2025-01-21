package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient is a wrapper for HTTP requests
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTPClient instance with a timeout
func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get executes a GET request and returns the response body
func (c *HTTPClient) Get(path string) ([]byte, error) {
	resp, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
