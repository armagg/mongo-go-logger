package main

import (
   "context"
   "time"
   "fmt"


   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
   "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
    client, err := mongo.NewClient(options.Client().ApplyURI())
    if err != nil {
		log.Fatal(err)
	
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
            log.Fatal(err)
    }
    defer client.Disconnect(ctx)
    }