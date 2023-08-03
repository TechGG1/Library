package handler

import (
	"encoding/json"
	"github.com/TechGG1/Library/internal/model"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) CreateRent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rent model.Rent
	err := json.NewDecoder(r.Body).Decode(&rent)
	if err != nil {
		h.logger.Log.Error("Error in decoding rent", zap.Error(err), zap.Stringp("url", &r.URL.Path))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rentId, err := h.service.CreateRent(r.Context(), &rent)
	if err != nil {
		h.logger.Log.Error("Error in retrieving rent", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]int{"rent_id": rentId})
	if err != nil {
		h.logger.Log.Error("Error in encoding rent", zap.Error(err), zap.Stringp("url", &r.URL.Path))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
