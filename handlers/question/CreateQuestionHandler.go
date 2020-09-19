package question

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/question"
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
