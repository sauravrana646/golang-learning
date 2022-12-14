package router

import (
	"github.com/gorilla/mux"
	"github.com/sauravrana646/golang-learning/Golang_Auth/jwthttpserver/controllers"
)

func Router(r *mux.Router) {
	r.HandleFunc("/", controllers.Base)

	r.HandleFunc("/login", controllers.Login).Methods("GET")
	r.HandleFunc("/home", controllers.Home).Methods("GET")
	r.HandleFunc("/refreshtoken", controllers.RefreshToken).Methods("GET")
}
