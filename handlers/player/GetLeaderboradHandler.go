package player

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/sarthakpranesh/Questioner/controllers"

	"github.com/sarthakpranesh/Questioner/connect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type leaderboardResponse struct {
	Message string       `json:"message"`
	Payload []PlayerInfo `json:"payload"`
}

// GetLeaderboardHandler retrives the Leaderboard of the current game
func GetLeaderboardHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	collection := connect.Collection("test", "player")
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"score", -1}})
	findOptions.SetLimit(15)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Println("Error:", err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(controllers.ResponseError(err))
	}
	var parr []PlayerInfo
	for cursor.Next(ctx) {
		var p PlayerInfo
		cursor.Decode(&p)
		parr = append(parr, p)
	}
	var glh leaderboardResponse = leaderboardResponse{
		Message: "Leader board fetched!",
		Payload: parr,
	}
	json.NewEncoder(response).Encode(glh)
}
