package archive

import (
	"memorial/model"
	"testing"
	"time"
)

func TestEligibility(t *testing.T) {
	e := model.NewEvent("e", "p", "x", "2000-01-01")
	e.Status = "confirmed"
	if !ShouldArchive(e, time.Now()) {
		t.Fatal()
	}
}
