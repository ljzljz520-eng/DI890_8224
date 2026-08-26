package api

import (
	"memorial/query"
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if h.Svc.Store.Health() != nil {
			http.Error(w, "down", 503)
			return
		}
		w.WriteHeader(200)
	})
	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		x, e := query.Upcoming(h.Svc.Store)
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		writeJSON(w, x)
	})
	return mux
}
