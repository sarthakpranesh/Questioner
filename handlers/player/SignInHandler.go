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

type signInResponse struct {
	Message string     `json:"message"`
	Token   string     `json:"token"`
	Payload PlayerInfo `json:"payload"`
}

// SignInHandler helps the user retrive the auth token
func SignInHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var p model.Player
	json.NewDecoder(request.Body).Decode(&p)
	var err error
	var result model.Player
	if p.UID != "" {
		h := sha256.New()
		h.Write([]byte(p.UID))
		uid := base64.URLEncoding.EncodeToString(h.Sum(nil))
		result, err = player.GetPlayer(model.Player{UID: uid})
	} else if p.Email != "" && p.Password != "" {
		result, err = player.GetPlayer(model.Player{Email: p.Email})
	} else {
		response.WriteHeader(http.StatusBadRequest)
		response.Write(controllers.ResponseString("Improper authentication details!"))
		return
	}

	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write(controllers.ResponseString("User not found!"))
		return
	}

	if p.Email != "" && p.Password != "" && p.UID == "" {
		h := sha256.New()
		h.Write([]byte(p.Password))
		PassHash := base64.URLEncoding.EncodeToString(h.Sum(nil))
		if PassHash != result.Password {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write(controllers.ResponseString("Incorrect Password"))
			return
		}
	}

	token := jwt.EncodeSegment([]byte(result.ID.Hex()))
	sir := signInResponse{
		Message: "Successful sign in",
		Token:   token,
		Payload: PlayerInfo{
			Username: result.Username,
			Email:    result.Email,
			Score:    result.Score,
			Level:    result.Level,
		},
	}
	json.NewEncoder(response).Encode(sir)
}
