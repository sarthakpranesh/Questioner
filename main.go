package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sarthakpranesh/Questioner/handlers/player"
	"github.com/sarthakpranesh/Questioner/handlers/question"

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
	router.HandleFunc("/player/sign_in", player.SignInHandler).Methods("POST")
	router.HandleFunc("/player", player.GetPlayerHandler).Methods("GET")
	router.HandleFunc("/player/play", player.AnswerQuesHandler).Methods("POST")
	router.HandleFunc("/player/leaderboard", player.GetLeaderboardHandler).Methods("GET")
	router.HandleFunc("/player/update", player.UpdatePlayerHandler).Methods("POST")
	router.HandleFunc("/questions", question.GetQuestionsHandler).Methods("GET")
	router.HandleFunc("/question", question.CreateQuestionHandler).Methods("POST")
	router.HandleFunc("/question/{id}", question.DeleteQuestionHandler).Methods("DELETE")

	err := http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), router)
	if err != nil {
		log.Fatalln("Unable to start server!, Error:", err.Error())
	}
}
