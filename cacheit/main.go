package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sauravrana646/cacheit/controller"
)

func main() {
	fmt.Println("Welcome to Cache It....")

	r := mux.NewRouter()
	r.HandleFunc("/", controller.HomeHandler).Methods("GET")
	r.HandleFunc("/getall", controller.GetAll).Methods("GET")
	r.HandleFunc("/get/{id}", controller.GetOne).Methods("GET")
	r.HandleFunc("/login/github", controller.OAuth).Methods("GET")
	r.HandleFunc("/add", controller.AddOne).Methods("POST")
	r.HandleFunc("/delete/{id}", controller.DeleteOne).Methods("DELETE")
	r.HandleFunc("/update/{id}", controller.UpdateOne).Methods("PUT")

	// controller.GetRaw()
	log.Fatal(http.ListenAndServe(":7000", r))
}
