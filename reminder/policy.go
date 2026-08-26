package reminder

import (
	"memorial/model"
	"strings"
)

func CanConfirm(e model.Event, actor string) bool {
	return !e.Archived && e.Status != "archived" && strings.TrimSpace(actor) != ""
}
func MergeNotes(old, new string) string {
	if old == "" {
		return new
	}
	if new == "" {
		return old
	}
	return old + " | " + new
}
func IsUrgent(e model.Event) bool { return e.Kind == "birthday" || e.Kind == "anniversary" }
