package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Amount   int                `bson:"amount"`
	Image    string             `bson:"image"`
	Category string             `bson:"category"`
	Price    float64            `bson:"price"`
}
