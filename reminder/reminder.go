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

// lockFor returns a per-event mutex so concurrent confirmations of the same
// event are serialized instead of clobbering each other's notes.
func (m *Manager) lockFor(eventID string) *sync.Mutex {
	v, _ := m.locks.LoadOrStore(eventID, &sync.Mutex{})
	return v.(*sync.Mutex)
}

func (m *Manager) Confirm(eventID, actor, note string) (model.Record, error) {
	mu := m.lockFor(eventID)
	mu.Lock()
	defer mu.Unlock()

	ev, e := m.Store.GetEvent(eventID)
	if e != nil {
		return model.Record{}, e
	}
	// Model the persistence round-trip latency. The per-event lock held above
	// serializes concurrent confirmations, so this window can no longer let one
	// confirmation overwrite another's note.
	time.Sleep(10 * time.Millisecond)
	if !CanConfirm(ev, actor) {
		return model.Record{}, fmt.Errorf("event not confirmable: %s", ev.Status)
	}
	rec := model.Record{ID: fmt.Sprintf("%s-%s-%d", eventID, actor, time.Now().UnixNano()), EventID: eventID, Content: note, ConfirmedBy: actor, ConfirmedAt: time.Now().UTC().Format(time.RFC3339), Version: 1}
	if e = m.Store.SaveRecord(rec); e != nil {
		return rec, e
	}
	// Merge with any note already recorded so both confirmations survive,
	// rather than the later submitter overwriting the earlier one.
	ev.Status = "confirmed"
	ev.ReminderNote = MergeNotes(ev.ReminderNote, note)
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
