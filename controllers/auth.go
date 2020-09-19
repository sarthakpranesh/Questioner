package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ParseToken checks the header Authorization to retrive the user Object id
func ParseToken(request *http.Request) (primitive.ObjectID, error) {
	token := strings.ReplaceAll(request.Header.Values("Authorization")[0], "Bearer ", "")
	bs, err := jwt.DecodeSegment(token)
	if err != nil {
		log.Println("Unable to DecodeSegment:", err.Error())
		return primitive.ObjectID{}, err
	}
	ID, err := primitive.ObjectIDFromHex(string(bs))
	if err != nil {
		log.Println("Unable to convert string id to ObjectID:", err.Error())
		return primitive.ObjectID{}, err
	}
	return ID, nil
}
