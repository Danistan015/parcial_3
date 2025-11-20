package main

// este struct es el cliente que vamos a guardar en mongo y mandar en json
type Client struct {
    ID     string `json:"id" bson:"_id,omitempty"`      // id del cliente, mongo lo pone solo
    Name   string `json:"name" bson:"name"`             // nombre del cliente
    Email  string `json:"email" bson:"email"`           // correo del cliente
    Phone  string `json:"phone" bson:"phone"`           // telefono del cliente
}
