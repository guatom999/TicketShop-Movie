package customer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Customer struct {
		Id         primitive.ObjectID `bson:"_id"`
		UserName   string             `bson:"username"`
		Email      string             `bson:"email"`
		Password   string             `bson:"password"`
		Created_At time.Time          `bson:"created_at"`
		Updated_At time.Time          `bson:"updated_at"`
	}
)
