package handler

import (
	"encoding/json"
	"github.com/TechGG1/Library/internal/model"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (h *Handler) Fine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Query().Get("rent_id"))
	if err != nil {
		h.logger.Log.Error("Error in parsing rent_id", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rent := &model.Rent{
		RentId: id,
	}
	rentWithFine, err := h.service.CalculateFine(r.Context(), rent.RentId)
	if err != nil {
		h.logger.Log.Error("Error in retrieving rent", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(rentWithFine)
	if err != nil {
		h.logger.Log.Error("Error in encoding rent", zap.Error(err), zap.String("url", r.URL.Path))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
