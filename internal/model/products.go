package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	IdProduct uuid.UUID          `bson:"id" validate:"required,uuid"`
	Name      string             `bson:"name" validate:"required,min=2,max=100"`
	Amount    int                `bson:"amount"`
	Image     string             `bson:"image"`
	Category  string             `bson:"category"`
	Price     float64            `bson:"price"`
}

var ProductFilterList = []string{"name", "amount", "category", "price"}
