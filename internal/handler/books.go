package handler

import (
	"encoding/json"
	"github.com/TechGG1/Library/internal/model"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ErrResponse struct {
	message string `json:"message"`
}

func (h *Handler) Books(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	books, err := h.service.Books(limit, page)
	if err != nil {
		h.logger.Log.Error("Error in retrieving Books", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		h.logger.Log.Error("Error in encoding Books", zap.Error(err), zap.String("url", r.URL.Path))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		h.logger.Log.Error("Error in decoding Book", zap.Error(err), zap.Stringp("url", &r.URL.Path))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookId, err := h.service.CreateBook(&book)
	if err != nil {
		h.logger.Log.Error("Error in retrieving Books", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]int{"book_id": bookId})
	if err != nil {
		h.logger.Log.Error("Error in encoding Book", zap.Error(err), zap.Stringp("url", &r.URL.Path))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
