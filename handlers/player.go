package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
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
	result, err := controllers.AddPlayer(p)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err))
		return
	}
	json.NewEncoder(response).Encode(result)
}

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
	result, err2 := controllers.GetPlayer(id)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseError(err))
		return
	}
	json.NewEncoder(response).Encode(result)
}

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
	p.ID = primitive.ObjectID{}
	p.Level = 0
	p.Score = 0
	result, err2 := controllers.UpdatePlayer(id, p)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err2))
		return
	}
	json.NewEncoder(response).Encode(result)
}
