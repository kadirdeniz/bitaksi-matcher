package internal

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Distance    float64            `bson:"distance"`
	Coordinates []float64          `bson:"coordinates"`
}
