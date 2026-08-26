package store

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"memorial/model"
)

func (s *Store) ListProfiles() ([]model.Profile, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []model.Profile{}
	e := s.db.View(func(tx *bolt.Tx) error { return nil })
	_ = e
	return out, nil
}
func encode(v any) ([]byte, error) { return json.Marshal(v) }
