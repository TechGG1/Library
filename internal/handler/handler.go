package handler

import (
	"github.com/TechGG1/Library/internal/logging"
	"github.com/TechGG1/Library/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	service *service.Service
	logger  logging.Logger
}

func NewHandler(service *service.Service, logger logging.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/book", h.Books).Methods(http.MethodGet)
	r.HandleFunc("/book", h.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/reader", h.CreateReader).Methods(http.MethodPost)
	r.HandleFunc("/reader", h.Reader).Methods(http.MethodGet)
	r.HandleFunc("/reader", h.UpdateReader).Methods(http.MethodPut)
	r.HandleFunc("/rent", h.CreateRent).Methods(http.MethodPost)
	r.HandleFunc("/rent", h.UpdateRent).Methods(http.MethodPut)
	r.HandleFunc("/fine", h.Fine).Methods(http.MethodGet)
	r.HandleFunc("/rent", h.Rents).Methods(http.MethodGet)
	return r
}
