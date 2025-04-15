package main

import (
	"log"
	"net/http"

	"github.com/cleisondev/apirestcrud/config"
	"github.com/cleisondev/apirestcrud/handlers"
	"github.com/cleisondev/apirestcrud/models"
	"github.com/gorilla/mux"
)

func main() {
	dbConnection := config.SetupDB()

	_, err := dbConnection.Exec(models.CreateTableSQL)
	defer dbConnection.Close()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.WriteTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
