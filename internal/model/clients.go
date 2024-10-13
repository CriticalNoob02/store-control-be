package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	IdClient string             `bson:"id" validate:"required,min=32"`
	Name     string             `bson:"name" validate:"required,min=2,max=100"`
	Cpf      string             `bson:"cpf" validate:"required,len=11"`
	Number   string             `bson:"number"`
}

var ClientFilterList = []string{"name", "cpf", "number"}
