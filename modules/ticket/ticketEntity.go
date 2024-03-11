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
		Seat       string             `json:"seat" bson:"seat"`
		CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
