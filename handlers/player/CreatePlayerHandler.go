package player

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sarthakpranesh/Questioner/controllers"
	"github.com/sarthakpranesh/Questioner/controllers/player"
	"github.com/sarthakpranesh/Questioner/model"
)

type CreatePlayerResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

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
	if p.Password != "" {
		h := sha256.New()
		h.Write([]byte(p.Password))
		p.Password = base64.URLEncoding.EncodeToString(h.Sum(nil))
	} else {
		h := sha256.New()
		h.Write([]byte(p.UID))
		p.UID = base64.URLEncoding.EncodeToString(h.Sum(nil))
	}
	result, err := player.AddPlayer(p)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err))
		return
	}
	token := jwt.EncodeSegment([]byte(result.InsertedID.(primitive.ObjectID).Hex()))
	var cpr *CreatePlayerResponse = &CreatePlayerResponse{
		Message: "Player Registered!",
		Token:   token,
	}
	json.NewEncoder(response).Encode(cpr)
}
