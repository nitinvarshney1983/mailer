package main

import (
	"context"
	"fmt"
	"time"

	persistence "../persistence"
	log "github.com/nitinvarshney1983/mailer/logging"
	"go.mongodb.org/mongo-driver/bson"
)

func printTime() {

	//for {
	time.Sleep(100 * time.Millisecond)
	log.Info(time.Now())
	conn := persistence.GetConnection("etDB")
	coll := conn.Collection("user")
	var result bson.A
	err := coll.FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	//}

}
