package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	IdClient uuid.UUID          `bson:"id" validate:"required,uuid"`
	Name     string             `bson:"name" validate:"required,min=2,max=100"`
	Cpf      string             `bson:"cpf" validate:"required,len=11"`
	Number   string             `bson:"number"`
}

var ClientFilterList = []string{"name", "cpf", "number"}
