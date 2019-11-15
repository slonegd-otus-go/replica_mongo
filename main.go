package main

import (
	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return
	}
	defer session.Close()
}
