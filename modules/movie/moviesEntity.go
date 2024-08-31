package movie

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Movie struct {
		Id              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title           string             `json:"name" bson:"title"`
		Description     string             `json:"description" bson:"description"`
		RunningTime     string             `json:"running_time" bson:"running_time"`
		Price           float64            `json:"price" bson:"price"`
		ImageUrl        string             `json:"image_url" bson:"image_url"`
		CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
		Category        string             `json:"category" bson:"category"`
		ReleaseAt       time.Time          `json:"release_at" bson:"release_at"`
		OutOfTheatersAt time.Time          `json:"out_of_theaters_at" bson:"out_of_theaters_at"`
	}

	MovieAvaliable struct {
		Id            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Movie_Id      string             `json:"movie_id" bson:"movie_id"`
		Title         string             `json:"name" bson:"title"`
		CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
		Showtime      time.Time          `json:"showtime" bson:"show_time"`
		SeatAvailable []SeatAvailable    `json:"seat_available" bson:"seat_available"`
	}

	SeatAvailable map[string]bool

	Category struct {
		Id   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Name string             `json:"name" bson:"name"`
	}
)
