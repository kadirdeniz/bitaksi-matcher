package mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"matcher/internal"
)

var Coordiantes = &internal.Coordinates{
	Latitude:  1.0,
	Longitude: 1.0,
}

var Location = &internal.Location{
	ID:          primitive.NewObjectID(),
	Coordinates: []float64{1.0, 1.0},
	Distance:    1.0,
}
