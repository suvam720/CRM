package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pin_code" bson:"pin_code"`
	Country string `json:"country" bson:"country"`
}

type User struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Age     int                `json:"age" bson:"age"`
	Gender  string             `json:"gender" bson:"gender"`
	Address Address            `json:"address" bson:"address"`
}

