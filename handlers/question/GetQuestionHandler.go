package question

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/question"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
