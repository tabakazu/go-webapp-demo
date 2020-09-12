package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/tabakazu/golang-webapi-demo/db"
	"github.com/tabakazu/golang-webapi-demo/model"
)

type itemsHandler struct {
	db db.DB
}

func (h *itemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.index(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func (h *itemsHandler) index(w http.ResponseWriter, r *http.Request) {
	items, err := h.db.GetAllItems(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&items); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *itemsHandler) create(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item.ID = uuid.New().String()

	if err := h.db.CreateItem(r.Context(), &item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
