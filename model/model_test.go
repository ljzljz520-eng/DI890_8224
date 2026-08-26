package model

import "testing"

func TestValidation(t *testing.T) {
	if NewEvent("", "p", "x", "bad").Validate() == nil {
		t.Fatal()
	}
	if NewProfile("p", "n", "r", "2030-01-01").Validate() != nil {
		t.Fatal()
	}
}
