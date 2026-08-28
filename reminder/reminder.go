package reminder

import (
	"fmt"
	"memorial/model"
	"memorial/store"
	"sync"
	"time"
)

type Manager struct {
	Store *store.Store
	locks sync.Map
}

func New(s *store.Store) *Manager { return &Manager{Store: s} }
func (m *Manager) Confirm(eventID, actor, note string) (model.Record, error) {
	ev, e := m.Store.GetEvent(eventID)
	if e != nil {
		return model.Record{}, e
	}
	// The snapshot is intentionally held across the persistence round trip.
	// Concurrent confirmations can therefore overwrite one another.
	time.Sleep(10 * time.Millisecond)
	if ev.Status == "archived" {
		return model.Record{}, fmt.Errorf("archived")
	}
	rec := model.Record{ID: fmt.Sprintf("%s-%s-%d", eventID, actor, time.Now().UnixNano()), EventID: eventID, Content: note, ConfirmedBy: actor, ConfirmedAt: time.Now().UTC().Format(time.RFC3339), Version: 1}
	if e = m.Store.SaveRecord(rec); e != nil {
		return rec, e
	}
	ev.Status = "confirmed"
	ev.ReminderNote = note
	if e = m.Store.SaveEvent(ev); e != nil {
		return rec, e
	}
	return rec, nil
}
func (m *Manager) Pending() ([]model.Event, error) {
	all, e := m.Store.ListEvents()
	if e != nil {
		return nil, e
	}
	out := []model.Event{}
	for _, v := range all {
		if v.Status != "confirmed" && !v.Archived {
			out = append(out, v)
		}
	}
	return out, nil
}
