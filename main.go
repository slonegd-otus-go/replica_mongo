package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

var mongourl string

func main() {
	flag.StringVar(&mongourl, "mongourl", "", "")
	flag.Parse()

	log.Printf("start dial to %s", mongourl)
	session, err := mgo.Dial(mongourl)
	if err != nil {
		message := fmt.Sprintf("dial to %s failed: %s", mongourl, err)
		log.Fatal(message)
	}
	defer session.Close()
	log.Printf("connect %v", session)
}
