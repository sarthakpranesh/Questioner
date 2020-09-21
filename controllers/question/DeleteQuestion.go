package question

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sarthakpranesh/Questioner/model"

	"github.com/sarthakpranesh/Questioner/connect"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteQuestion takes in a question id and deletes that question
func DeleteQuestion(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := connect.Collection("test", "question")
	result, err := collection.DeleteOne(ctx, model.Question{ID: id})
	if err != nil {
		log.Println("Error while deleteing question,", err.Error())
		return &mongo.DeleteResult{}, err
	}
	return result, nil
}
