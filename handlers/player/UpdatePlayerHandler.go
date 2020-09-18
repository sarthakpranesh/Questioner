package player

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/player"
	"github.com/sarthakpranesh/Questioner/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdatePlayerHandler is used to update username of the player
func UpdatePlayerHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	s := strings.ReplaceAll(request.Header.Values("Authorization")[0], "Bearer ", "")
	id, _ := primitive.ObjectIDFromHex(s)
	var p model.Player
	err := json.NewDecoder(request.Body).Decode(&p)
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
