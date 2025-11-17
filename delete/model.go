package main

type Client struct {
    ID    string `json:"id" bson:"_id,omitempty"`
    Name  string `json:"name" bson:"name"`
    Email string `json:"email" bson:"email"`
    Phone string `json:"phone" bson:"phone"`
}
