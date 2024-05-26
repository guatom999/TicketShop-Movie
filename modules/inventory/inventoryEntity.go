package inventory

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	CustomerTicket struct {
		Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		CustomerId string             `json:"customer_id" bson:"customer_id"`
		MovieId    string             `json:"movie_id" bson:"movie_id"`
		MovieName  string             `json:"movie_name" bson:"movie_name"`
		Created_At string             `json:"created_at" bson:"created_at"`
		Price      string             `json:"price" bson:"price"`
		Seat       []string           `json:"seat" bson:"seat"`
	}
)
