package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		result := bson.M{}
		err = session.DB("test").Run(bson.D{{"serverStatus", 1}}, &result)
		if err != nil {
			log.Printf("get server status failed: %s", err)
		} else {
			fmt.Println(result)
		}
	}

}
