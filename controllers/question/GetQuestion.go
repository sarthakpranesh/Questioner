package question

import (
	"context"
	"log"
	"time"

	"github.com/sarthakpranesh/Questioner/connect"
	"github.com/sarthakpranesh/Questioner/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetQuestion retrives a player info from the database
func GetQuestion(id primitive.ObjectID) (model.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := connect.Collection("test", "question")
	var q model.Question
	err := collection.FindOne(ctx, model.Question{ID: id}).Decode(&q)
	if err != nil {
		log.Println("Error in GetQuestion:", err.Error())
		return model.Question{}, err
	}
	return q, nil
}
