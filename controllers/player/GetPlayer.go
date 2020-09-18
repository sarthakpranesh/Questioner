package player

import (
	"context"
	"log"
	"time"

	"github.com/sarthakpranesh/Questioner/connect"
	"github.com/sarthakpranesh/Questioner/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPlayer retrives a player info from the database
func GetPlayer(id primitive.ObjectID) (model.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := connect.Collection("test", "player")
	var p model.Player
	err := collection.FindOne(ctx, model.Player{ID: id}).Decode(&p)
	if err != nil {
		log.Println("Error in GetPlayer:", err.Error())
		return model.Player{}, err
	}
	return p, nil
}
