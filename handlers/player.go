package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sarthakpranesh/Questioner/controllers"

	"github.com/sarthakpranesh/Questioner/model"
)

// CreatePlayerHandler creates a player if all checks pass
func CreatePlayerHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var p model.Player
	json.NewDecoder(request.Body).Decode(&p)
	b, s := p.Valid()
	if b == false {
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseString(s))
		return
	}
	result, err := model.AddPlayer(p)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err))
		return
	}
	json.NewEncoder(response).Encode(result)
}
