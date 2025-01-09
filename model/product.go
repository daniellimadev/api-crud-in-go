package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Product represents the structure in the database.
type Product struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Price    float64            `json:"price" bson:"price"`
	Quantity int                `json:"quantity" bson:"quantity"`
}
