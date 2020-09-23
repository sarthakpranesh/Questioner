package player

import (
	"encoding/json"
	"net/http"

	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/player"
	"github.com/sarthakpranesh/Questioner/model"
)

type PlayerInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Score    uint32 `json:"score"`
	Level    uint8  `json:"level"`
}

type getPlayerResponse struct {
	Message string     `json:"message"`
	Payload PlayerInfo `json:"payload"`
}

// GetPlayerHandler retrives the player using player id
func GetPlayerHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	id, err := controllers.ParseToken(request)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write(controllers.ResponseError(err))
		return
	}
	result, err := player.GetPlayer(model.Player{ID: id})
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseError(err))
		return
	}
	gpr := getPlayerResponse{
		Message: "Player information retrived",
		Payload: PlayerInfo{
			Username: result.Username,
			Email:    result.Email,
			Score:    result.Score,
			Level:    result.Level,
		},
	}
	json.NewEncoder(response).Encode(gpr)
}
