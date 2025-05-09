package customer

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Customer struct {
		Id         primitive.ObjectID `bson:"_id,omitempty"`
		UserName   string             `bson:"username"`
		Email      string             `bson:"email"`
		ImageUrl   string             `bson:"image_url"`
		Password   string             `bson:"password"`
		Created_At time.Time          `bson:"created_at"`
		Updated_At time.Time          `bson:"updated_at"`
	}

	Claims struct {
		Id string `json:"customer_id"`
	}

	Credential struct {
		Id           primitive.ObjectID `bson:"_id,omitempty"`
		CustomerId   string             `bson:"customer_id,omitempty"`
		Rolecode     int                `bson:"role_code"`
		AccessToken  string             `bson:"access_token"`
		RefreshToken string             `bson:"refresh_token"`
		CreatedAt    time.Time          `bson:"created_at"`
		UpdatedAt    time.Time          `bson:"updated_at"`
	}

	UpdateRefreshToken struct {
		CustomerId   string    `bson:"customer_id"`
		AccessToken  string    `bson:"access_token"`
		RefreshToken string    `bson:"refresh_token"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}

	AuthClaims struct {
		*Claims
		jwt.RegisteredClaims
	}
)
