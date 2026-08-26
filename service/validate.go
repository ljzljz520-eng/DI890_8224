package service

import (
	"errors"
	"memorial/model"
)

func ValidateSubmission(p model.Profile, e model.Event) error {
	if err := p.Validate(); err != nil {
		return err
	}
	if err := e.Validate(); err != nil {
		return err
	}
	if p.ID != e.ProfileID {
		return errors.New("profile mismatch")
	}
	return nil
}
