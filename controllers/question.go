package controllers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sarthakpranesh/Questioner/connect"
	"github.com/sarthakpranesh/Questioner/model"
)

// AddQuestion inserts a question into the database
func AddQuestion(q model.Question) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := connect.Collection("test", "question")
	result, err := collection.InsertOne(ctx, q)
	if err != nil {
		log.Println("Unable to insert question, Error:", err.Error())
		return &mongo.InsertOneResult{}, err
	}
	return result, nil
}
