package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sale struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"user"`
	IdSale  string             `bson:"id"`
	Client  string             `bson:"client"`
	Items   []string           `bson:"items"`
	Amounts []int              `bson:"amounts"`
	Cost    float64            `bson:"cost"`
	Payment bool               `bson:"payment"`
	Date    time.Time          `bson:"date"`
}
