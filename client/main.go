package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

type TestDoc struct {
	ID int `bson:"id"`
}

var mongourl string

func main() {
	flag.StringVar(&mongourl, "mongourl", "", "")
	flag.Parse()

	log.Printf("start dial to %s", mongourl)
	// replSet initiate got NodeNotElectable: This node, mongo2:27017, with _id MemberId(1)
	// is not electable under the new configuration version 1 for replica set rs0 while validating
	// session, err := mgo.Dial(mongourl1 + "," + mongourl2 + "/?connect=direct")
	session, err := mgo.Dial(mongourl)
	if err != nil {
		message := fmt.Sprintf("dial to %s failed: %s", mongourl, err)
		log.Fatal(message)
	}
	defer session.Close()
	log.Printf("connect %v", session)

	// session.SetMode(mgo.Monotonic, true)

	db := session.DB("test")

	collection := db.C("test")

	ticker := time.NewTicker(1 * time.Second)
	i := 1

	for range ticker.C {
		doc := TestDoc{i}
		err := collection.Insert(doc)
		i++
		if err != nil {
			log.Printf("insert %v failed: %s", doc, err)
		}

		n, err := collection.Count()
		if err != nil {
			log.Printf("get count failed: %s", err)
			continue
		}
		log.Printf("got count: %d", n)

	}
}
