package model

import (
	"context"
	"sync"
	"time"

	"github.com/sarthakpranesh/Questioner/connect"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Player struct describes a players attributes
type Player struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Score     uint32             `json:"score,omitempty" bson:"score,omitempty"`
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

	var wg sync.WaitGroup
	checkArray := []Player{Player{Email: p.Email}, Player{Username: p.Username}}
	i := len(checkArray)
	check := make(chan bool, i)
	whatFailed := make(chan string, i)
	defer close(check)
	defer close(whatFailed)
	wg.Add(i)
	for _, v := range checkArray {
		val := v
		go func() {
			b, s := checkUniqueProperty(val)
			check <- b
			whatFailed <- s
			wg.Done()
		}()
	}

	wg.Wait()
	for ; i > 0; i-- {
		if <-check == false {
			return false, <-whatFailed
		}
		<-whatFailed
	}

	return true, ""
}

func checkUniqueProperty(p Player) (bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := connect.Collection("test", "player")
	var found, empty Player
	_ = collection.FindOne(ctx, p).Decode(&found)
	if found != empty {
		return false, "Username/Email already in use!"
	}
	return true, ""
}
