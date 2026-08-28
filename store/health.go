package store

import "fmt"

func (s *Store) Health() error {
	if s == nil || s.db == nil {
		return fmt.Errorf("store unavailable")
	}
	return nil
}
func (s *Store) CountBuckets() int { return len(buckets) }
