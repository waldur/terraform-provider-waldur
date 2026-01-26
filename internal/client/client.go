package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client is a Waldur API client
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// Config holds the client configuration
type Config struct {
	Endpoint   string
	Token      string
	HTTPClient *http.Client // Optional: for testing with VCR or custom transport
}

// NewClient creates a new Waldur API client
func NewClient(config *Config) (*Client, error) {
	if config.Endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}
	if config.Token == "" {
		return nil, fmt.Errorf("token is required")
	}

	// Parse and validate endpoint URL
	baseURL, err := url.Parse(config.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}

	// Use provided HTTP client or create default one
	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &Client{
		baseURL:    baseURL.String(),
		token:      config.Token,
		httpClient: httpClient,
	}, nil
}

// doRequest performs an HTTP request with authentication
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	// Construct full URL, avoiding double slashes
	fullURL := strings.TrimSuffix(c.baseURL, "/") + path

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, fullURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Token "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// GetURL performs a GET request
func (c *Client) GetURL(ctx context.Context, path string, result interface{}) error {
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := c.checkResponse(resp); err != nil {
		return err
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	resp, err := c.doRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := c.checkResponse(resp); err != nil {
		return err
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// Patch performs a PATCH request
func (c *Client) Patch(ctx context.Context, path string, body interface{}, result interface{}) error {
	resp, err := c.doRequest(ctx, http.MethodPatch, path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := c.checkResponse(resp); err != nil {
		return err
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// DeleteURL performs a DELETE request
func (c *Client) DeleteURL(ctx context.Context, path string) error {
	resp, err := c.doRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return c.checkResponse(resp)
}

// checkResponse checks the HTTP response for errors
func (c *Client) checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	// Try to read error message from response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HTTP %d: failed to read error response", resp.StatusCode)
	}

	// Try to parse as JSON error
	var errorResp map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &errorResp); err == nil {
		return fmt.Errorf("HTTP %d: %v", resp.StatusCode, errorResp)
	}

	// Return raw body if not JSON
	return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(bodyBytes))
}

// List performs a GET request with query parameters for filtering
func (c *Client) List(ctx context.Context, path string, filters map[string]string, result interface{}) error {
	// Build query string from filters
	if len(filters) > 0 {
		query := url.Values{}
		for key, value := range filters {
			query.Add(key, value)
		}
		path = path + "?" + query.Encode()
	}
	return c.GetURL(ctx, path, result)
}

// Get retrieves a single resource by UUID
func (c *Client) Get(ctx context.Context, path string, uuid string, result interface{}) error {
	// Replace {uuid} placeholder in path, or append if not present
	fullPath := strings.Replace(path, "{uuid}", uuid, 1)
	// Ensure trailing slash
	if !strings.HasSuffix(fullPath, "/") {
		fullPath += "/"
	}
	return c.GetURL(ctx, fullPath, result)
}

// Update updates a resource
func (c *Client) Update(ctx context.Context, path string, uuid string, body interface{}, result interface{}) error {
	// Replace {uuid} placeholder in path, or append if not present
	fullPath := strings.Replace(path, "{uuid}", uuid, 1)
	// Ensure trailing slash
	if !strings.HasSuffix(fullPath, "/") {
		fullPath += "/"
	}
	return c.Patch(ctx, fullPath, body, result)
}

// Delete deletes a resource by UUID
func (c *Client) Delete(ctx context.Context, path string, uuid string) error {
	// Replace {uuid} placeholder in path, or append if not present
	fullPath := strings.Replace(path, "{uuid}", uuid, 1)
	// Ensure trailing slash
	if !strings.HasSuffix(fullPath, "/") {
		fullPath += "/"
	}
	return c.DeleteURL(ctx, fullPath)
}

// ExecuteAction executes an action on a resource
func (c *Client) ExecuteAction(ctx context.Context, pathTemplate string, uuid string, body interface{}, result interface{}) error {
	path := strings.Replace(pathTemplate, "{uuid}", uuid, 1)
	return c.Post(ctx, path, body, result)
}
