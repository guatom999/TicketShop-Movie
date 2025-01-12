package inventory

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CustomerTicket struct {
		Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		CustomerId   string             `json:"customer_id" bson:"customer_id"`
		MovieId      string             `json:"movie_id" bson:"movie_id"`
		MovieName    string             `json:"movie_name" bson:"movie_name"`
		PosterUrl    string             `json:"poster_url" bson:"poster_url"`
		OrderNumber  string             `json:"order_number" bson:"order_number"`
		Ticket_Image string             `json:"ticket_image" bson:"ticket_image"`
		Created_At   time.Time          `json:"created_at" bson:"created_at"`
		Price        int64              `json:"price" bson:"price"`
		Seat         []string           `json:"seat" bson:"seat"`
	}
)
