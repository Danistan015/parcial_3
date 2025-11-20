package main

import (
    "context"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// este repo es el que se encarga de hablar con mongo
type ClientRepository struct{}

// Insert guarda un cliente en mongo
func (r *ClientRepository) Insert(client Client) error {

    // creo un contexto con tiempo limite por si mongo se demora mucho
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // leo la uri desde las variables de entorno del docker-compose
    uri := os.Getenv("MONGO_URI")

    // me conecto a mongo con esa uri
    mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return err // si falla la conexion, devuelvo el error
    }
    defer mongoClient.Disconnect(ctx) // al final cierro la conexion

    // tambien leo el nombre de la base y la coleccion desde env
    db := os.Getenv("MONGO_DB")
    col := os.Getenv("MONGO_COLLECTION")

    // aca hago el insert tal cual
    _, err = mongoClient.Database(db).Collection(col).InsertOne(ctx, client)

    // si InsertOne falla, retorno ese error, si no, queda en nil
    return err
}
