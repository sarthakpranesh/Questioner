package main

import (
	"log"
	"net/http"

	"github.com/sarthakpranesh/Questioner/handlers"

	"github.com/joho/godotenv"
	"github.com/sarthakpranesh/Questioner/connect"

	"github.com/gorilla/mux"
)

func main() {
	godotenv.Load(".env")

	cancel := connect.Mongo()
	defer cancel()

	router := mux.NewRouter()
	router.HandleFunc("/player", handlers.CreatePlayerHandler).Methods("POST")
	router.HandleFunc("/player/update", handlers.UpdatePlayerHandler).Methods("POST")
	router.HandleFunc("/player/{id}", handlers.GetPlayerHandler).Methods("GET")
	router.HandleFunc("/question", handlers.CreateQuestion).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Unable to start server!, Error:", err.Error())
	}
}
