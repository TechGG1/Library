package handler

import (
	"encoding/json"
	"fmt"
	"github.com/TechGG1/Library/internal/model"
	"go.uber.org/zap"
	"net/http"
	"strconv"
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

func (h *Handler) UpdateRent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rent model.Rent
	err := json.NewDecoder(r.Body).Decode(&rent)
	if err != nil {
		h.logger.Log.Error("Error in decoding rent", zap.Error(err), zap.Stringp("url", &r.URL.Path))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rentId, err := h.service.UpdateRent(r.Context(), &rent)
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
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Rents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	readerId, err := strconv.Atoi(r.URL.Query().Get("reader_id"))
	if err != nil {
		h.logger.Log.Error("Error in parsing reader_id", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		h.logger.Log.Error("Error in parsing page", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		h.logger.Log.Error("Error in parsing limit", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Print(page, limit)

	rents, pageFromReq, err := h.service.Rents(r.Context(), limit, page, readerId)
	if err != nil {
		h.logger.Log.Error("Error in retrieving rents", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(rents)
	if err != nil {
		h.logger.Log.Error("Error in encoding rents", zap.Error(err), zap.String("url", r.URL.Path))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("page", strconv.Itoa(pageFromReq))
	w.WriteHeader(http.StatusOK)
}
