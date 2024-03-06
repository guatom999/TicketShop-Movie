package ticket

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Ticket struct {
		Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		MovieId    string             `json:"movie_id" bson:"movie_id"`
		CustomerId string             `json:"customer_id" bson:"customer_id"`
		Title      string             `json:"name" bson:"title"`
		Price      float64            `json:"price" bson:"price"`
		Seat       string             `json:"seat" bson:"seat"`
		Date       string             `json:"date" bson:"date"`
		Time       string             `json:"time" bson:"time"`
	}
)
