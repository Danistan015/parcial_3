package main

import (
    "context"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type ClientRepository struct{}

func (r *ClientRepository) Insert(client Client) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := os.Getenv("MONGO_URI")
    mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return err
    }
    defer mongoClient.Disconnect(ctx)

    db := os.Getenv("MONGO_DB")
    col := os.Getenv("MONGO_COLLECTION")

    _, err = mongoClient.Database(db).Collection(col).InsertOne(ctx, client)
    return err
}
