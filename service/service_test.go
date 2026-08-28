package service

import (
	"memorial/model"
	"memorial/store"
	"path/filepath"
	"testing"
)

func TestRegister(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "a"))
	defer s.Close()
	x := New(s)
	if e := x.Register(model.NewProfile("p", "n", "r", "2030-01-01"), model.NewEvent("e", "p", "x", "2030-01-01")); e != nil {
		t.Fatal(e)
	}
}
