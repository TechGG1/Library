package handler

import (
	"github.com/gorilla/mux"
	"library/internal/logging"
	"library/internal/service"
)

type Handler struct {
	service *service.Service
	logger  logging.Logger
}

func InitRotes() *mux.Router {
	r := mux.NewRouter()
	//r.HandleFunc("/book", Controller.NewHomeController).Methods("GET")          //1
	//r.HandleFunc("/book", Controller.NewSaveBookController).Methods("POST")     //2
	//r.HandleFunc("/reader", Controller.NewSaveReederController).Methods("POST") //3
	//r.HandleFunc("/reader", Controller.NewGetReedersController).Methods("GET")  //4
	//r.HandleFunc("/rent", Controller.NewSaveRentController).Methods("POST")     //5
	//r.HandleFunc("/rent/complete", Controller.NewCompleteRentController).Methods("PUT")
	//r.HandleFunc("/rent", Controller.NewGetRentController).Methods("GET")
	//r.HandleFunc("/cover", Controller.NewSaveCoverController).Methods("POST")
	//r.HandleFunc("/cover", Controller.NewGetCoverController).Methods("GET")
	//r.HandleFunc("/coverId", Controller.NewGetCoverController1).Methods("GET")
	return r
}
