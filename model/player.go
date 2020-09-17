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
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Score    uint32             `json:"score,omitempty" bson:"score,omitempty"`
	Level    uint8              `json:"level,omitempty" bson:"level,omitempty"`
}

// CheckUsername verifies the username update the player does
func (p *Player) CheckUsername() (bool, string) {
	if len(p.Username) < 6 || len(p.Username) > 40 {
		return false, "Username should be 6 to 40 characters long."
	}

	return checkUniqueProperty(Player{Username: p.Username})
}

// Valid checks this validity of the Type Person
func (p *Player) Valid() (bool, string) {
	if len(p.Username) < 6 || len(p.Username) > 40 {
		return false, "Username should be 6 to 40 characters long."
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
