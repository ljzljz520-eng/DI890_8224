package store

import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"memorial/model"
	"sync"
	"time"
)

var buckets = []string{"profiles", "events", "records", "audits"}

type Store struct {
	db *bolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bolt.Tx) error {
		for _, n := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(n)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
	}
	return s, e
}
func (s *Store) Close() error { return s.db.Close() }
func put[T any](s *Store, b string, id string, v T) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bolt.Tx) error { return tx.Bucket([]byte(b)).Put([]byte(id), raw) })
}
func get[T any](s *Store, b, id string) (T, error) {
	var out T
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte(b)).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("not found")
		}
		return json.Unmarshal(v, &out)
	})
	return out, e
}
func (s *Store) SaveProfile(v model.Profile) error { return put(s, "profiles", v.ID, v) }
func (s *Store) GetProfile(id string) (model.Profile, error) {
	return get[model.Profile](s, "profiles", id)
}
func (s *Store) SaveEvent(v model.Event) error           { return put(s, "events", v.ID, v) }
func (s *Store) GetEvent(id string) (model.Event, error) { return get[model.Event](s, "events", id) }
func (s *Store) SaveRecord(v model.Record) error         { return put(s, "records", v.ID, v) }
func (s *Store) SaveAudit(v model.Audit) error           { return put(s, "audits", v.ID, v) }
func (s *Store) ListEvents() ([]model.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []model.Event{}
	e := s.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("events")).ForEach(func(_, v []byte) error {
			var x model.Event
			if e := json.Unmarshal(v, &x); e != nil {
				return e
			}
			out = append(out, x)
			return nil
		})
	})
	return out, e
}
func NewAudit(event, actor, action, detail string) model.Audit {
	return model.Audit{ID: fmt.Sprintf("%s-%d", event, time.Now().UnixNano()), EventID: event, Actor: actor, Action: action, Detail: detail, CreatedAt: time.Now().UTC().Format(time.RFC3339)}
}
