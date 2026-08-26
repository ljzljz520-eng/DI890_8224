package service

import (
	"memorial/model"
	"memorial/query"
)

func (s *Service) Pending() ([]model.Event, error) { return query.Upcoming(s.Store) }
func (s *Service) Status(status string) ([]model.Event, error) {
	return query.ByStatus(s.Store, status)
}
func (s *Service) Summary() (map[string]int, error) {
	all, e := s.Store.ListEvents()
	if e != nil {
		return nil, e
	}
	m := map[string]int{}
	for _, v := range all {
		m[v.Status]++
	}
	return m, nil
}
