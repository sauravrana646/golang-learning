package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sauravrana646/cacheit/utils"
)

const pathToConfig string = "/home/unthinkable/saurav/golang-learning/cacheit/config"

func OAuth(w http.ResponseWriter, r *http.Request) {
	config, err := utils.LoadConfig(pathToConfig)
	if err != nil {
		log.Fatal("Cannot Load Config:", err)
	}

	fmt.Println("Config is ", config)
}
