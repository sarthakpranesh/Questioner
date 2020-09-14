package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sarthakpranesh/Questioner/connect"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Score     uint32             `json:"score" bson:"score"`
}

// Valid checks this validity of the Type Person
func (p *Player) Valid() (bool, string) {
	if len(p.Username) < 6 || len(p.Username) > 40 {
		return false, "Username should be 6 to 40 characters long."
	}

	if len(p.Firstname) < 3 || len(p.Firstname) > 30 {
		return false, "Firstname should be 3 to 30 characters long."
	}

	if len(p.Lastname) < 3 || len(p.Lastname) > 30 {
		return false, "Lastname should be 3 to 30 characters long."
	}

	return true, ""
}

// AddPlayer adds a player to the player collection
func AddPlayer(p Player) (*mongo.InsertOneResult, error) {
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
