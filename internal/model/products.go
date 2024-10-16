package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	IdProduct string             `bson:"id" validate:"required,min=32"`
	Name      string             `bson:"name" validate:"required,min=2,max=100"`
	Amount    int                `bson:"amount" validate:"required,min=1"`
	Image     string             `bson:"image" validate:"required"`
	Category  string             `bson:"category" validate:"required"`
	Price     float64            `bson:"price" validate:"required"`
}

var ProductFilterList = []string{"name", "amount", "category", "price"}
