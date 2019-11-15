package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongourl1 string
var mongourl2 string

func main() {
	flag.StringVar(&mongourl1, "mongourl1", "", "")
	flag.StringVar(&mongourl2, "mongourl2", "", "")
	flag.Parse()

	log.Printf("start dial to %s", mongourl1)
	// replSet initiate got NodeNotElectable: This node, mongo2:27017, with _id MemberId(1)
	// is not electable under the new configuration version 1 for replica set rs0 while validating
	// session, err := mgo.Dial(mongourl1 + "," + mongourl2 + "/?connect=direct")
	session, err := mgo.Dial(mongourl1 + "/?connect=direct")
	if err != nil {
		message := fmt.Sprintf("dial to %s failed: %s", mongourl1, err)
		log.Fatal(message)
	}
	defer session.Close()
	log.Printf("connect %v", session)

	session.SetMode(mgo.Monotonic, true)

	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		// https://stackoverflow.com/questions/44196113/is-it-possible-to-run-mongo-replicaset-commands-using-mgo-driver
		config := bson.M{
			"_id": "rs0",
			"members": []bson.M{
				{"_id": 0, "host": mongourl1},
				{"_id": 1, "host": mongourl2, "priority": 0}, // slave
			},
		}
		result := bson.M{}
		err := session.Run(bson.M{"replSetInitiate": config}, &result)
		if err == nil {
			log.Printf("replica set initiate result: %v", result)
			break
		}
		log.Printf("replica set initiate failed: %s", err)

	}

	for range ticker.C {
		result, err := exec(session, "admin", "replSetGetStatus")
		if err != nil {
			log.Printf("get replica set status failed: %s", err)
		} else {
			fmt.Println(result)
		}
	}
}

func exec(session *mgo.Session, name, command string) (bson.M, error) {
	result := bson.M{}
	err := session.DB(name).Run(bson.D{{command, 1}}, &result)
	return result, err
}
