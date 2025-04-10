package main

import (
	"log"
	"net/http"

	"github.com/cleisondev/apirestcrud/config"
)

func main(){
	dbConnection := config.SetupDB()

	defer dbConnection.Close()
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}