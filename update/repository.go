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

func (r *ClientRepository) Update(id string, client Client) error {

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("MONGO_DB")
    colName := os.Getenv("MONGO_COLLECTION")

    mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return err
    }

    collection := mongoClient.Database(dbName).Collection(colName)

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    update := bson.M{
        "$set": bson.M{
            "name":  client.Name,
            "email": client.Email,
            "phone": client.Phone,
        },
    }

    _, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
    return err
}
