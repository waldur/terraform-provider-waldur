package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListWithFilter(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify query parameters
		if r.URL.Query().Get("name") != "test-project" {
			t.Errorf("Expected name=test-project, got %s", r.URL.Query().Get("name"))
		}

		// Return test data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"uuid": "abc-123", "name": "test-project"}]`))
	}))
	defer server.Close()

	// Create client
	client, err := NewClient(&Config{
		Endpoint: server.URL,
		Token:    "test-token",
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test ListWithFilter
	var results []map[string]interface{}
	filters := map[string]string{
		"name": "test-project",
	}

	err = client.ListWithFilter(context.Background(), "/api/projects/", filters, &results)
	if err != nil {
		t.Fatalf("ListWithFilter failed: %v", err)
	}

	// Verify results
	if len(results) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(results))
	}

	if results[0]["uuid"] != "abc-123" {
		t.Errorf("Expected uuid=abc-123, got %v", results[0]["uuid"])
	}

	if results[0]["name"] != "test-project" {
		t.Errorf("Expected name=test-project, got %v", results[0]["name"])
	}
}

func TestGetByUUID_WithTemplate(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify path has UUID substituted
		expectedPath := "/api/projects/abc-123/"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Return test data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"uuid": "abc-123", "name": "test-project"}`))
	}))
	defer server.Close()

	// Create client
	client, err := NewClient(&Config{
		Endpoint: server.URL,
		Token:    "test-token",
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test GetByUUID with templated path
	var result map[string]interface{}
	err = client.GetByUUID(context.Background(), "/api/projects/{uuid}/", "abc-123", &result)
	if err != nil {
		t.Fatalf("GetByUUID failed: %v", err)
	}

	// Verify result
	if result["uuid"] != "abc-123" {
		t.Errorf("Expected uuid=abc-123, got %v", result["uuid"])
	}
}

func TestUpdate_WithTemplate(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify method
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH, got %s", r.Method)
		}

		// Verify path has UUID substituted
		expectedPath := "/api/projects/abc-123/"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Return test data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"uuid": "abc-123", "name": "updated-project"}`))
	}))
	defer server.Close()

	// Create client
	client, err := NewClient(&Config{
		Endpoint: server.URL,
		Token:    "test-token",
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test Update with templated path
	var result map[string]interface{}
	body := map[string]string{"name": "updated-project"}

	err = client.Update(context.Background(), "/api/projects/{uuid}/", "abc-123", body, &result)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	// Verify result
	if result["name"] != "updated-project" {
		t.Errorf("Expected name=updated-project, got %v", result["name"])
	}
}

func TestDeleteByUUID_WithTemplate(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify method
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}

		// Verify path has UUID substituted
		expectedPath := "/api/projects/abc-123/"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	// Create client
	client, err := NewClient(&Config{
		Endpoint: server.URL,
		Token:    "test-token",
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test DeleteByUUID with templated path
	err = client.DeleteByUUID(context.Background(), "/api/projects/{uuid}/", "abc-123")
	if err != nil {
		t.Fatalf("DeleteByUUID failed: %v", err)
	}
}
