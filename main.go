package main

import (
	"log"
	"net/http"

	"github.com/sarthakpranesh/Questioner/handlers"
	"github.com/sarthakpranesh/Questioner/handlers/player"

	"github.com/joho/godotenv"
	"github.com/sarthakpranesh/Questioner/connect"

	"github.com/gorilla/mux"
)

func main() {
	godotenv.Load(".env")

	cancel := connect.Mongo()
	defer cancel()

	router := mux.NewRouter()
	router.HandleFunc("/player", player.CreatePlayerHandler).Methods("POST")
	router.HandleFunc("/player/update", player.UpdatePlayerHandler).Methods("POST")
	router.HandleFunc("/player/{id}", player.GetPlayerHandler).Methods("GET")
	router.HandleFunc("/question", handlers.CreateQuestionHandler).Methods("POST")
	router.HandleFunc("/question/{id}", handlers.GetQuestionHandler).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Unable to start server!, Error:", err.Error())
	}
}
