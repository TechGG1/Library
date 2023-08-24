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

	r.Use(corsMiddleware)

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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
