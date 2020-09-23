package question

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sarthakpranesh/Questioner/connect"
	"github.com/sarthakpranesh/Questioner/model"
	"go.mongodb.org/mongo-driver/bson"
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

// GetQuestions retrives a player info from the database
func GetQuestions() ([]model.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := connect.Collection("test", "question")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error in GetQuestion:", err.Error())
		return []model.Question{}, err
	}
	var qarr []model.Question
	for cursor.Next(ctx) {
		var q model.Question
		cursor.Decode(&q)
		q.Answer = []string{}
		qarr = append(qarr, q)
	}
	return qarr, nil
}
