package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	IdClient string             `bson:"id"`
	Name     string             `bson:"name"`
	Cpf      string             `bson:"cpf"`
	Number   string             `bson:"number"`
}
