package handler

import (
	"encoding/json"
	"github.com/TechGG1/Library/internal/model"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (h *Handler) CreateReader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reader model.Reader
	err := json.NewDecoder(r.Body).Decode(&reader)
	if err != nil {
		h.logger.Log.Error("Error in decoding reader", zap.Error(err), zap.Stringp("url", &r.URL.Path))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	readerId, err := h.service.CreateReader(r.Context(), &reader)
	if err != nil {
		h.logger.Log.Error("Error in retrieving reader", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]int{"reader_id": readerId})
	if err != nil {
		h.logger.Log.Error("Error in encoding reader", zap.Error(err), zap.Stringp("url", &r.URL.Path))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Reader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	readers, pageFromReq, err := h.service.Readers(r.Context(), limit, page)
	if err != nil {
		h.logger.Log.Error("Error in retrieving readers", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(readers)
	if err != nil {
		h.logger.Log.Error("Error in encoding readers", zap.Error(err), zap.String("url", r.URL.Path))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("page", strconv.Itoa(pageFromReq))
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateReader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reader model.Reader
	err := json.NewDecoder(r.Body).Decode(&reader)
	if err != nil {
		h.logger.Log.Error("Error in decoding reader", zap.Error(err), zap.Stringp("url", &r.URL.Path))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	readerId, err := h.service.UpdateReader(r.Context(), &reader)
	if err != nil {
		h.logger.Log.Error("Error in retrieving reader", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]int{"reader_id": readerId})
	if err != nil {
		h.logger.Log.Error("Error in encoding reader", zap.Error(err), zap.Stringp("url", &r.URL.Path))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
