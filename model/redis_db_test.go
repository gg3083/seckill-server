package model

import (
	"log"
	"testing"
)

func TestGetOneId(t *testing.T) {
	id, err := GetOneId("goods")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(id)
}