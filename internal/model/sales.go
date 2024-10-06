package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sale struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty"`
	Client  primitive.ObjectID   `bson:"client"`
	Items   []primitive.ObjectID `bson:"items"`
	Amounts []int                `bson:"amounts"`
	Cost    float64              `bson:"cost"`
	Payment bool                 `bson:"payment"`
	Date    time.Time            `bson:"date"`
}
