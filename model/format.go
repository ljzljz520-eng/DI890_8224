package model

import "time"

func DateLabel(v string) string {
	t, e := time.Parse("2006-01-02", v)
	if e != nil {
		return "invalid"
	}
	return t.Format("02 Jan 2006")
}
func NormalizeStatus(v string) string {
	switch v {
	case "confirmed", "archived":
		return v
	default:
		return "pending"
	}
}
