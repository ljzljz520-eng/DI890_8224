package archive

import (
	"memorial/model"
	"memorial/store"
)

func Eligible(e model.Event) bool    { return e.Status == "confirmed" && !e.Archived }
func Mark(e model.Event) model.Event { e.Archived = true; e.Status = "archived"; return e }
func Audit(s *store.Store, e model.Event, actor string) error {
	return s.SaveAudit(store.NewAudit(e.ID, actor, "archive", "eligible event archived"))
}
