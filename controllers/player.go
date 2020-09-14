package controllers

import (
	"context"
	"log"
	"time"

	"github.com/sarthakpranesh/Questioner/connect"
	"github.com/sarthakpranesh/Questioner/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddPlayer adds a player to the player collection
func AddPlayer(p model.Player) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := connect.Collection("test", "player")
	result, err := collection.InsertOne(ctx, p)
	if err != nil {
		log.Println("Mongo Insertion Error:", err.Error())
		return nil, err
	}
	return result, nil
}
