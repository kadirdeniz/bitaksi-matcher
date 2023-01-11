package internal

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID          primitive.ObjectID `json:"driver_id_id,omitempty"`
	Distance    float64            `json:"distance"`
	Coordinates []float64          `json:"coordinates"`
}
