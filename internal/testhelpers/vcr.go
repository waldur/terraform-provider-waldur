package testhelpers

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

// SetupVCR creates a VCR recorder for recording/replaying HTTP interactions.
// By default, it runs in replay-only mode. Set VCR_RECORD=true to record new cassettes.
//
// Usage:
//
//	rec, cleanup := testhelpers.SetupVCR(t, "my_test_cassette")
//	defer cleanup()
//	client := &http.Client{Transport: rec}
func SetupVCR(t *testing.T, cassetteName string) (*recorder.Recorder, func()) {
	t.Helper()

	mode := recorder.ModeReplayOnly

	// Check if we should record new cassettes
	if os.Getenv("VCR_RECORD") == "true" {
		mode = recorder.ModeRecordOnce
	}

	// Cassettes are stored in testdata/
	// Note: go-vcr automatically appends .yaml extension
	// We need to find the project root for cassettes to work from any package
	cassettePath := filepath.Join("testdata", cassetteName)

	// If the cassette path doesn't exist, try to find from project root
	if _, err := os.Stat(cassettePath + ".yaml"); os.IsNotExist(err) {
		// Try from current directory up to find project root (where go.mod is)
		wd, _ := os.Getwd()
		for wd != "/" {
			testPath := filepath.Join(wd, cassettePath)
			if _, err := os.Stat(testPath + ".yaml"); err == nil {
				cassettePath = testPath
				break
			}
			wd = filepath.Dir(wd)
		}
	}

	r, err := recorder.NewWithOptions(&recorder.Options{
		CassetteName: cassettePath,
		Mode:         mode,
	})
	if err != nil {
		t.Fatalf("Failed to create VCR recorder: %v", err)
	}

	// Cleanup function to stop the recorder
	cleanup := func() {
		if err := r.Stop(); err != nil {
			t.Errorf("Failed to stop VCR recorder: %v", err)
		}
	}

	return r, cleanup
}

// GetHTTPClient creates an HTTP client configured with VCR for testing.
// This is a convenience wrapper around SetupVCR for tests that just need a client.
func GetHTTPClient(t *testing.T, cassetteName string) (*http.Client, func()) {
	t.Helper()

	rec, cleanup := SetupVCR(t, cassetteName)
	client := &http.Client{
		Transport: rec,
	}

	return client, cleanup
}
