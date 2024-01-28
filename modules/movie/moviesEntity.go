package movie

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Movie struct {
		Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title     string             `json:"name" bson:"Title"`
		Price     float64            `json:"price" bson:"price"`
		ImageUrl  string             `json:"image_url" bson:"image_url"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
		Avaliable int                `json:"valiable" bson:"avaliable"`
		Category  string             `json:"category" bson:"category"`
	}

	Category struct {
		Id   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Name string             `json:"name" bson:"name"`
	}
)
