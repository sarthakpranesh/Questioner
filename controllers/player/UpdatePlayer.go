package player

import (
	"context"
	"log"
	"time"

	"github.com/sarthakpranesh/Questioner/connect"
	"github.com/sarthakpranesh/Questioner/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdatePlayerUsername updates the given player by matching it with id
func UpdatePlayerUsername(id primitive.ObjectID, username string) (model.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := connect.Collection("test", "player")
	update := bson.M{
		"$set": bson.M{"username": username},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var pNew model.Player
	err := collection.FindOneAndUpdate(
		ctx,
		model.Player{ID: id},
		update,
		&opt,
	).Decode(&pNew)
	if err != nil {
		log.Println("Error in UpdatePlayer:", err.Error())
		return model.Player{}, err
	}
	return pNew, nil
}
