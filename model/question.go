package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Question struct represents an actual question
type Question struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Question string             `json:"question,omitempty" bson:"question,omitempty"`
	Answer   []string           `json:"answer,omitempty" bson:"answer,omitempty"`
	Level    uint8              `json:"level" bson:"level"`
	Extra    string             `json:"extra,omitempty" bson:"extra,omitempty"` // this should be encoded JSON of whatever extra things that you want to store, like hints, descriptions, etc.
}

// Valid method validates the question
func (q Question) Valid() (bool, string) {
	if len(q.Name) < 3 && len(q.Name) > 20 {
		return false, "Question's name should be 3 to 20 characters long!"
	}

	if len(q.Question) < 6 && len(q.Question) > 140 {
		return false, "Question should be 6 to 140 characters long!"
	}

	if len(q.Answer) < 1 && len(q.Answer) > 140 {
		return false, "Question's answer should be 1 to 140 characters long!"
	}

	if q.Level <= 0 {
		return false, "Question's level cannot be negative!"
	}

	return true, ""
}
