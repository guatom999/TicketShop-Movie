package customer

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Customer struct {
		Id       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email    string             `json:"email" bson:"email,omitempty"`
		UserName string
	}
)
