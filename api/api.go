package api

import (
	"encoding/json"
	"memorial/model"
	"memorial/service"
	"net/http"
)

type Handler struct{ Svc *service.Service }

func New(s *service.Service) *Handler { return &Handler{Svc: s} }
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		e, x := h.Svc.Find(id)
		if x != nil {
			http.Error(w, x.Error(), 404)
			return
		}
		json.NewEncoder(w).Encode(e)
	case http.MethodPost:
		var e model.Event
		if json.NewDecoder(r.Body).Decode(&e) != nil {
			http.Error(w, "bad json", 400)
			return
		}
		if x := h.Svc.Store.SaveEvent(e); x != nil {
			http.Error(w, x.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(e)
	default:
		http.Error(w, "method", 405)
	}
}
