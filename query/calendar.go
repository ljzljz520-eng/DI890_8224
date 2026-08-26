package query

import (
	"memorial/model"
	"memorial/store"
	"sort"
)

func ForProfile(s *store.Store, pid string) ([]model.Event, error) {
	all, e := s.ListEvents()
	if e != nil {
		return nil, e
	}
	out := []model.Event{}
	for _, v := range all {
		if v.ProfileID == pid {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].EventDate < out[j].EventDate })
	return out, nil
}
func Active(events []model.Event) []model.Event {
	out := make([]model.Event, 0, len(events))
	for _, v := range events {
		if !v.Archived {
			out = append(out, v)
		}
	}
	return out
}
