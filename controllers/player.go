package controllers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

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

// UpdatePlayer updates the given player by matching it with id
func UpdatePlayer(id primitive.ObjectID, p model.Player) (model.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := connect.Collection("test", "player")
	var pNew model.Player
	err := collection.FindOneAndUpdate(
		ctx,
		model.Player{ID: id},
		p,
	).Decode(&pNew)
	if err != nil {
		log.Println("Error in UpdatePlayer:", err.Error())
		return model.Player{}, err
	}
	return pNew, nil
}
