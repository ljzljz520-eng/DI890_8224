package memorial

import (
	"memorial/model"
	"memorial/service"
	"memorial/store"
	"path/filepath"
	"testing"
)

func setup(t *testing.T) *service.Service {
	s, _ := store.Open(filepath.Join(t.TempDir(), "a"))
	t.Cleanup(func() { s.Close() })
	return service.New(s)
}
func TestWorkflowOne(t *testing.T) {
	s := setup(t)
	if e := s.Register(model.NewProfile("p", "Alice", "family", "2030-01-01"), model.NewEvent("e", "p", "Anniversary", "2030-01-01")); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s := setup(t)
	s.Register(model.NewProfile("p", "A", "f", "2030-01-01"), model.NewEvent("e", "p", "X", "2030-01-01"))
	if _, e := s.Confirm("e", "a", "one"); e != nil {
		t.Fatal(e)
	}
	if e := s.Archive("e", "a"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s := setup(t)
	s.Register(model.NewProfile("p", "A", "f", "2030-01-01"), model.NewEvent("e", "p", "X", "2030-01-01"))
	if _, e := s.Confirm("e", "a", "note"); e != nil {
		t.Fatal(e)
	}
}
func TestBusinessChain01(t *testing.T) {
	s := setup(t)
	s.Register(model.NewProfile("p", "A", "f", "2030-01-01"), model.NewEvent("e", "p", "X", "2030-01-01"))
	done := make(chan bool, 2)
	for _, v := range []string{"first", "second"} {
		go func(n string) { _, e := s.Confirm("e", n, n); done <- e == nil }(v)
	}
	if !(<-done && <-done) {
		t.Fatal()
	}
	e, err := s.Find("e")
	if err != nil {
		t.Fatal(err)
	}
	if e.ReminderNote != "first | second" && e.ReminderNote != "second | first" {
		t.Fatalf("lost concurrent confirmation notes: %q", e.ReminderNote)
	}
}
