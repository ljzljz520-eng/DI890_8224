package reminder

import (
	"memorial/model"
	"memorial/store"
	"path/filepath"
	"testing"
)

func TestConfirm(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "a"))
	defer s.Close()
	s.SaveEvent(model.NewEvent("e", "p", "x", "2030-01-01"))
	m := New(s)
	if _, e := m.Confirm("e", "a", "note"); e != nil {
		t.Fatal(e)
	}
}
