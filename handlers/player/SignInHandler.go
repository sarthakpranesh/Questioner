package player

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/player"

	"github.com/sarthakpranesh/Questioner/model"
)

type SignInResponse struct {
	Message string            `json:"message"`
	Token   string            `json:"token"`
	Player  GetPlayerResponse `json:"player"`
}

// SignInHandler helps the user retrive the auth token
func SignInHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var p model.Player
	json.NewDecoder(request.Body).Decode(&p)
	if p.Email == "" || p.Password == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseString("Improper authentication details!"))
		return
	}
	result, err := player.GetPlayer(model.Player{Email: p.Email})
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write(controllers.ResponseString("User not found!"))
		return
	}
	if result.Password == "" || result.UID != "" {
		response.WriteHeader(http.StatusNotAcceptable)
		response.Write(controllers.ResponseString("Please use OAuth!"))
		return
	}
	h := sha256.New()
	h.Write([]byte(p.Password))
	PassHash := base64.URLEncoding.EncodeToString(h.Sum(nil))
	if PassHash != result.Password {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write(controllers.ResponseString("Incorrect Password"))
		return
	}
	token := jwt.EncodeSegment([]byte(result.ID.Hex()))
	sir := SignInResponse{
		Message: "Successful sign in",
		Token:   token,
		Player: GetPlayerResponse{
			Message:  "Player information",
			Username: result.Username,
			Email:    result.Email,
			Score:    result.Score,
			Level:    result.Level,
		},
	}
	json.NewEncoder(response).Encode(sir)
}
