package question

import (
	"encoding/json"
	"net/http"

	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/question"
	"github.com/sarthakpranesh/Questioner/model"
)

type getQuestionResponse struct {
	Message string           `json:"message"`
	Payload []model.Question `json:"payload"`
}

// GetQuestionsHandler retrives all questions
func GetQuestionsHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	qarr, err2 := question.GetQuestions()
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err2))
		return
	}
	var gqh getQuestionResponse = getQuestionResponse{
		Message: "Questions retrived!",
		Payload: qarr,
	}
	json.NewEncoder(response).Encode(gqh)
}
