package api

import (
	"memorial/service"
	"memorial/store"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealth(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "a"))
	defer s.Close()
	r := httptest.NewRecorder()
	New(service.New(s)).Routes().ServeHTTP(r, httptest.NewRequest("GET", "/health", nil))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
