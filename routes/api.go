package routes

import (
	"api-go/controller"

	"github.com/gorilla/mux"
)

func AddApiRoute(r *mux.Router) {
	r.HandleFunc("/api/banks", controller.GetUser).Methods("GET")
	r.HandleFunc("/api/banks", controller.CreateUser).Methods("POST")
}
