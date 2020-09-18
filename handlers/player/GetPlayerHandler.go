package player

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/player"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPlayerHandler retrives the player using player id
func GetPlayerHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	id, err := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Unable to retrive object Id!, Error:", err.Error())
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseError(err))
		return
	}
	result, err2 := player.GetPlayer(id)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseError(err))
		return
	}
	result.Password = ""
	json.NewEncoder(response).Encode(result)
}
