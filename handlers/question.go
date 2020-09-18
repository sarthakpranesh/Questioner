package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/question"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sarthakpranesh/Questioner/model"
)

// CreateQuestionHandler handler validates and creates the question document
func CreateQuestionHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var q model.Question
	err := json.NewDecoder(request.Body).Decode(&q)
	if err != nil {
		log.Println("Unable to decode body, Error:", err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err))
		return
	}
	b, s := q.Valid()
	if b == false {
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseString(s))
		return
	}
	result, err := question.AddQuestion(q)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err))
		return
	}
	json.NewEncoder(response).Encode(result)
}

// GetQuestionHandler retrives the question using question id
func GetQuestionHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	id, err := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error from GetQuestion:", err.Error())
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseError(err))
		return
	}
	q, err2 := question.GetQuestion(id)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err))
		return
	}
	json.NewEncoder(response).Encode(q)
}
