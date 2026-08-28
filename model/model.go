package model

import (
	"errors"
	"strings"
	"time"
)

type Profile struct {
	ID, Name, Relation, ImportantDate, Notes string
	Archived                                 bool
}
type Event struct {
	ID, ProfileID, Title, EventDate, Kind, Status, ReminderNote string
	Archived                                                    bool
}
type Record struct {
	ID, EventID, Content, ConfirmedBy, ConfirmedAt string
	Version                                        int
}
type Audit struct{ ID, EventID, Actor, Action, Detail, CreatedAt string }

func (p Profile) Validate() error {
	if strings.TrimSpace(p.ID) == "" {
		return errors.New("profile id required")
	}
	if p.Name == "" {
		return errors.New("name required")
	}
	if _, e := time.Parse("2006-01-02", p.ImportantDate); e != nil {
		return e
	}
	return nil
}
func (e Event) Validate() error {
	if e.ID == "" || e.ProfileID == "" {
		return errors.New("event identity required")
	}
	if e.Title == "" {
		return errors.New("title required")
	}
	if _, x := time.Parse("2006-01-02", e.EventDate); x != nil {
		return x
	}
	return nil
}
func (r Record) Validate() error {
	if r.ID == "" || r.EventID == "" || r.Content == "" {
		return errors.New("record incomplete")
	}
	return nil
}
func (a Audit) Validate() error {
	if a.ID == "" || a.EventID == "" || a.Action == "" {
		return errors.New("audit incomplete")
	}
	return nil
}
func NewEvent(id, profile, title, date string) Event {
	return Event{ID: id, ProfileID: profile, Title: title, EventDate: date, Kind: "anniversary", Status: "pending"}
}
func NewProfile(id, name, relation, date string) Profile {
	return Profile{ID: id, Name: name, Relation: relation, ImportantDate: date}
}
