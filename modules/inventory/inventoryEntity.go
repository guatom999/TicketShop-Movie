package inventory

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Inventory struct {
		Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		CustomerId string             `json:"customer_id" bson:"customer_id"`
		MovieId    string             `json:"movie_id" bson:"movie_id"`
	}
)
