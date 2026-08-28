package query

import (
	"memorial/model"
	"memorial/store"
	"path/filepath"
	"testing"
)

func TestUpcoming(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "a"))
	defer s.Close()
	s.SaveEvent(model.NewEvent("e", "p", "x", "2030-01-01"))
	x, e := Upcoming(s)
	if e != nil || len(x) != 1 {
		t.Fatal(x, e)
	}
}
