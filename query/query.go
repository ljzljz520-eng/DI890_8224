package query

import (
	"memorial/model"
	"memorial/store"
	"sort"
	"strings"
)

func ByStatus(s *store.Store, status string) ([]model.Event, error) {
	all, e := s.ListEvents()
	if e != nil {
		return nil, e
	}
	out := []model.Event{}
	for _, v := range all {
		if strings.EqualFold(v.Status, status) {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].EventDate < out[j].EventDate })
	return out, nil
}
func Upcoming(s *store.Store) ([]model.Event, error) {
	all, e := s.ListEvents()
	if e != nil {
		return nil, e
	}
	out := []model.Event{}
	for _, v := range all {
		if !v.Archived && v.Status != "archived" {
			out = append(out, v)
		}
	}
	return out, nil
}
func Count(s *store.Store) (int, error) { x, e := s.ListEvents(); return len(x), e }
