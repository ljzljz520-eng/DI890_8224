package service

import (
	"fmt"
	"memorial/model"
	"memorial/reminder"
	"memorial/store"
)

type Service struct {
	Store    *store.Store
	Reminder *reminder.Manager
}

func New(s *store.Store) *Service { return &Service{Store: s, Reminder: reminder.New(s)} }
func (s *Service) Register(p model.Profile, e model.Event) error {
	if x := p.Validate(); x != nil {
		return x
	}
	if x := e.Validate(); x != nil {
		return x
	}
	if x := s.Store.SaveProfile(p); x != nil {
		return x
	}
	return s.Store.SaveEvent(e)
}
func (s *Service) Confirm(id, actor, note string) (model.Record, error) {
	return s.Reminder.Confirm(id, actor, note)
}
func (s *Service) Find(id string) (model.Event, error) { return s.Store.GetEvent(id) }
func (s *Service) Archive(id, actor string) error {
	e, x := s.Store.GetEvent(id)
	if x != nil {
		return x
	}
	if e.Status != "confirmed" {
		return fmt.Errorf("cannot archive %s", e.Status)
	}
	e.Archived = true
	e.Status = "archived"
	if x = s.Store.SaveEvent(e); x != nil {
		return x
	}
	return s.Store.SaveAudit(store.NewAudit(id, actor, "archive", "event archived"))
}
