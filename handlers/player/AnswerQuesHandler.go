package player

import (
	"encoding/json"
	"net/http"

	"github.com/sarthakpranesh/Questioner/model"

	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/player"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnswerRequest struct {
	QID    string `json:"qId"`
	Answer string `json:"answer"`
}

type AnswerResponse struct {
	Message   string       `json:"message"`
	IsCorrect bool         `json:"isCorrect"`
	Player    model.Player `json:"player"`
}

func AnswerQuesHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	id, err := controllers.ParseToken(request)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write(controllers.ResponseError(err))
		return
	}
	var a AnswerRequest
	json.NewDecoder(request.Body).Decode(&a)
	quesID, err := primitive.ObjectIDFromHex(a.QID)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseError(err))
		return
	}
	p, b, err := player.AnswerQues(id, quesID, a.Answer)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err))
		return
	}
	var ar AnswerResponse = AnswerResponse{
		Player:    p,
		IsCorrect: b,
		Message:   "Answer checking successful",
	}
	json.NewEncoder(response).Encode(ar)
}
