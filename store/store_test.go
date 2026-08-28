package store

import (
	"memorial/model"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "x.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveEvent(model.NewEvent("e1", "p1", "day", "2030-01-01")); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetEvent("e1"); e != nil {
		t.Fatal(e)
	}
}
