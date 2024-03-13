package ticket

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Ticket struct {
		Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		MovieId    string             `json:"movie_id" bson:"movie_id"`
		CustomerId string             `json:"customer_id" bson:"customer_id"`
		MovieName  string             `json:"movie_name" bson:"movie_name"`
		Seat       string             `json:"seat" bson:"seat"`
		Date       string             `json:"date" bson:"date"`
		Time       string             `json:"time" bson:"time"`
		Price      float64            `json:"price" bson:"price"`
		CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
