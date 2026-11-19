package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

// personalAccessTokenPrefix starts every Waldur personal access token.
const personalAccessTokenPrefix = "w_"

// AuthorizationHeader returns the Authorization header value for a token.
// Waldur takes a personal access token as a Bearer token and an account token
// under the Token scheme; sent under the other scheme, either is refused with 401.
func AuthorizationHeader(token string) string {
	if strings.HasPrefix(token, personalAccessTokenPrefix) {
		return "Bearer " + token
	}
	return "Token " + token
}

// doRequest performs an HTTP request with authentication
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	// Construct full URL, avoiding double slashes and double 'api' segments
	baseURL := strings.TrimSuffix(c.baseURL, "/")
	if strings.HasSuffix(baseURL, "/api") && strings.HasPrefix(path, "/api/") {
		baseURL = strings.TrimSuffix(baseURL, "/api")
	}
	fullURL := baseURL + path

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
	req.Header.Set("Authorization", AuthorizationHeader(c.token))
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

// listPageSize is the page size ListAll asks for.
const listPageSize = 100

// maxListPages bounds ListAll, so that a server which never stops offering a
// next page fails the read instead of looping forever.
const maxListPages = 1000

// ListAll is List across every page: the API returns one page per request, so
// List alone sees at most the first page of a long result.
//
// It follows the API's Link header pagination: it asks for the next page only
// while the response links one (rel="next"), and stops once it holds
// X-Result-Count items. Asking past the last page is an error (404 "Invalid
// page."), so a full last page must not be taken to mean there is more. An
// endpoint that does not paginate sends no next link and is read once; one that
// ignores the page parameter yet links a next page is caught when a page starts
// with the same item as the page before it.
func (c *Client) ListAll(ctx context.Context, path string, filters map[string]string, result interface{}) error {
	all := []json.RawMessage{}
	previousFirst := ""
	for page := 1; ; page++ {
		if page > maxListPages {
			return fmt.Errorf("listing %s: more than %d pages of %d items", path, maxListPages, listPageSize)
		}
		query := url.Values{}
		for key, value := range filters {
			query.Set(key, value)
		}
		query.Set("page", fmt.Sprintf("%d", page))
		query.Set("page_size", fmt.Sprintf("%d", listPageSize))

		items, header, err := c.getPage(ctx, path+"?"+query.Encode())
		if err != nil {
			return err
		}
		if len(items) == 0 {
			break
		}
		if page > 1 && string(items[0]) == previousFirst {
			break // the page parameter is ignored: this page repeats the last one
		}
		previousFirst = string(items[0])
		all = append(all, items...)

		if !hasNextLink(header.Get("Link")) {
			break
		}
		if total, err := strconv.Atoi(header.Get("X-Result-Count")); err == nil && len(all) >= total {
			break
		}
	}
	data, err := json.Marshal(all)
	if err != nil {
		return fmt.Errorf("failed to collect list results: %w", err)
	}
	return json.Unmarshal(data, result)
}

// getPage reads one page of a list along with the response headers, which
// carry the pagination links.
func (c *Client) getPage(ctx context.Context, path string) ([]json.RawMessage, http.Header, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if err := c.checkResponse(resp); err != nil {
		return nil, nil, err
	}
	var items []json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return items, resp.Header, nil
}

// hasNextLink reports whether an RFC 8288 Link header links a next page.
func hasNextLink(link string) bool {
	for _, part := range strings.Split(link, ",") {
		for _, param := range strings.Split(part, ";")[1:] {
			if strings.EqualFold(strings.ReplaceAll(strings.TrimSpace(param), " ", ""), `rel="next"`) {
				return true
			}
		}
	}
	return false
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
