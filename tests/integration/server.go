package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// StartTestServer starts a local server serving files from tests/fixtures
func StartTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	// Serve static files from fixtures directory
	fs := http.FileServer(http.Dir("../fixtures"))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[TestServer] Request: %s\n", r.URL.Path)
		fs.ServeHTTP(w, r)
	}))

	t.Cleanup(func() {
		server.Close()
	})

	return server
}
