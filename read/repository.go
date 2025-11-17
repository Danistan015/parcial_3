package main

import (
    "context"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type ClientRepository struct{}

func (r *ClientRepository) GetAll() ([]Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := os.Getenv("MONGO_URI")
    mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }

    db := os.Getenv("MONGO_DB")
    col := os.Getenv("MONGO_COLLECTION")
    collection := mongoClient.Database(db).Collection(col)

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }

    var clients []Client
    if err := cursor.All(ctx, &clients); err != nil {
        return nil, err
    }

    return clients, nil
}

func (r *ClientRepository) GetByID(id string) (Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := os.Getenv("MONGO_URI")
    mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return Client{}, err
    }

    db := os.Getenv("MONGO_DB")
    col := os.Getenv("MONGO_COLLECTION")
    collection := mongoClient.Database(db).Collection(col)

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return Client{}, err
    }

    var client Client
    err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&client)
    return client, err
}
