package player

import (
	"encoding/json"
	"net/http"

	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/player"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type answerRequest struct {
	QID    string `json:"qId"`
	Answer string `json:"answer"`
}

type answerResponse struct {
	Message   string     `json:"message"`
	IsCorrect bool       `json:"isCorrect"`
	Payload   PlayerInfo `json:"payload"`
}

// AnswerQuesHandler used to check if the user Answer is correct
func AnswerQuesHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	id, err := controllers.ParseToken(request)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write(controllers.ResponseError(err))
		return
	}
	var a answerRequest
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
	var ar answerResponse = answerResponse{
		Message:   "Answer checking successful",
		IsCorrect: b,
		Payload: PlayerInfo{
			Username: p.Username,
			Email:    p.Email,
			Score:    p.Score,
			Level:    p.Level,
		},
	}
	json.NewEncoder(response).Encode(ar)
}
