package handler

import (
	"github.com/TechGG1/Library/internal/logging"
	"github.com/TechGG1/Library/internal/service"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/book", h.Books).Methods("GET")
	r.HandleFunc("/book", h.CreateBook).Methods("POST")
	r.HandleFunc("/reader", h.CreateReader).Methods("POST")
	r.HandleFunc("/reader", h.Reader).Methods("GET")
	r.HandleFunc("/reader", h.UpdateReader).Methods("PUT")
	r.HandleFunc("/rent", h.CreateRent).Methods("POST")
	r.HandleFunc("/rent", h.UpdateRent).Methods("PUT")
	r.HandleFunc("/fine", h.Fine).Methods("GET")
	r.HandleFunc("/rent", h.Rents).Methods("GET")
	return r
}
