package player

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/player"
	"github.com/sarthakpranesh/Questioner/model"
)

// UpdatePlayerHandler is used to update username of the player
func UpdatePlayerHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	id, err := controllers.ParseToken(request)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write(controllers.ResponseError(err))
	}
	var p model.Player
	err = json.NewDecoder(request.Body).Decode(&p)
	if err != nil {
		log.Println("json decode request.Body failed in UpdatePlayerHandler:", err.Error())
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseError(err))
		return
	}
	b, s := p.CheckUsername()
	if b == false {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseString(s))
		return
	}
	result, err2 := player.UpdatePlayerUsername(id, p.Username)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err2))
		return
	}
	json.NewEncoder(response).Encode(result)
}
