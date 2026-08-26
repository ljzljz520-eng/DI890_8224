package archive

import (
	"memorial/model"
	"time"
)

func ShouldArchive(e model.Event, now time.Time) bool {
	if e.Archived {
		return false
	}
	if e.Status != "confirmed" {
		return false
	}
	d, x := time.Parse("2006-01-02", e.EventDate)
	return x == nil && d.Before(now)
}
func Reason(e model.Event) string {
	if e.Archived {
		return "already archived"
	}
	if e.Status != "confirmed" {
		return "awaiting confirmation"
	}
	return "eligible"
}
