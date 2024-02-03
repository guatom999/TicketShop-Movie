package book

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	BookMovieReq struct {
		Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		CustomerId string             `bson:"customer_id"`
		MovieName  string             `json:"movie_name"`
		Quantity   int                `bson:"quantity"`
		Theater    string             `json:"theater"`
		Seat       []Seat             `json:"seat"`
	}

	Seat struct {
		Number string `bson:"number"`
	}
)
