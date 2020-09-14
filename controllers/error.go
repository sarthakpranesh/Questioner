package controllers

import (
	"encoding/json"
	"log"
)

type RespError struct {
	Message string `json:"message"`
}

// ResponseError formats a json error
func ResponseError(err error) []byte {
	bs, err := json.Marshal(RespError{err.Error()})
	if err != nil {
		log.Println("Json Marshal error while ResponseError creation, ", err.Error())
		return []byte{}
	}
	return bs
}

func ResponseString(s string) []byte {
	bs, err := json.Marshal(RespError{s})
	if err != nil {
		log.Println("Json Marshal error while ResponseError creation, ", err.Error())
		return []byte{}
	}
	return bs
}
