package player

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/sarthakpranesh/Questioner/connect"
	"github.com/sarthakpranesh/Questioner/controllers/question"
	"github.com/sarthakpranesh/Questioner/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AnswerQues(pId, qId primitive.ObjectID, answer string) (model.Player, bool, error) {
	errChan := make(chan error, 2)
	var p model.Player
	var q model.Question
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		var err error
		p, err = GetPlayer(model.Player{ID: pId})
		if err != nil {
			errChan <- err
		}
		wg.Done()
	}()
	go func() {
		var err error
		q, err = question.GetQuestion(qId)
		if err != nil {
			errChan <- err
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		return model.Player{}, false, err
	default:

	}
	var checkAns bool = false
	for _, v := range q.Answer {
		if v == answer {
			checkAns = true
			break
		}
	}
	if checkAns == true {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := connect.Collection("test", "player")
		update := bson.M{
			"$set": bson.M{
				"score": p.Score + 100,
				"level": p.Level + 1,
			},
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
			model.Player{ID: pId},
			update,
			&opt,
		).Decode(&pNew)
		p.Score += 100
		p.Level++
		if err != nil {
			log.Println("Error:", err.Error())
			return model.Player{}, false, err
		}
		return pNew, true, nil
	}
	return p, false, nil
}
