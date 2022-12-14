package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sauravrana646/golang-learning/Golang_Auth/jwthttpserver/router"
)

func main() {
	fmt.Println("Welcome to Golang JWT HTTP Server....")
	r := mux.NewRouter()
	router.Router(r)
	log.Fatal(http.ListenAndServe(":7000", r))
}
